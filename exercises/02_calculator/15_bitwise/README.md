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
