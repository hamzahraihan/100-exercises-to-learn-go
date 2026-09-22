---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Runes"
weight: 13
draft: false
---

# Runes

A Go `string` is bytes; a `rune` is one character. `"é"` is 2 bytes
(`len` reports 2) but a single rune. Indexing a string yields bytes —
ranging over it yields runes.

## Task

Fix `CountRunes` in `calc.go` so multibyte characters count once.

## Check

```bash
go test ./exercises/02_calculator/13_runes/ -v
```

---

*Source: `exercises/02_calculator/13_runes/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
