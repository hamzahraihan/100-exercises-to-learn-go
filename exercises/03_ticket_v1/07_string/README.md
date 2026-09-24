# String

`fmt.Println(tk)` prints something today. Something *useless* — field
braces and raw values, the debug dump every Go struct gets for free. This
exercise teaches your type to introduce itself properly, and smuggles in
your first interface along the way.

## A method the whole standard library calls

The `fmt` package prints any value. But before falling back to the dump, it
asks one question: does this value have a `String() string` method?

```go
// Syntax: no parameters, one string result
func (t Ticket) String() string {
    return fmt.Sprintf("Ticket: %s", t.Title)
}
```

Define that method (adding `import "fmt"` first — the packages lesson's
unused-import rule runs in reverse: use it or lose the build) and every — `fmt.Sprint`, `fmt.Println`,
`%v` and `%s` verbs, log lines, test failures — uses your text instead.
One method, total coverage. The test proves it through the public door,
`fmt.Sprint(tk)`, never calling `String()` directly: the point is that
*callers don't need to know the method exists*.

## Satisfaction without declaration

Here's the smuggled lesson: nothing registered `Ticket` anywhere. There is
no `implements` keyword in Go, no list of conformed protocols. Defining the
method **is** the conformance — `Ticket` now satisfies the `fmt.Stringer`
interface implicitly, by behavior alone.

```go
// fmt.Stringer, the entire contract:
type Stringer interface {
    String() string
}
```

This is Go's signature idea, and the interfaces section will spend a whole
chapter on it. For today, absorb the shape: small interfaces (one method is
ideal), satisfied silently, discovered by usage. Code written against
`fmt.Stringer` accepts your ticket without either side knowing the other's
name in advance.

## Verbs and one trap

`fmt.Sprintf` builds the text with verbs: `%s` for strings, `%q` for
quoted strings, `%d` for integers, `%v` for "whatever fits." Test output
throughout this course is `Sprintf` doing quiet work — every `t.Fatalf`
message you read was formatted, not concatenated.

And the trap, so you never fall in: inside `String()`, never format the
ticket itself with `%v` or `%s`:

```go
// DON'T: String calls Sprintf calls String calls Sprintf...
return fmt.Sprintf("%v", t)
```

Infinite recursion — `fmt` sees the `String` method, calls it, which calls
`fmt`, forever. Format the *fields* (`t.Title`), never the whole value.
One day this will save you from a stack overflow wearing a friendly face.

## Task

Implement `String` in `ticket.go` so the printed form includes the ticket's
title (the `fmt` package is available for formatting), and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/07_string/ -v
```
