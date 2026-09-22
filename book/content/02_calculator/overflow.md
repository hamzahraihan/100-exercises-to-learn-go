---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Overflow"
weight: 8
draft: false
---

# Overflow

Unsigned arithmetic in Go wraps silently on overflow: `uint8(200) + 100`
gives 44 with no panic (unlike Rust, which panics in debug builds).
Idiomatic Go checks explicitly: if `sum < a`, the addition wrapped around.

## Task

Fix `AddUint8` in `calc.go` to report wraparound via the `overflow` flag.

## Check

```bash
go test ./exercises/02_calculator/08_overflow/ -v
```

---

*Source: `exercises/02_calculator/08_overflow/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
