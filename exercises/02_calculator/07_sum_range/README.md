# Sum Range

Counting loops count. But most real loops walk through *things* — and for
that, Go gives `for` a second companion: `range`. This exercise also smuggles
in your first collection type. Two lessons, one stub.

## Your first slice

```go
[]int{1, 2, 3, 4} // a slice of ints: an ordered, growable sequence
```

A **slice** is Go's workhorse list: `[]int` reads "slice of int," and the
literal in braces holds the elements. Slices appear everywhere ahead —
test tables already used them anonymously — so get comfortable at a glance.
One property matters today: a `nil` slice (the zero value, no elements at
all) is perfectly valid, and ranging over it simply runs zero times. The
test's `SumRange(nil)` row proves it.

## Ranging

```go
// Syntax: for <index>, <value> := range <collection> { <body> }
func SumRange(vals []int) int {
    total := 0
    for _, v := range vals {
        total += v
    }
    return total
}
```

`range` hands you each element in order — no counters, no boundary
operators, no off-by-one. The index arrives too, first: `for i, v := range
vals`. Need only the values? The blank identifier `_` swallows the index,
the same `_` that discarded indexes in the workflow lesson and return
values in the panics test. One symbol, one job: *I know this is here, and
I'm choosing to ignore it.*

Compare with the counting loop from the last two exercises. The counter
version of this function needs `vals[i]`, a length, and a `<` that had
better be right. `range` deletes all three concerns. Prefer it whenever you
touch every element — which, honestly, is most loops.

## Task

Fix `SumRange` in `calc.go` to add all slice elements with `for ... range`.

## Check

```bash
go test ./exercises/02_calculator/07_sum_range/ -v
```
