---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Integers"
weight: 1
draft: false
---

# Integers

Go has `int`, `int8/16/32/64`, `uint`, `uint8/16/32/64`, `uintptr`.
Unlike Rust there is no `u128`, and `int` is 32 or 64 bits depending on platform.
Go never implicitly converts between integer types: `uint32(m)` is required.

## Task

Use `multiplier` (a `uint8`) in the computation. Convert it to `uint32` first.

## Check

```bash
go test ./exercises/02_calculator/01_integers/ -v
```

---

*Source: `exercises/02_calculator/01_integers/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
