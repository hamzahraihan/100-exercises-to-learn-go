# Update a Ticket

Creation appends, retrieval reads — both leave existing data alone.
`PUT /tickets/{id}` does not: it *replaces* a stored ticket with the
client's new representation. Replacement is the most dangerous write,
because every step can corrupt what a previous step validated. The
ordering below is the whole lesson.

## Validate before touching

```go
func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        http.Error(w, "bad id", http.StatusBadRequest)
        return
    }
    var t Ticket
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    if t.Title == "" {
        http.Error(w, "title required", http.StatusBadRequest)
        return
    }
    updated, ok := s.store.Update(id, t)
    if !ok {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updated)
}
```

Parse, decode, validate, *then* store — in that order, no exceptions.
`Update` runs last because everything before it can reject: malformed id,
broken JSON, empty title each bounce before a single stored byte moves.
Validate-then-mutate keeps bad input from corrupting good data halfway
through; the reverse order would require rollback logic nobody wants to
write. Cheap checks first, irreversible acts last — a principle that
outlives HTTP entirely.

`PUT` here means full replace: the body is the new ticket, complete, with
the path id authoritative (the store stamps `t.ID = id` regardless of what
the body claims — URL wins over payload when they disagree, always).

## Three tests, three verdicts

```go
PUT /tickets/1  {"title":"B",...} // 200 + title echoed — the happy path
PUT /tickets/99 {...}             // 404 — well-formed, absent
PUT /tickets/1  {"title":""}      // 400 — present, invalid
```

Each rejection gets its own assertion, and the suite reads as the
handler's contract in miniature: success echoes, absence 404s, invalidity
400s. Note the 404 case sends `{}` — valid JSON, decodable, but for a
ghost id. The 400 case sends a real id with an empty title. Every
combination of *routable / decodable / valid / present* earns separate
coverage, because each exercises a different branch — and untested
branches are where replacement bugs hibernate.

## Task

Fill in `handleUpdate` in `update.go`:

```go
func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Parse the path id, decode and validate the replacement, then update or
answer `404`. The stub answers `501`, so the tests fail on status and body.

## Check

```bash
go test ./exercises/15_capstone/03_update/ -v
```
