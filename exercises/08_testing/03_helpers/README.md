# Test Helpers

Repeated assertions belong in a helper function, and calling `t.Helper()` as
its first line marks it as infrastructure: when the helper calls `t.Fatalf`,
the failure is reported at the caller's line, not inside the helper. Without
that call you would see the same helper line for every failure and have to
guess which input broke.

## Task

Fill in `Mean` in `mean.go` so it averages the slice, returning 0 for empty
input:

```go
func Mean(xs []float64) float64 {
	// ...
}
```

Sum the elements and divide by their count, guarding `len(xs) == 0` first so
you never divide by zero.

## Check

```bash
go test ./exercises/08_testing/03_helpers/ -v
```
