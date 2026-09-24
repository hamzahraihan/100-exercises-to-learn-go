# Setters

Methods with value receivers read. But a ticket's title needs *changing* —
and the previous lesson's receiver hands every method a disposable copy.
Watch what happens if `SetTitle` takes that copy and writes to it: the
assignment lands, the method returns, the copy evaporates. The caller's
ticket never moves. This exercise fixes that with one character.

## The copy in the mirror

```go
func (t Ticket) SetTitle(title string) {
    t.Title = title // writes to the copy; caller sees nothing
}
```

Every method call copies the receiver into `t`, exactly as if you'd passed
it as an argument. For readers like `Validate` that's harmless. For
writers it's fatal — mutations die with the copy. Go's answer isn't
reference semantics or annotations. It's a different receiver:

```go
// Syntax: pointer receiver — the star changes everything
func (t *Ticket) SetTitle(title string) {
    t.Title = title // writes through to the caller's ticket
}
```

`t *Ticket` receives the *address* of the caller's value, so `t.Title =
title` lands in the original. Same call syntax at the call site —
`tk.SetTitle("new")` — completely different fate for the data. Go
automatically takes the address when you call a pointer method on an
addressable value, so callers never juggle `&` themselves.

## Which receiver, when

The rule of thumb fits in one line: **pointer receivers mutate, value
receivers observe.** A few refinements you'll internalize with practice:

- If *any* method on a type needs a pointer receiver, give them all
  pointer receivers. Mixed receiver sets confuse readers and complicate
  interfaces (a value doesn't implement pointer-receiver methods — the
  interfaces section will make this concrete).
- Small, immutable-feeling types (time instants, coordinates) often stay
  all-value. Tickets, with their changing titles, point the other way.
- The test calls `tk.SetTitle("new")` on a plain `Ticket` variable and
  then reads `tk.Title`. Addressable variable plus pointer method: Go
  bridges the gap silently, mutation persists.

So the fix is genuinely minimal — a star and an assignment:

```go
func (t *Ticket) SetTitle(title string) {
    t.Title = title
}
```

One character of syntax, and the mirror becomes a window.

## Task

Implement `SetTitle` in `ticket.go` so it changes the ticket's title and
the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/04_setters/ -v
```
