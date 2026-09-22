---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Sum To"
weight: 6
draft: false
---

# Sum To

Go has only `for`: there is no `while` keyword. The same C-style header
`for i := 1; i <= n; i++` covers counting loops, and `for condition { }`
covers while-style loops.

## Task

Fix `SumTo` in `calc.go` to return `1 + 2 + ... + n` (0 for `n <= 0`).

## Check

```bash
go test ./exercises/02_calculator/06_sum_to/ -v
```

---

*Source: `exercises/02_calculator/06_sum_to/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
