# Struct

Until now a ticket has been two loose strings drifting through function
signatures. That arrangement works for exactly one exercise longer. The task
in `exercises/03_ticket_v1/01_struct` gives those strings a home: a `Ticket`
type with a shape the compiler can check.

Structs repay deep study, but deep study can wait. Today is about three
skills only — declaring one, building one, and deciding which fields
outsiders may touch. Everything else structs can do will arrive, as always,
right when you need it.

## The problem

Every ticket needs a title and a description. So far that's meant passing
two strings everywhere:

```go
func PrintTicket(title, description string)
```

Two parameters is tolerable. Five is a mess, and nothing stops you from
swapping two same-typed arguments by accident. But how do we **combine**
these pieces into a single entity the compiler understands?

## Defining a `struct`

A `struct` defines a **new Go type** by bundling fields together:

```go
// Syntax: type <Name> struct { <field> <type> ... }
type Ticket struct {
    Title       string
    Description string
}
```

`Ticket` can now be used like `string` or `int` — as a parameter, a return
value, or a field of another struct. Later sections will attach behavior to
it; for now it only carries data.

## Defining fields

Each field is a name plus a type, one per line. Fields don't have to share
a type, as the `Counter` below shows:

```go
type Counter struct {
    Name  string
    Count int
}
```

Field order is your choice and carries no meaning — but keep related fields
together. Future readers (including you, next week) will thank you.

## Building one

You create a value with a struct literal, naming each field:

```go
// Syntax: <StructName>{ <field>: <value>, ... }
tk := Ticket{Title: "Fix bug", Description: "Crash on login"}
```

Named fields are order-independent and survive the struct gaining new
fields later. Prefer them — always.

You *can* omit the names and list values positionally:

```go
b := Ticket{"Fix bug", "Crash on login"}
```

This compiles, but it's a trap: reorder the fields in the declaration and
every positional literal silently means something else. Positional literals
are only acceptable for tiny structs that will never change. When in doubt,
name the fields.

Fields you leave out get their **zero value** — `""` for strings, `0` for
numbers, `nil` for slices and maps, `false` for booleans:

```go
c := Ticket{Title: "Fix bug"} // c.Description == ""
```

The zero value is always usable; no initialization ritual is required before
a struct works. You'll lean on this property for the rest of the course.

## Reading one

Fields are accessed with the `.` operator:

```go
// Field access
title := tk.Title
```

Assignment works the same way, provided you own the value:

```go
tk.Title = "Fix bug faster"
```

## Who can touch the fields

Go has no access-modifier keywords. Visibility comes from the first letter
of the name: a field starting with a capital is **exported** — readable and
writable from any package — while a lowercase field is visible only inside
its own package:

```go
type Ticket struct {
    Title string // exported: usable as ticket.Ticket.Title
    notes string // unexported: only usable inside package ticket
}
```

The rule covers everything with a name: types, functions, methods. `NewTicket`
can be called from another package; `newTicket` cannot.

Why does this matter already? Because it's the encapsulation mechanism for
everything ahead. When a field must stay valid — a status with only a few
legal values, a description with a length limit — you keep it unexported and
hand out controlled access through functions. The next exercises build
exactly that, one step at a time.

## `NewTicket`: constructors are just functions

Go has no constructors. The convention is an ordinary function — usually
called `New` followed by the type name — that hands back a ready-to-use
value:

```go
func NewTicket(title, description string) Ticket {
    return Ticket{Title: title, Description: description}
}
```

Nothing about the name is special; the compiler treats `NewTicket` like any
other function. What makes it a "constructor" is agreement among Go
programmers that `NewX` builds an `X`. When tickets gain validation rules
in the next exercises, this same function will grow into
`func NewTicket(title, description string) (Ticket, error)` — the signature
evolves, the call sites barely notice.

## Reading test failures

The test prints the whole struct when it fails:

```go
t.Fatalf("got %+v", tk)
```

`%+v` shows field names alongside values. If you see
`got {Title: Description:}`, you returned the zero value — the function ran
but never assigned its parameters to the fields. If the values are swapped,
you'll see that too. Read the struct, don't guess.

## Task

Populate `NewTicket` in `ticket.go` to return the struct with the given
title and description so the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/01_struct/ -v
```
