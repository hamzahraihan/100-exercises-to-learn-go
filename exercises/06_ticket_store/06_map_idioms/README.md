# Map Idioms

The map store indexed tickets by id. This one indexes by title — and in
doing so becomes a catalog of every map habit worth owning. Three idioms,
one small type, zero constructors.

## Lookup: the comma-ok, again

```go
func (r *Registry) Lookup(title string) (Ticket, bool) {
    tk, ok := r.byTitle[title]
    return tk, ok
}
```

`v, ok := m[k]` separates "present with a value" from "absent" — the same
form the map store used for ids, now keyed by string. It works on a `nil`
map too: lookups against nothing return zero values with `ok == false`, no
panic, no special case. Reading never needs permission; only writing does.

## Removal: delete and move on

```go
delete(r.byTitle, title)
```

`delete` removes a key, and deleting a missing key is a silent no-op — no
error, no boolean, no panic. Maps are the rare Go structure where removal
can't fail loudly enough to matter, so the language declines to report it
at all. (This exercise doesn't ask for removal, but every registry grows
one eventually; now you know it's one line with no edge cases.)

## Storage: allocate on first write

The sharp edge, same as the map store: assigning into a `nil` map panics.
And the test again declares its registry bare:

```go
var r Registry
r.Put(Ticket{Title: "a"})
```

So `Put` equips the map lazily:

```go
func (r *Registry) Put(t Ticket) {
    if r.byTitle == nil {
        r.byTitle = make(map[string]Ticket)
    }
    r.byTitle[t.Title] = t
}
```

`make` builds the map on first write; reads before that moment hit the
`nil`-map path and report absent. Zero value usable, first write
self-sufficient, no constructor anywhere. Three exercises now — slice
store, map store, registry — all honor the same contract: `var x T` works.
That repetition is the point. Zero-value usability isn't a trick per type;
it's the house style, and these stores are its showroom.

Note the method set: pointer receivers throughout (`*Registry`), because
`Put` mutates — the setters lesson, applied to containers without a second
thought.

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
