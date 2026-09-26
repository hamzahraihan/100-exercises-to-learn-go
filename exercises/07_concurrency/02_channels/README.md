# Channel Pipeline

The mutex lesson shared state and guarded it. This one inverts the
strategy: don't share the slice at all. Each input travels to its own
goroutine, each result travels back over a channel, and the output is
assembled where exactly one goroutine — the caller — ever touches it.

## Sending work out, collecting results back

```go
// Syntax: make the channel, send with <-, receive with <-
ch := make(chan int) // unbuffered: every send waits for a receiver
ch <- 42             // send (blocks until received)
v := <-ch            // receive (blocks until sent)
```

A **channel** is a typed conduit between goroutines. `go` launches the
worker; the channel carries the answer home:

```go
ch := make(chan int)
go func() {
    ch <- 21 * 2
}()
got := <-ch // 42
```

No lock, no shared variable, no race — the handoff *is* the
synchronization. Receiving waits for the send, so `got` is ready by
construction. Values cross goroutine boundaries; the slice never does.

## Order is the actual problem

Doubling is trivially parallel, but the test demands results *in input
order* — and goroutines finish in whatever order the scheduler fancies.
Collecting bare values in arrival order scrambles the output. Two cures:

**Tag each result with its index**, then place by tag:

```go
type result struct {
    i, v int
}

func DoubleAll(nums []int) []int {
    out := make([]int, len(nums))
    ch := make(chan result)
    for i, n := range nums {
        go func() {
            ch <- result{i, n * 2}
        }()
    }
    for range nums {
        r := <-ch
        out[r.i] = r.v
    }
    return out
}
```

Workers run wild; the collector — the only writer to `out` — restores
order from the tags. Distinct indices mean no two writes collide, and a
single writer means no lock.

**Or write `out[i]` from the worker directly** and synchronize separately
(a `WaitGroup`, or a done-channel drained `len(nums)` times). Fewer moving
parts, but the ordering burden moves into reasoning about who writes what.
Either shape passes; the tagged channel shows the channel doing real work,
so prefer it here.

One historical footnote, already handled: pre-1.22 Go reused loop variables
across iterations, so goroutines capturing `i` all saw the final value.
Modern Go (this course requires 1.23+) binds fresh variables per iteration
— the closure above is safe as written. If you ever backport concurrent
loops to older Go, `i := i` inside the loop is the inoculation.

## Task

Complete `DoubleAll` in `pipeline.go` so each input is doubled via a
goroutine and a channel, results land in input order, and the test passes.

## Check

```bash
go test -race ./exercises/07_concurrency/02_channels/ -v
```

The `-race` flag is required here: goroutines share nothing by design, and
the detector proves it instead of scheduling luck.
