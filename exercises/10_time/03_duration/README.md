# Durations

A `time.Time` is a moment. The gap between two moments is something else:
a `time.Duration` — and Go models it as what it is, a number of
nanoseconds wearing a good API. No separate interval type, no calendar
math surprises at this scale.

## Subtraction returns a quantity

```go
func HoursBetween(a, b time.Time) float64 {
    return b.Sub(a).Hours()
}
```

`b.Sub(a)` is "b minus a" as a `Duration`: positive when `b` is later,
negative when earlier, zero for identical moments. The test's three shapes
cover the space — 90 minutes apart yields `1.5`, same instant yields `0` —
and the signature's direction (`b` minus `a`, not `a` minus `b`) decides
the sign. Name the order in the function name (`HoursBetween(a, b)` reads
"from a to b") so callers never guess.

## A number with manners

Underneath, `Duration` is an `int64` of nanoseconds. What makes it pleasant
is the vocabulary around the number:

```go
time.Hour        // a Duration: exactly 3,600,000,000,000 ns
(90 * time.Minute).Hours()   // 1.5 — scale, then convert
d.String()       // "1h30m0s" — human-readable, free
d.Round(time.Minute) // shed sub-minute noise before displaying
```

Constants like `time.Hour` and `time.Minute` replace magic numbers —
`timeout := 5 * time.Second` reads as prose and compiles to integer math.
Converters (`Hours`, `Minutes`, `Seconds`) return fractional floats; 90
minutes is honestly `1.5`, no remainder ceremony. And `String()` renders
`"1h30m0s"` for logs without a formatting verb in sight.

One honesty clause: `Duration` counts *physical* elapsed time. Across a
daylight-saving transition, "24 hours later" and "tomorrow at the same
clock time" disagree by an hour — `AddDate` handles calendar days, `Add`
handles stopwatch time. Ninety-nine percent of durations never cross that
line; knowing which function serves which master covers the remaining one.

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
