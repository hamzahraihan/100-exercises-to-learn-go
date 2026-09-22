---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Conversions"
weight: 10
draft: false
---

# Conversions

Go never converts numbers implicitly: assigning an `int32` to an `int64`
needs an explicit conversion written as `T(x)`, e.g. `int64(x)`. There is no
`as` cast keyword. Widening conversions like `int32` to `int64` always
preserve the value.

## Task

Fix `ToInt64` in `calc.go` to convert `x` with `int64(x)`.

## Check

```bash
go test ./exercises/02_calculator/10_conversions/ -v
```

---

*Source: `exercises/02_calculator/10_conversions/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
