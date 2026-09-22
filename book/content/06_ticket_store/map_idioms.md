---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Map Idioms"
weight: 6
draft: false
---

# Map Idioms

A map indexes values by key for direct lookup. Reading uses the comma-ok
idiom, which separates "present with a value" from "absent", and removal
uses the `delete` builtin, which is safe even when the key is missing. Maps
have one sharp edge: a `nil` map reads fine but writing to it panics, so a
type designed to work in its zero value must allocate its map lazily on
first write with `make` before storing. `Registry` follows that pattern,
indexing tickets by title so the zero `Registry` is usable with no
constructor. This is the Go equivalent of the keyed-registry step in the
ticket-store section of the Rust course this section is adapted from.

## Task

Complete `Put` and `Lookup` in `store.go` so stored tickets are retrievable
by title, missing titles report absent via the comma-ok idiom, and the test
passes:

```go
func (r *Registry) Put(t Ticket) {
	// ...
}

func (r *Registry) Lookup(title string) (Ticket, bool) {
	// ...
}
```

## Check

```bash
go test ./exercises/06_ticket_store/06_map_idioms/ -v
```

---

*Source: `exercises/06_ticket_store/06_map_idioms/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
