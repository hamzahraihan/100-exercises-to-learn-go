# Value Receivers

The setters lesson ended with a rule: pointer receivers mutate. This one is
its mirror image — and its warning label. Sometimes the *point* is to leave
the original untouched.

## Change without changing

```go
func (t Ticket) Renamed(newTitle string) Ticket {
    t.Title = newTitle
    return t
}
```

Same assignment as a setter, opposite fate. The value receiver copies the
ticket, the copy gets the new title, and the copy comes home as the return
value. The caller's original never moves:

```go
orig := Ticket{Title: "a", Description: "long enough description"}
got := orig.Renamed("b")
// got.Title == "b", orig.Title == "a" — both true at once
```

The test pins both halves: the result carries the new title *with the
description preserved*, and `orig` still reads `"a"`. Note the second
assertion's quiet demand — `Renamed` must copy the *whole* ticket, not
construct a fresh one from the title alone. Starting from `t` and changing
one field gets this right; building `Ticket{Title: newTitle}` from scratch
would silently drop the description.

## When copies win

Setters mutate in place; `Renamed` forks history. Prefer the fork when:

- **The old value must survive.** Drafts, undo stacks, before/after
  comparisons — anywhere two versions coexist.
- **The value is small.** Two strings copy cheaply. A ten-megabyte buffer
  would argue otherwise.
- **Callers chain.** `tk.Renamed("b").Renamed("c")` reads left to right
  because each step returns the next value. Pointer setters return nothing
  and break chains.

This is the copy-out style: pure function manners on a struct. No aliasing
surprises downstream, no "who else holds this pointer" audits. The price is
a copy per call — usually negligible, occasionally (huge structs, hot
loops) worth measuring rather than assuming.

Read the two lessons together and you hold the complete receiver doctrine:
`*Ticket` to change the thing, `Ticket` to derive from it. The interfaces
section will add the one remaining wrinkle — which receiver set a value
actually implements.

## Task

Implement `Renamed` in `ticket.go` so it returns a copy of the ticket with
the new title while leaving the original unchanged, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/05_receivers/ -v
```
