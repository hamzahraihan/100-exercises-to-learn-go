---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Building Strings Efficiently"
weight: 5
draft: false
---

# Building Strings Efficiently

Strings in Go are immutable, so each `+=` in a loop allocates a brand-new
string and copies everything so far — fine once, wasteful thousands of times.
`strings.Builder` exists for exactly this shape: write pieces into it, then ask
for the finished string once, with far fewer allocations. The stub returns
`""` always, so the repeat case fails (the zero-count case passes vacuously).

## Task

Fill in `ConcatN` in `concat.go`:

```go
func ConcatN(s string, n int) string {
	// ...
}
```

Loop `n` times writing into a `strings.Builder`, then return the built string.

## Check

```bash
go test ./exercises/09_strings/05_builder/ -v
```

---

*Source: `exercises/09_strings/05_builder/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
