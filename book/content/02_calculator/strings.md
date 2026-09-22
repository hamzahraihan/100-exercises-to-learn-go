---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Strings"
weight: 14
draft: false
---

# Strings

Strings support `==` and `+`, but you cannot assign to an index:
`s[0] = 'H'` does not compile. Build a new string with slicing and
concatenation instead. The `strings` package has the rest.

## Task

Fix `ToUpperFirst` in `calc.go`. Keep `""` mapping to `""`.

## Check

```bash
go test ./exercises/02_calculator/14_strings/ -v
```

---

*Source: `exercises/02_calculator/14_strings/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
