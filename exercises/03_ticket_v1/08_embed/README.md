# Embedding

Go has no inheritance. No parent classes, no `extends`, no method override
chains to untangle at midnight. And yet real domains share structure —
every ticket, user, and comment in a system wants an ID and timestamps.
Without inheritance, how do fields get reused? By *containing* instead of
*descending*.

## Composition, literally

```go
type Meta struct {
    ID int
}

type Ticket struct {
    Meta        // no field name — the type name is the field
    Title string
}
```

A field declared with only a type is **embedded**. `Ticket` doesn't inherit
from `Meta`; it *contains* a `Meta`, addressable as `tk.Meta`. So far,
ordinary nesting. The magic is what Go adds on top: **promotion**. The
embedded struct's fields (and methods) are lifted to the outer type, so
both spellings work:

```go
tk.Meta.ID // the honest path: through the embedded struct
tk.ID      // the promoted shortcut: same field, less typing
```

Promotion is convenience, not copying — one field, two addresses. And when
two embedded structs promote the *same* name, Go refuses to guess: the
shortcut errors on ambiguity while the explicit paths keep working. The
language promotes eagerly but disambiguates never.

## Building through the nesting

Construction stays explicit — promotion doesn't extend to literals:

```go
func NewTicketWithID(id int, title string) Ticket {
    return Ticket{Meta: Meta{ID: id}, Title: title}
}
```

`Ticket{ID: id, ...}` would not compile; the literal names the embedded
*field* (`Meta`), and the `Meta` literal names *its* field (`ID`). Reading
uses shortcuts, writing uses full paths. Asymmetric, but honest: every
level of structure is visible exactly where values are created.

## Embedding versus naming

When should the field be anonymous (`Meta`) versus named (`Meta Meta`)?
Embed when the inner type's API should feel like the outer type's own —
`tk.ID` reading as the ticket's identity, methods promoting upward to
enrich the ticket. Name the field when the relationship deserves daylight:
`ticket.Author.Name` says *a ticket has an author*; `ticket.Name` would lie
about whose name it is.

This same mechanism scales one level further: embedded *interfaces* lend
whole method sets, and the interfaces section puts that to work. For now,
one struct inside another, promoted fields, explicit construction — Go's
entire answer to reuse, with no hierarchy attached.

## Task

Implement `NewTicketWithID` in `ticket.go` so it returns a `Ticket` with the
given id and title, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/08_embed/ -v
```
