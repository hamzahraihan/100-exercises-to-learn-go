# Validation

A struct can hold anything — including nonsense. `Ticket{Title: "",
Description: "x"}` compiles happily, and every function downstream pays for
it. This exercise attaches judgment to the type: a method that inspects a
ticket and reports what's wrong with it.

## Behavior attaches with receivers

Functions live at package level. **Methods** live on a type, declared with
a **receiver** between `func` and the name:

```go
// Syntax: func (<recv> <Type>) <Name>(<params>) <results>
func (t Ticket) Validate() error {
    // t is the ticket this was called on
}
```

`t` is just a parameter with a privileged position — the value before the
dot in `tk.Validate()`. A value receiver like this one gets a *copy*, which
is fine: validation only reads. (Mutation needs a different receiver, two
exercises from now.)

## Errors are values

Go functions don't throw. They **return** problems as ordinary values of
the built-in `error` type — and the universal convention is `nil` for
success:

```go
if err := tk.Validate(); err != nil {
    // something is wrong; err describes it
}
// err == nil: proceed, the ticket is sound
```

So `Validate` returns `nil` for a good ticket and a populated error
otherwise. Build simple errors with `errors.New`, formatted ones with
`fmt.Errorf`. Both are already in the standard library; no custom types
needed today.

## One rule per branch, named in the message

The test demands two rejections — empty title, short description — and it
reads the message:

```go
if err := NewTicket("ok", "x").Validate(); err == nil ||
    !strings.Contains(err.Error(), "description") {
```

An error nobody reads is a log line; an error that names its field is a
diagnosis. Mention *which* rule broke (`"title must not be empty"`,
`"description must be at least 10 characters"`), and callers can act
without opening your source. Write each rule as its own early return —
flat branches, one complaint each — instead of accumulating cleverness:

```go
func (t Ticket) Validate() error {
    if t.Title == "" {
        return errors.New("title must not be empty")
    }
    if len(t.Description) < 10 {
        return errors.New("description must be at least 10 characters")
    }
    return nil
}
```

The shape should look familiar: it's the guard-clause habit from the
calculator section, now returning explanations instead of zeros.

## Task

Implement `Validate` in `ticket.go`: return an error when `Title` is empty
or `Description` is shorter than 10 characters, and `nil` otherwise, so the
test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/02_validation/ -v
```
