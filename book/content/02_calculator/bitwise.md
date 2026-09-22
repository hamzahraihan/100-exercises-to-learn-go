---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Bitwise operators"
weight: 15
draft: false
---

# Bitwise operators

`<<` shifts bits, `&`/`|`/`^` combine them. With `iota`, each constant in
a block gets a successive value, so `1 << iota` builds power-of-two flags
that combine with `|` and test with `&`.

## Task

Fix `Has` in `calc.go` using `&`.

## Check

```bash
go test ./exercises/02_calculator/15_bitwise/ -v
```

---

*Source: `exercises/02_calculator/15_bitwise/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
