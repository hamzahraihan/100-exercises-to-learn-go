# Delete a Ticket

The smallest handler with the sharpest contract. `DELETE /tickets/{id}`
parses, removes, and answers — no body in, no body out, one status that
means exactly what happened. Fewer lines than any sibling, and the one
most worth reading twice.

## 204 means "gone, nothing to say"

```go
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        http.Error(w, "bad id", http.StatusBadRequest)
        return
    }
    if !s.store.Delete(id) {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
```

`204 No Content` is deletion's native tongue: success, *by definition*
bodiless. Write the header and return — no JSON, no confirmation text, no
`Content-Type`. Clients distinguish "deleted" from "here's the deleted
thing" by this code, and caches invalidate on it. The `400`/`404` guards
above it are the familiar vocabulary (malformed id, absent ticket),
unchanged since the get-one lesson — new verbs, same grammar.

## Proof by resurrection failure

```go
// DELETE /tickets/1 → 204
// GET /tickets/1    → 404 (was 200 before the delete)
```

The test doesn't trust the status alone — it reads the ticket *back*.
Before deletion the id answers; after, it's gone. Read-after-delete is
the only proof deletion happened rather than being merely claimed: a
handler returning 204 without removing anything would pass a status-only
assertion and fail this one. State-changing endpoints earn state-verifying
tests. Always confirm through a second operation that the world actually
changed.

Notice the setup gift: this directory wires the *working* `handleGet`
beside the broken `handleDelete`. Stuck? Read the neighbor — same file
family, same store, same conventions, one handler already correct. Real
codebases teach the same way: adjacent working code is the best
documentation for the piece you're writing. Study the getter as your
model, then write its mirror.

## Task

Fill in `handleDelete` in `delete.go`:

```go
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Parse the path id, delete or answer `404`, and write `204` with no body. The
stub answers `501`, so the tests fail on status.

## Check

```bash
go test ./exercises/15_capstone/04_delete/ -v
```
