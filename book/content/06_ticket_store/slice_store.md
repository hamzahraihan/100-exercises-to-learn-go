---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Slice Store"
weight: 1
draft: false
---

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

---

*Source: `exercises/06_ticket_store/01_slice_store/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
