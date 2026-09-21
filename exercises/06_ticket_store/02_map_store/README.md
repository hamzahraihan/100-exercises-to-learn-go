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
