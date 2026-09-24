# Wrap

The previous exercise wrapped a sentinel at construction time. This one
wraps on the *lookup* path — searching a list, missing, and reporting the
miss so precisely that callers can distinguish "not found" from every other
failure. Same verbs, new territory: the verb choice is the lesson.

## The search that can miss

```go
func FindTicket(ids []int, id int) (int, error) {
    for i, v := range ids {
        if v == id {
            return i, nil
        }
    }
    return -1, fmt.Errorf("ticket %d: %w", id, ErrNotFound)
}
```

Two conventions share this small function. The index pair `(-1, err)`:
`-1` is never a valid index, so it unambiguously means "no position," while
the error carries *why*. And the loop itself — `for i, v := range ids`,
both variables this time, because position is the product being sold.

## `%w` versus its impostors

Three verbs format errors. Only one preserves identity:

```go
fmt.Errorf("ticket %d: %w", id, ErrNotFound) // wraps: Is() sees ErrNotFound
fmt.Errorf("ticket %d: %v", id, ErrNotFound) // flattens: message only, Is() blind
fmt.Errorf("ticket %d: %s", id, ErrNotFound) // flattens, explicitly textual
```

`%v` and `%s` render the sentinel's *text* into the message and discard the
link — output looks right, `errors.Is` fails, and the bug hides until the
one caller who branches on identity runs. The test here checks exactly that
branch:

```go
_, err = FindTicket([]int{10, 20}, 99)
if !errors.Is(err, ErrNotFound) {
```

A `%v` implementation prints a lovely message and fails this assertion.
When reviewing error code, the verb after the colon deserves the same
scrutiny as the logic around it.

## Unwrap is a chain, not a box

`errors.Is` doesn't compare — it *walks*. Each wrapped layer exposes its
inner error via `Unwrap()`, and `Is` descends until it finds the sentinel
or runs out of layers. Deep call stacks wrap at every level (`"handler:
service: store: not found"`), and identity survives the whole descent.
That property is what makes wrapping scale: add context freely at every
boundary, secure that detection at the top still works.

## Task

Complete `FindTicket` in `ticket.go` so it returns the index of `id`, or a
wrapped `ErrNotFound` when the id is absent.

## Check

```bash
go test ./exercises/05_ticket_v2/03_wrap/ -v
```
