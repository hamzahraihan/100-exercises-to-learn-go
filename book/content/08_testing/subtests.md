---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Subtests"
weight: 2
draft: false
---

# Subtests

`t.Run` splits one test into named subtests — "above", "below", "inside" here
— and each reports separately, so a failure tells you which case broke
instead of just which function. Run with `-v` and you get a line per
subtest; `go test -run 'TestClamp/below'` runs only the "below" case, which
is handy when debugging one branch of the logic.

## Task

Fill in `Clamp` in `clamp.go` so `n` is constrained to `[lo, hi]`:

```go
func Clamp(n, lo, hi int) int {
	// ...
}
```

Return `lo` when `n` is below it, `hi` when `n` is above it, otherwise `n`
itself.

## Check

```bash
go test ./exercises/08_testing/02_subtests/ -v
```

---

*Source: `exercises/08_testing/02_subtests/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
