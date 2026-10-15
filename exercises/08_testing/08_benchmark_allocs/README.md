# Benchmark Allocations

`ns/op` tells you how long. `allocs/op` tells you *why*.

A function can be slow for two unrelated reasons: too much work, or too many heap allocations. The allocator has a hidden cost — each heap object means a GC scan later. `b.ReportAllocs()` exposes what the profiler would eventually tell you, without needing `pprof`.

## ReportAllocs and ResetTimer

```go
func BenchmarkJoin(b *testing.B) {
    b.ReportAllocs() // print allocs/op and B/op columns
    words := []string{"the", "quick", "brown", "fox"}
    b.ResetTimer()   // exclude setup from measurement
    for b.Loop() {
        _ = strings.Join(words, " ")
    }
}
```

`b.ReportAllocs()` adds two columns to the output: `allocs/op` (heap objects created per call) and `B/op` (bytes allocated per call). Call it at the top of the benchmark before the loop.

`b.ResetTimer()` discards setup time; use it after expensive fixture creation so that allocation counts reflect only the work inside the loop.

> **Note:** `b.Loop()` is the modern Go 1.24+ idiom replacing the classic `for i := 0; i < b.N; i++` pattern. It handles warm-up and teardown correctly and avoids common off-by-one mistakes.

## Acting on the numbers

```
BenchmarkConcatSlow-8    1000    1234 ns/op    15 allocs/op    960 B/op
BenchmarkConcatFast-8    5000     210 ns/op     1 allocs/op     64 B/op
```

`ConcatSlow` uses `+=` on strings. Because Go strings are immutable, each `+=` allocates a brand-new backing array and copies both sides into it — O(n) allocations for n repetitions.

`ConcatFast` uses `strings.Builder`. The Builder maintains a mutable `[]byte` buffer, growing it geometrically when needed. With `Grow` pre-called, the buffer is sized exactly once, so the final `b.String()` call is the only heap allocation that escapes.

The allocation count is the tell — not the duration, which varies by machine and load.

## b.ReportAllocs is a contract

Once you add it to a benchmark, a regression that adds allocations shows up in CI without a profiler. It is a cheap, durable performance test that catches "oops I added a `fmt.Sprintf` in the hot path" at review time.

<details><summary>Hint</summary>

In `ConcatFast`, declare `var b strings.Builder`, call `b.Grow(len(s) * n)` to pre-allocate the exact number of bytes you will write, then call `b.WriteString(s)` in a loop, and finally return `b.String()`.

With `Grow` called before any writes, the Builder will not resize its internal buffer, so the only allocation that survives to the heap is the one backing the returned string — exactly 1 `allocs/op`.

</details>

## Task

1. Complete `ConcatFast(s string, n int) string` in `concat.go` using `strings.Builder` with `Grow` pre-allocation so that it allocates exactly **1** heap object regardless of `n`.
2. Complete `BenchmarkConcatFast` in `concat_test.go` — add `b.ReportAllocs()`, `b.ResetTimer()`, and the `b.Loop()` body.

The test `TestAllocsPerRun` uses `testing.AllocsPerRun` to assert that `ConcatFast` allocates at most 1 heap object for n=50. This makes the allocation constraint machine-checkable without running benchmarks manually.

## Check

```bash
go test ./exercises/08_testing/08_benchmark_allocs/ -v
go test -bench . -benchmem -run '^$' ./exercises/08_testing/08_benchmark_allocs/
```
