# Test Helpers

The third time you type the same assertion, stop. Extract it — into a
helper function that takes `t` and does the checking. But helpers create a
new problem: when one fails, *whose line* does the failure point at? This
exercise's test calls the answer before you write a line of your own.

## Infrastructure wears a mark

```go
func mustMean(t *testing.T, xs []float64, want float64) {
    t.Helper()
    if got := Mean(xs); got != want {
        t.Fatalf("Mean(%v) = %v, want %v", xs, got, want)
    }
}
```

`t.Helper()` as the first line marks the function as scaffolding. When the
`Fatalf` inside fires, Go skips helper frames and reports the *caller's*
line — the `mustMean(t, ...)` invocation in the test, with its specific
input. Without that call, every failure would point inside the helper at
the same `t.Fatalf` line, and you'd guess which input broke. One line of
marking, failures that locate themselves.

The `must` prefix is convention, not syntax: `mustX` helpers fail the test
on the spot (`Fatalf` inside) rather than returning errors. `MustParse`
from the panic-boundary lesson wore the same prefix for the same reason —
*this either succeeds or the test is over*. Name helpers after what they
guarantee, and readers know the stakes before reading the body.

## The function under test

`Mean` averages a slice, with one trap laid deliberately:

```go
func Mean(xs []float64) float64 {
    if len(xs) == 0 {
        return 0
    }
    total := 0.0
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}
```

An empty slice divides by zero — `0/0` in floats is `NaN`, not a panic,
which is worse: a quiet poison that spreads through every downstream
computation. The `len(xs) == 0` guard returns an honest 0 instead (a
guard clause doing its oldest job), and the test pins it with a `nil`
input — `Mean(nil)` must be 0, exercising both the guard and the nil-slice
ranging rule from the calculator section in one call.

Note the conversion: `total / float64(len(xs))`. Integer division's ghost
from the floats lesson — operands decide the operation, so the count
converts before dividing. Same trap, new costume; the rule transfers
unchanged.

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
