# Slice Store

A slice is Go's growable sequence: `append` adds elements, returning the
updated slice you must keep. Here `Store` pairs a `[]Ticket` with a `nextID`
counter, so each `Add` hands out a fresh id. Because the test declares the
store with `var s Store`, the zero value must work — a `nil` slice appends
fine, so no constructor is needed. To find a ticket, `Get` scans the slice
and reports whether the id was present. That scan is the Go equivalent of
the linear lookup in the ticket-store section of the Rust course this
section is adapted from.

## Task

Complete `Add` and `Get` in `store.go` so added tickets are retrievable by
id, unknown ids report missing, and the test passes.

## Check

```bash
go test ./exercises/06_ticket_store/01_slice_store/ -v
```
