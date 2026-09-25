# Benchmarks

Tests answer "does it work?" Benchmarks answer "how fast?" — and Go keeps
the two side by side in the same file, sharing helpers and fixtures. This
exercise's package holds both: `TestFib` guards correctness while
`BenchmarkFib` times it, and a third test guards the benchmark itself.

## Measuring, not asserting

```go
func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fib(20)
    }
}
```

`BenchmarkXxx` takes `*testing.B`, and `b.N` is the runner's dial: it calls
your function with growing `N` until timings stabilize, then reports
nanoseconds per operation. You don't choose the iteration count — the
harness calibrates it. Your only job is putting the operation inside the
loop, exactly once per iteration, with setup hoisted outside.

Benchmarks never run under plain `go test`. Speed costs time, so it's
opt-in:

```bash
go test -bench . -run '^$' ./exercises/08_testing/06_benchmarks/
```

`-bench .` selects every benchmark; `-run '^$'` deselects every test
(the regex matches nothing), so the run measures without re-proving.
Memorize the pair — it's the standard incantation, and half of all
benchmark confusion is running the wrong combination.

## Two Fibs, one lesson

```go
// Iterative: two running values, O(n) time
func Fib(n int) int {
    a, b := 0, 1
    for i := 0; i < n; i++ {
        a, b = b, a+b
    }
    return a
}
```

The README suggests the loop first, and the numbers explain why. Naive
recursion recomputes the same values exponentially — `Fib(20)` already
performs tens of thousands of redundant calls — while the accumulator
version walks forward once. "Recursion also works" is permission for a
first draft, not a recommendation: run the benchmark on both and watch the
ns/op column differ by orders of magnitude. That comparison *is* the
lesson. Correctness tests can't distinguish the two implementations; the
benchmark can, which is why performance-critical code measures instead of
assuming.

`Fib(0) == 0, Fib(1) == 1` are the base facts the loop must reproduce —
trace `n = 0` (body never runs, `a` stays 0) and `n = 1` (one step: `a, b`
become `1, 1`, return 1) before trusting it. Edge-first reading, the table
habit applied to loops.

## Who watches the watcher

```go
func TestBenchmarkRuns(t *testing.T) {
    if res := testing.Benchmark(BenchmarkFib); res.N <= 0 {
        t.Fatal("benchmark did not run")
    }
}
```

Benchmarks are code, and code rots — a benchmark that stops executing (bad
flag, renamed function, skipped file) fails silently, reporting nothing
forever. This guard runs the benchmark *as a test* via `testing.Benchmark`
and asserts it actually iterated. Infrastructure testing itself: the same
instinct as `t.Helper` marking, one level up. When your own suites grow
benchmarks, steal this guard — an unmeasured benchmark is decoration.

## Task

Fill in `Fib` in `fib.go` so it returns the n-th Fibonacci number with
`Fib(0) == 0` and `Fib(1) == 1`:

```go
func Fib(n int) int {
    // ...
}
```

Accumulate two running values in a loop (recursion also works and is a fine
first draft, though much slower — compare the benchmark numbers after both).

## Check

```bash
go test ./exercises/08_testing/06_benchmarks/ -v
go test -bench . -run '^$' ./exercises/08_testing/06_benchmarks/
```
