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
