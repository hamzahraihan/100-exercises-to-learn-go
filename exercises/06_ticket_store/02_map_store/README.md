# Map Store

Same store, new engine. The slice version found tickets by scanning every
element — correct, and increasingly slow. This exercise keeps the identical
`Add`/`Get` shape and swaps the storage for a map: one key leading straight
to its value, no scanning involved.

## Indexing instead of scanning

```go
type Store struct {
    tickets map[int]Ticket
    nextID  int
}

func (s *Store) Add(title string) int {
    if s.tickets == nil {
        s.tickets = make(map[int]Ticket)
    }
    id := s.nextID
    s.nextID++
    s.tickets[id] = Ticket{ID: id, Title: title}
    return id
}
```

`map[int]Ticket` reads "from int keys to Ticket values." Insertion is
assignment-shaped: `s.tickets[id] = ...`, no append, no result to keep.
Where the slice grew by reallocation, the map absorbs entries into its own
structure — lookups then land directly instead of walking.

## The nil map's split personality

Maps have one sharp edge, and this exercise is built around it. A `nil` map
**reads** perfectly — lookups return zero values, `range` runs zero times.
But **writing** to a `nil` map panics on the spot.

Since the test again declares `var s Store` with no constructor, `Add` must
allocate on first use:

```go
if s.tickets == nil {
    s.tickets = make(map[int]Ticket)
}
```

`make` builds a ready map; the `nil` check keeps it lazy, so stores that
never receive a ticket never allocate one. This is the zero-value design
habit applied to containers: usable at declaration, equipped on demand.
Slices needed no such care (append handles `nil`); maps demand one `if`.
Asymmetric, worth memorizing: *nil slices append, nil maps must be made.*

## Asking with comma-ok

```go
func (s *Store) Get(id int) (Ticket, bool) {
    tk, ok := s.tickets[id]
    return tk, ok
}
```

Map lookup has a second return built in: `v, ok := m[k]`. `ok` reports
presence; `v` is the value or the zero value when absent. No sentinel, no
error — presence is data, not failure, so it travels as a boolean in the
paired-returns rhythm you've used since divmod.

Compare the two engines side by side. Slice `Get` scans and synthesizes
`false`; map `Get` indexes and reports the language's own verdict. Same
signature, same tests, different physics — and the reason large stores
index instead of scan.

## Task

Complete `Add` and `Get` in `store.go` so added tickets are retrievable by
id, unknown ids report missing via the comma-ok idiom, and the test passes.

## Check

```bash
go test ./exercises/06_ticket_store/02_map_store/ -v
```
