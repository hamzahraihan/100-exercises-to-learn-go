---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Saturating Add"
weight: 9
draft: false
---

# Saturating Add

Saturating arithmetic clamps at the type's maximum instead of wrapping:
`SaturatingAdd(200, 100)` on `uint8` gives 255, not 44. Since Go wraps
silently, you detect the wrap (`sum < a`) and substitute the max value.

## Task

Fix `SaturatingAdd` in `calc.go` to return 255 when the sum overflows.

## Check

```bash
go test ./exercises/02_calculator/09_saturating/ -v
```

---

*Source: `exercises/02_calculator/09_saturating/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
