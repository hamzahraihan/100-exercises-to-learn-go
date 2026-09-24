# Stringer

You've already done this exercise once — you just didn't know its name.
Back in the ticket section you gave `Ticket` a `String() string` method and
`fmt` started using it. That wasn't a trick of the `fmt` package. It was
your first interface, hiding in plain sight.

## Naming the concept

An **interface** is a set of method signatures — a description of behavior,
with no implementation attached:

```go
// Syntax: type <Name> interface { <methods> }
type Stringer interface {
    String() string
}
```

Any type with a `String() string` method *is* a `Stringer`. Nothing to
declare, nowhere to register. `Order` below qualifies the moment its method
exists:

```go
type Order struct {
    ID   int
    Item string
}

func (o Order) String() string {
    return fmt.Sprintf("Order %d: %s", o.ID, o.Item)
}
```

Functions written against the interface — `fmt.Sprint`, `fmt.Println`,
anything taking a `fmt.Stringer` — accept `Order` without either side
knowing the other's name in advance. The ticket lesson showed the effect;
this lesson names the cause.

## Satisfaction is silent

Most languages make conformance explicit: a declaration, a keyword, a list
maintained beside the type. Go's bet is that this bookkeeping costs more
than it saves. Interfaces stay small (one method is ideal), types satisfy
them by accident as much as by design, and old types retroactively satisfy
new interfaces nobody imagined when they were written.

The test asserts through the public door, as before:

```go
if got := fmt.Sprint(o); !strings.Contains(got, "7") || !strings.Contains(got, "book") {
```

`fmt.Sprint` takes `any` and asks each value for a `String` method at print
time. Your implementation decides the text; the test demands both facts —
id *and* item — appear in it. A `String` that prints only the id would
compile, satisfy the interface, and still fail: conformance is necessary,
content is graded.

## Task

Implement `String` in `order.go` so the returned string contains both the
order id and the item name, and the test passes.

## Check

```bash
go test ./exercises/04_interfaces/01_stringer/ -v
```
