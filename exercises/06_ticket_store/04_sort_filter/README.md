# Sort and Filter

Two transforms power most list code: reorder what you have, keep what you
want. Both look trivial. Both hide the same trap from the previous lesson —
the input's backing array, shared unless you intervene. These two functions
learn opposite disciplines for it.

## Sort a copy, not the caller

```go
func SortedByID(ts []Ticket) []Ticket {
    out := slices.Clone(ts)
    slices.SortFunc(out, func(a, b Ticket) int {
        return a.ID - b.ID
    })
    return out
}
```

`slices.SortFunc` orders with your comparator: negative when `a` goes
first, zero when tied, positive when `b` does. (`a.ID - b.ID` ascends;
swap the operands and it descends.) But sorting rearranges *in place* —
run it on `ts` directly and the caller's order is destroyed as a side
effect. The name promises a sorted *result*, so the function clones first
and sorts the clone. Clone-then-sort is the whole discipline: never reorder
what you don't own.

The test feeds `{3, 1, 2}` and expects `{1, 2, 3}` — via `reflect.DeepEqual`,
order-sensitive this time. Sorting is one of the few operations where
sequence matters to equality, and the test leans on exactly that.

## Filter by appending

```go
func OpenOnly(ts []Ticket) []Ticket {
    var out []Ticket
    for _, t := range ts {
        if !t.Closed {
            out = append(out, t)
        }
    }
    return out
}
```

Filtering builds: walk the input, append each match into a fresh result.
`var out []Ticket` starts `nil`, and appending to `nil` works — the
zero-value habit from the slice store, earning interest. The input is never
written, only read, so no clone is needed; the fresh `out` array guarantees
independence by construction.

Note what's *not* here: no deletion from the original, no in-place
compaction. Removing from a slice while ranging over it is a classic source
of skipped elements. Build the new list instead of surgically editing the
old one — simpler to read, impossible to skip with.

## Two transforms, one contract

Sorted and filtered results share a promise: *the input comes back
untouched*. Different mechanisms (clone-then-sort, build-fresh), same
guarantee. Callers should never wonder whether a function borrowed their
array permanently. When your own functions take slices, decide the aliasing
story up front and make the name tell it: `SortedByID` returns something
new, and its body had better agree.

## Task

Complete `SortedByID` and `OpenOnly` in `store.go` so one returns tickets
ordered by `ID` ascending and the other keeps only tickets that are not
closed, and the tests pass:

```go
func SortedByID(ts []Ticket) []Ticket {
    // ...
}

func OpenOnly(ts []Ticket) []Ticket {
    // ...
}
```

## Check

```bash
go test ./exercises/06_ticket_store/04_sort_filter/ -v
```
