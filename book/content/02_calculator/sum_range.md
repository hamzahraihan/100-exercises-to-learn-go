---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Sum Range"
weight: 7
draft: false
---

# Sum Range

`for ... range` iterates over slices, arrays, and maps. The form
`for _, v := range vals` skips the index with the blank identifier `_`
and binds each element to `v`. Ranging over a `nil` slice simply runs zero times.

## Task

Fix `SumRange` in `calc.go` to add all slice elements with `for ... range`.

## Check

```bash
go test ./exercises/02_calculator/07_sum_range/ -v
```

---

*Source: `exercises/02_calculator/07_sum_range/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
