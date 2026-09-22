---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Variables"
weight: 2
draft: false
---

# Variables

Go offers two ways to declare a variable: `var x int = 1` states the type
explicitly, while `x := 1` lets the compiler infer it. The short form `:=`
only works inside functions.

## Task

Fix `Double` in `calc.go` to return twice `x`, using `:=` to declare the result.

## Check

```bash
go test ./exercises/02_calculator/02_variables/ -v
```

---

*Source: `exercises/02_calculator/02_variables/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
