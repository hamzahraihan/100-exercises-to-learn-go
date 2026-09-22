---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Slice Functions (Clone)"
weight: 3
draft: false
---

# Slice Functions (Clone)

A slice header is a small descriptor — pointer, length, capacity — over a
shared backing array. Assigning one slice to another copies the header, not
the elements, so both headers still point at the same array: writing through
one alias is visible through the other. `len` reports how many elements are
addressable, `cap` how many the array can hold before `append` must allocate
a new one. A true clone needs a fresh array with the elements copied over;
the standard library offers `slices.Clone` by name for exactly this, and the
manual equivalent is an append-into-empty or copy-into-made slice. This is
the Go equivalent of the cloning step in the ticket-store section of the Rust
course this section is adapted from.

## Task

Complete `CloneTickets` in `store.go` so the result equals the input but
shares no backing array with it, and the test passes:

```go
func CloneTickets(ts []Ticket) []Ticket {
	// ...
}
```

## Check

```bash
go test ./exercises/06_ticket_store/03_slices_fn/ -v
```

---

*Source: `exercises/06_ticket_store/03_slices_fn/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
