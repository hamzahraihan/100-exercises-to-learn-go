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
