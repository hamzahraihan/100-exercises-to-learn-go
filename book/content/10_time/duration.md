---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Durations"
weight: 3
draft: false
---

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

---

*Source: `exercises/10_time/03_duration/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
