# Mutex Counter

One hundred goroutines, one integer, each adding one. The answer should be
100. Run the check and watch it come back 97, 99, 100 — a different wrong
number every time. This exercise is about why one innocent line can't be
trusted, and the one lock that makes it trustworthy.

## One step that is three

```go
c.n += n
```

Reads as a single step. Compiles as three: load `c.n` into a register, add
`n`, store the result back. Two goroutines can interleave fatally — both
load 41, both add 1, both store 42 — and one increment vanishes without a
trace. No error, no panic, just a count that drifts downward under load.
Every shared-state bug in Go is some version of this interleaving.

## The test's machinery, decoded

Before the fix, read the harness — it introduces two concurrency tools at
once:

```go
var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() { defer wg.Done(); c.Add(1) }()
}
wg.Wait()
```

`go func() { ... }()` launches each increment as a **goroutine**: a
lightweight thread managed by the runtime, thousands per program without
breaking a sweat. `sync.WaitGroup` counts the outstanding ones — `Add(1)`
per launch, `Done()` per finish (deferred so it runs even on panic), and
`Wait()` blocks until the counter drains. Launch, track, rendezvous: the
standard fan-out shape you'll reuse in every concurrent exercise ahead.

Note what the test does *not* do: sleep, retry, or assume timing. `Wait`
makes completion deterministic; only the *interleaving* varies, which is
exactly the nondeterminism the lock must tame.

## The lock

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c *Counter) Add(n int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n += n
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.n
}
```

`sync.Mutex` is mutual exclusion: whoever holds the lock proceeds, everyone
else waits. `Lock` at the top, `defer Unlock()` beside it — deferred so
every return path, including panics, releases. Lock the *reader* too:
`Value` racing a concurrent `Add` is the same three-step interleaving with
a read in the middle. Unsynchronized reads are races, full stop — the
detector below says so emphatically.

And keep the critical section tiny: lock, touch `n`, unlock. Holding a
mutex across I/O, channel operations, or other locks is how deadlocks are
born. This course's later sections will show the alternative — sharing by
communicating instead of locking — but some state is genuinely shared, and
for that, the mutex is the honest tool.

## Prove it with -race

```bash
go test -race ./exercises/07_concurrency/01_mutex/ -v
```

The `-race` flag instruments memory accesses and reports `DATA RACE`
warnings with both goroutines' stack traces. It's required here for a
cruel reason: without it, the racy stub *sometimes passes*. Scheduling luck
masks the bug; the detector doesn't gamble. Run `-race` on concurrent code
as reflexively as `gofmt` on new code — luck is not a test strategy.

## Task

Guard `n` in `counter.go` with the existing `mu` field so 100 concurrent
`Add(1)` calls always leave `Value()` at 100, and the test passes.

## Check

```bash
go test -race ./exercises/07_concurrency/01_mutex/ -v
```

The `-race` flag is required here: without it the racy stub can
flake-pass, but under the race detector it reports `DATA RACE` warnings.
