---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Floats"
weight: 12
draft: false
---

# Floats

`float64` holds fractional numbers. Converting after integer division keeps
the truncation: `float64(5/2)` is `2.0`, while `float64(5)/2` is `2.5`.
The `math` package (`math.Pi`, `math.Sqrt`) covers the rest.

## Task

Fix `Half` in `calc.go` so odd inputs keep their `.5`.

## Check

```bash
go test ./exercises/02_calculator/12_floats/ -v
```

---

*Source: `exercises/02_calculator/12_floats/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
