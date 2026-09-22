---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Div Mod"
weight: 11
draft: false
---

# Div Mod

Go functions can return multiple values, so quotient and remainder come back
together: `func DivMod(a, b int) (int, int)`. Integer `/` truncates toward
zero and `%` gives the remainder with the sign of the dividend.

## Task

Fix `DivMod` in `calc.go` to return `a / b` and `a % b`.

## Check

```bash
go test ./exercises/02_calculator/11_divmod/ -v
```

---

*Source: `exercises/02_calculator/11_divmod/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
