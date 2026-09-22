---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Factorial"
weight: 5
draft: false
---

# Factorial

Go has a single loop keyword, `for`, which handles every loop shape.
A counting loop looks like `for i := 1; i <= n; i++ { ... }`.
Start the accumulator at 1 so that `Factorial(0)` is 1.

## Task

Fix `Factorial` in `calc.go` to return `n!` using a `for` loop.

## Check

```bash
go test ./exercises/02_calculator/05_factorial/ -v
```

---

*Source: `exercises/02_calculator/05_factorial/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
