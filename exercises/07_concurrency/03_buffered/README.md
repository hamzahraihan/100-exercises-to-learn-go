# Buffered Collect

The pipeline lesson's channel was unbuffered: every send blocked until a
receiver arrived. Sender and collector moved in lockstep, each paced by the
other. Sometimes that's exactly the backpressure you want. Here it isn't —
`n` workers shouldn't queue behind one drain. This exercise gives the
channel a waiting room.

## Unbuffered is a handshake

```go
ch := make(chan int) // capacity 0: send waits for receive, and vice versa
```

Every value changes hands directly, sender to receiver, in the same
instant. Elegant — and coupled: a fast worker idles until the collector
asks, a slow collector stalls every worker behind the current one. For
one-to-one rhythms that's fine. For fan-out, it's a traffic jam with extra
steps.

## Buffered is a queue

```go
ch := make(chan int, n) // capacity n: up to n sends complete immediately
```

A **buffered** channel holds up to `n` values without any receiver
present. Sends proceed until the buffer fills; receives proceed while it's
non-empty. Workers deposit results and exit; the collector drains at
leisure. Nobody waits on anybody — the pace decouples in both directions.

The sizing rule for this pattern: **buffer for every worker**. `n`
goroutines each send once, so capacity `n` guarantees no sender ever
blocks:

```go
func Collect(n int) []int {
    type result struct{ i, v int }
    ch := make(chan result, n)
    for i := 0; i < n; i++ {
        go func() {
            ch <- result{i, 2 * i}
        }()
    }
    out := make([]int, n)
    for range n {
        r := <-ch
        out[r.i] = r.v
    }
    return out
}
```

Size the buffer to the sends and the deadlock question answers itself:
every send has a slot, every slot gets drained, the function always
returns. Undersize it and workers park mid-send — with no receiver
running yet (collection happens after launching), the program would hang
forever. When a channel program deadlocks, count sends against buffer plus
receivers first; the arithmetic is almost always the bug.

The index tags do the same duty as the pipeline lesson: arrival order is
scheduler chaos, `out[r.i]` restores input order. Same trick, new reason
to need it — buffered collection scrambles even harder, since nothing
paces the workers at all.

## Task

Complete `Collect` in `collect.go` so `n` goroutines each send one doubled
index over a buffered channel and results land in order:

```go
func Collect(n int) []int {
    // ...
}
```

## Check

```bash
go test -race ./exercises/07_concurrency/03_buffered/ -v
```

The `-race` flag is required here: `n` goroutines writing through one
channel is exactly the shape the detector was built to verify.
