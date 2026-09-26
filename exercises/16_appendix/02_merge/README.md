# Merge

The import lesson chose replace semantics deliberately: validated backups
become the entire state, never a blend. But some updates *are* blends — a
sync that overlays remote changes onto local ones, keeping both sides'
contributions. This exercise is the merge the course promised and
postponed: same ids overwritten, order preserved, newcomers appended.

## Overlay, don't append

```go
a := []Ticket{{ID: 1, Title: "A"}, {ID: 2, Title: "B"}}
b := []Ticket{{ID: 2, Title: "B2"}, {ID: 3, Title: "C"}}

Merge(a, b) // [{ID:1 Title:A} {ID:2 Title:B2} {ID:3 Title:C}]
```

Ticket 2 exists on both sides: `b`'s version wins. Ticket 1 exists only in
`a`: kept, in place. Ticket 3 exists only in `b`: appended, in `b`'s
order. The naive alternative — concatenate then deduplicate — scrambles
positions and buries the overwrite rule inside cleanup logic. Index first,
then overlay; the structure of the code mirrors the structure of the
decision.

## Index, overlay, preserve

```go
func Merge(a, b []Ticket) []Ticket {
    pos := make(map[int]int, len(a)) // id → index in out
    out := make([]Ticket, 0, len(a)+len(b))
    for _, t := range a {
        pos[t.ID] = len(out)
        out = append(out, t)
    }
    for _, t := range b {
        if i, ok := pos[t.ID]; ok {
            out[i] = t // known id: overwrite in place, order untouched
        } else {
            pos[t.ID] = len(out) // new id: claim a position, append
            out = append(out, t)
        }
    }
    return out
}
```

The map answers "have I seen this id, and where?" in one lookup; `out`
grows exactly once per distinct id thanks to the capacity hint. Overwrites
land by index, so `a`'s order survives contact with `b`. Newcomers append
in `b`'s order, so repeated merges stay deterministic. One pass per input,
zero re-scans — the linear-scan habit from the slice store, retired by the
map habit from the map store, in a single function.

Note the newcomer bookkeeping: claiming `pos[t.ID]` *before* appending
keeps the index truthful even if `b` repeats an id within itself (last
occurrence wins, consistently). Invariants maintained at every step beat
invariants restored at the end.

## Task

Fill in `Merge` in `merge.go` so same-ID tickets come from `b`, order
follows `a`, and `b`-only tickets append in `b`'s order:

```go
func Merge(a, b []Ticket) []Ticket {
    // ...
}
```

Index `a` by ID, walk `b` overwriting known ids and appending newcomers.

## Check

```bash
go test ./exercises/16_appendix/02_merge/ -v
```
