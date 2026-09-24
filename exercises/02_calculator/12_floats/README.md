# Floats

Integers had a strict chaperone: no mixing without conversion. Floats look
friendlier — until you notice the stub's bug hiding in plain sight:

```go
func Half(n int) float64 {
    return float64(n / 2) // looks right, halves wrong
}
```

`Half(5)` returns 2, not 2.5. The conversion is correct and the division is
correct; the *order* is the bug. Integer division runs first (`5/2 == 2`,
truncated, remainder abandoned), and converting the corpse to `float64`
can't resurrect the `.5`.

## Convert first, divide second

```go
func Half(n int) float64 {
    return float64(n) / 2 // 5 -> 5.0, then 5.0/2 == 2.5
}
```

`float64(n)` promotes the operand *before* the operator sees it, so `/`
becomes float division. The untyped constant `2` adapts to the occasion —
constants in Go are fluid until pinned down, which is why `float64(n) / 2`
needs no `float64(2)`. One conversion reroutes the whole expression.

This is the general rule behind three exercises' worth of traps: **in a
mixed expression, the types of the operands decide the operation, and the
operation decides what survives.** Truncation, wrapping, and narrowing all
strike at operation time. Convert at the leaves, before the damage, never
at the root after it.

## A word on float equality

The test compares with `==`:

```go
if got, want := Half(5), 2.5; got != want {
```

That's safe here — 2.5 is exactly representable in binary — but treat it
as an exception, not a license. Most decimal fractions (0.1, the classic)
are *not* exact, and `==` on computed floats compares bit patterns, not
intentions. Production float code compares within a tolerance or avoids
equality entirely. The `math` package (`math.Pi`, `math.Sqrt`, `math.Abs`)
covers the operations; judgment covers the comparisons.

## Task

Fix `Half` in `calc.go` so odd inputs keep their `.5`.

## Check

```bash
go test ./exercises/02_calculator/12_floats/ -v
```
