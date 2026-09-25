# Pagination

Lists grow; screens don't. Somewhere between a thousand tickets and a user
interface sits a window: show *these* ten, starting *there*. This exercise
cuts that window — and discovers that the callers choosing its edges can't
be trusted.

## The window

```go
// Five tickets, IDs 1..5:
Page(ts, 1, 2) // IDs [2 3] — skip 1, take 2
```

`offset` says where the window starts, `limit` caps how many it holds. The
happy path is one slice expression: `ts[offset : offset+limit]`, trimmed
when the window overhangs the end. Simple — until the inputs arrive from
HTTP query strings, config files, and other programs written by people in
a hurry.

## The enemies

The test sends four attacks, and each must return empty — never panic:

```go
Page(ts, 99, 10)  // offset past the end
Page(ts, 0, 0)    // non-positive limit
Page(ts, -1, -5)  // negatives everywhere
```

A raw `ts[offset:offset+limit]` dies on all three: index out of range is a
panic, the emergency brake from the panics lesson, and slicing past the end
pulls it. Callers you don't control means inputs you don't control, so the
function clamps *first* and slices *second*:

```go
func Page(ts []Ticket, offset, limit int) []Ticket {
    if offset < 0 {
        offset = 0
    }
    if limit <= 0 || offset >= len(ts) {
        return nil
    }
    end := offset + limit
    if end > len(ts) {
        end = len(ts)
    }
    return ts[offset:end]
}
```

Negatives pinned to zero, hopeless cases bounced early (the guard-clause
habit from the calculator section, now defending a slice expression), the
end trimmed to reality. Every bound is proven safe *before* the brackets
run. Read any slicing code with this suspicion: if the indexes aren't
demonstrably inside `[0, len]`, the panic is a matter of input, not of if.

## Empty is empty

The test asserts emptiness with `len(got) != 0` — deliberately loose. `nil`
and zero-length non-nil slices both qualify, so the implementation needn't
prefer one. This is Go's pragmatic streak: emptiness is a *length*, not an
identity. Return `nil` for the degenerate cases and move on; callers
ranging or measuring can't tell the difference, and no allocation is spent
manufacturing an empty non-nil slice nobody can distinguish.

## Task

Complete `Page` in `store.go` so it returns up to `limit` tickets starting
at `offset`, clamps out-of-range and non-positive inputs to an empty (len 0)
result without panicking, and the test passes:

```go
func Page(ts []Ticket, offset, limit int) []Ticket {
    // ...
}
```

## Check

```bash
go test ./exercises/06_ticket_store/05_pagination/ -v
```
