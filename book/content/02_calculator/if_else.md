---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "If Else"
weight: 3
draft: false
---

# If Else

Go's `if/else` looks familiar but drops the parentheses around the condition,
and the braces are required even for one-line bodies: `if a >= b { ... } else { ... }`.

## Task

Fix `Max` in `calc.go` to return the larger of `a` and `b` with an `if/else`.

## Check

```bash
go test ./exercises/02_calculator/03_if_else/ -v
```

---

*Source: `exercises/02_calculator/03_if_else/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
