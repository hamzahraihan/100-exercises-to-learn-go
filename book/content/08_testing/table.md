---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Table-Driven Tests"
weight: 1
draft: false
---

# Table-Driven Tests

Most Go tests follow one pattern: a slice of `{input, want}` cases, a loop
that calls the function under test, and `t.Errorf` (not `Fatalf`) on mismatch
so every case reports instead of stopping at the first failure. The test here
checks four scores against four grades; the stub returns `""` for everything,
so all four rows fail at once — which is exactly the point.

## Task

Fill in `Grade` in `grade.go` so each score maps to its letter:

```go
func Grade(score int) string {
	// ...
}
```

Use an if/else chain on the 90/80/70 cutoffs (90+ is "A", 80+ is "B", 70+
is "C", anything below is "F").

## Check

```bash
go test ./exercises/08_testing/01_table/ -v
```

---

*Source: `exercises/08_testing/01_table/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
