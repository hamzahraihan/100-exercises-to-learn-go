---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Test Helpers"
weight: 3
draft: false
---

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

---

*Source: `exercises/08_testing/03_helpers/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
