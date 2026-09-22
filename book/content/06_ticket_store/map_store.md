---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Map Store"
weight: 2
draft: false
---

# Map Store

A map gives indexed lookup: one key leads straight to its value instead of
scanning every element. `Store` keeps a `map[int]Ticket` plus a `nextID`
counter, mirroring the previous slice exercise but swapping the storage.
Two map rules matter here. First, reading a missing key is fine, and the
comma-ok form (`v, ok := m[k]`) tells present apart from absent. Second, a
`nil` map reads fine but writing to it panics, so `Add` must allocate the
map on first use — and since the test uses `var s Store`, that lazy setup
keeps the zero value usable with no constructor. This is the Go equivalent
of the keyed ticket-store step in the Rust course this section is adapted
from.

## Task

Complete `Add` and `Get` in `store.go` so added tickets are retrievable by
id, unknown ids report missing via the comma-ok idiom, and the test passes.

## Check

```bash
go test ./exercises/06_ticket_store/02_map_store/ -v
```

---

*Source: `exercises/06_ticket_store/02_map_store/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
