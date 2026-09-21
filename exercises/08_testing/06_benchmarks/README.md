# Benchmarks

Functions named `BenchmarkXxx` taking `*testing.B` measure speed, not
correctness: the runner calls them with increasing `b.N` until the timing is
stable, then reports nanoseconds per operation and throughput. They live next
to ordinary tests — this package has both `TestFib` for correctness and
`BenchmarkFib` for speed, plus a `TestBenchmarkRuns` guard proving the
benchmark actually executes. Benchmarks never run during a plain `go test`;
you opt in explicitly.

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
