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
