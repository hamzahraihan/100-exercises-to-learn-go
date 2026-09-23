# Struct

Don't jump ahead!
Complete the previous ticket exercise before starting this one. The exercise
for this section is located in `exercises/03_ticket_v1/01_struct`.

You've been passing titles and descriptions around as loose strings. This
section bundles them into a single domain type. We won't cover everything
structs can do — *just enough* to define one, build one, and control who
can touch its fields. One step at a time!

## Structs

A `struct` groups related fields into one named type. Each field has a name
and a type:

```go
type Ticket struct {
    Title       string
    Description string
}
```

`Ticket` is now a type you can use like `string` or `int` — as a parameter,
a return value, or a field of another struct. Later sections will attach
behavior to it; for now it only carries data.

## Creating values

The usual way to build a struct is a literal with **named fields**:

```go
tk := Ticket{Title: "Fix bug", Description: "Crash on login"}
fmt.Println(tk.Title) // "Fix bug"
```

Named fields are order-independent: `Ticket{Description: "...", Title: "..."}`
means the same thing. Prefer them always — adding or reordering fields
later won't break the code that uses them.

You can also write the literal positionally, without field names:

```go
b := Ticket{"Fix bug", "Crash on login"}
```

This compiles, but it is fragile: reorder the fields in the struct
declaration and every positional literal silently means something else.
It is considered idiomatic to avoid positional literals for structs with
more than one or two fields.

Fields you omit get their **zero value** (`""` for strings, `0` for
numbers, `nil` for slices and maps, `false` for booleans):

```go
c := Ticket{Title: "Fix bug"} // c.Description == ""
```

The zero value is always valid Go — no initialization step is required
before using a struct. You will rely on this property throughout the
course.

## Exported fields

Go has no access-modifier keywords. Visibility is decided by the first
letter of the name: a field starting with a capital letter is **exported**
(usable from other packages), anything else is unexported:

```go
type Ticket struct {
    Title string // exported: readable as ticket.Ticket.Title
    notes string // unexported: only usable inside package ticket
}
```

The same rule covers functions, types, and methods. `NewTicket` can be
called from another package; `newTicket` cannot. When you want outside code
to read a value but not write it freely, keep the field unexported and
expose controlled access through functions — the pattern the next few
exercises build up.

## Constructors are just functions

Go has no constructors. The convention is a plain function, usually named
`New` followed by the type name, that returns a ready-to-use value:

```go
func NewTicket(title, description string) Ticket {
    return Ticket{Title: title, Description: description}
}
```

There is nothing special about the name — the compiler doesn't treat
`NewTicket` differently from any other function. What makes it a
"constructor" is the agreement among Go programmers that `NewX` builds an
`X`. Later, when tickets need validation, this same function will return
`(Ticket, error)` instead.

## Reading test failures

The test for this exercise prints the whole struct on failure:

```go
t.Fatalf("got %+v", tk)
```

`%+v` shows field names alongside values. If you see `got {Title: Description:}`,
you returned the zero value — the function ran but never assigned the
parameters to the fields. Check that each parameter actually lands in its
field.

## Common mistakes

- **Returning an empty literal:** `return Ticket{}` compiles and the test
  failure (`got {Title: Description:}`) tells you exactly what happened.
- **Assigning the wrong way round:** `Ticket{Title: description, ...}`
  compiles because both are strings. The test output shows the swap.
- **Reaching for unexported fields across packages:** code outside
  `package ticket` cannot name unexported fields at all. That is the
  encapsulation mechanism — use the constructor instead.

## Task

Populate `NewTicket` in `ticket.go` to return the struct with the given
title and description so the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/01_struct/ -v
```
