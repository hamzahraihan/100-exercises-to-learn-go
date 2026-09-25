# Select with Timeout

Channels so far had one conversation each: send, receive, done. Real
programs juggle — data arriving *or* a deadline expiring, whichever comes
first. Waiting on one channel while another fires is how fetches hang
forever. `select` is how they don't.

## Waiting on whichever answers first

```go
// Syntax: select runs the first ready case, blocking until one is
select {
case v := <-dataCh:
    // data arrived first
case <-time.After(200 * time.Millisecond):
    // deadline hit first
}
```

`select` blocks until **one** of its cases can proceed, then runs exactly
that case. It reads like `switch`, but the branches are channel operations
and the condition is *readiness*, not equality. If several cases are ready
simultaneously, one is picked at random — fairness by design, so study the
implication: never write cases whose correctness depends on priority order.

`time.After(d)` returns a channel that delivers one value after `d`
elapses. It's a timer wearing a channel's clothes, which makes deadlines
composable with ordinary channel logic instead of a separate timeout
mechanism bolted on sideways.

## The fetch with a fuse

```go
func Fetch(d time.Duration) (string, error) {
    select {
    case <-time.After(d):
        return "data", nil
    case <-time.After(200 * time.Millisecond):
        return "", errors.New("fetch timed out")
    }
}
```

Two timers race. Fast simulated fetch (`d` = 10ms): the data case fires
first, success returns. Slow fetch (`d` = 500ms): the 200ms fuse blows
first, and the caller gets an error instead of a 500ms stare. The function
returns at 200ms either way — bounded latency, the whole point.

The test margins are engineered, and worth noticing: 10ms and 500ms sit
*far* from the 200ms boundary. Real schedulers jitter by milliseconds, so
asserting near the deadline would flake. The tests assert *outcomes*
(value vs error), never durations — and keep the cases clear of the fence.
When you write timing tests later, steal both habits: assert results, and
give the boundary a wide berth.

One loose end, honestly labeled: after a timeout win, the losing timer
still fires into the void — harmless here (garbage-collected channel, one
value, nobody listening). Production code with heavy timeout traffic
graduates to `time.NewTimer` with explicit `Stop`. For learning the shape,
`After` is exactly right.

## Task

Complete `Fetch` in `fetch.go` so a simulated fetch taking `d` succeeds on
the fast path and fails with a timeout error past 200ms:

```go
func Fetch(d time.Duration) (string, error) {
    // ...
}
```

## Check

```bash
go test -race ./exercises/07_concurrency/04_select/ -v
```
