# Durations

Subtracting two times with `Sub` gives a `time.Duration` — an `int64` count
of nanoseconds with friendly helpers. `Hours()` converts to fractional
hours, and constants like `time.Hour` let you build or scale durations
without magic numbers. The stub returns `0` for everything, so the
fractional case fails on the assertion.

## Task

Fill in `HoursBetween` in `hours.go`:

```go
func HoursBetween(a, b time.Time) float64 {
	// ...
}
```

Subtract with `b.Sub(a)` and convert the result with `Hours()`.

## Check

```bash
go test ./exercises/10_time/03_duration/ -v
```
