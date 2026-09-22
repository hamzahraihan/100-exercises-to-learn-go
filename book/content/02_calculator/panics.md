---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Panics"
weight: 4
draft: false
---

# Panics

Go signals unrecoverable errors with `panic`. Calling `panic("message")` stops
normal execution and unwinds the stack; tests recover from it with `recover()`.
Unlike returning an error, a panic crashes the program unless recovered.

## Task

Fix `Divide` in `calc.go` to `panic("division by zero")` when `b == 0`,
otherwise return `a / b`.

## Check

```bash
go test ./exercises/02_calculator/04_panics/ -v
```

---

*Source: `exercises/02_calculator/04_panics/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
