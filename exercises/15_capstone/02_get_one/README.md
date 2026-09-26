# Get One Ticket

Collections list; individuals hide behind ids. `GET /tickets/{id}` names
one ticket in the path — and three different things can go wrong turning
that name into an answer. This handler's whole job is telling those three
apart with the right status each.

## Three failures, three codes

```go
func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        http.Error(w, "bad id", http.StatusBadRequest)
        return
    }
    t, ok := s.store.Get(id)
    if !ok {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(t)
}
```

Read the chain as a vocabulary lesson. `"abc"` as id: the *client*
malformed the request — `400`. Well-formed id, no such ticket: the
resource simply isn't there — `404`. Each status names *whose fault*,
which is what status codes are for: machines branch on them, and the
branching only works when handlers choose honestly.

`r.PathValue("id")` pulls the captured segment (the routing lesson's
`{id}` pattern, registered and waiting), and `strconv.Atoi` converts it —
the border-crossing lesson's converter, now parsing URLs instead of form
fields. Conversion failure is client error, full stop: the server's
contract promises numeric ids, and `"abc"` violates the caller's side.

## The store does its half

`s.store.Get(id)` returns `(Ticket, bool)` — the paired-returns rhythm,
unchanged since divmod. Present: encode with the implicit `200` (success
needs no announcement). Absent: `404`. The handler adds no logic between
store verdict and HTTP verdict; translation layers should be thin enough
to read at a glance, and this one is four lines.

The test walks both paths: seed one ticket directly into the store
(`s.store.Add`, bypassing HTTP — setup through the back door, assertions
through the front), fetch `/tickets/1` for `200` plus title, fetch
`/tickets/99` for `404`. Seeding via the store while asserting via HTTP
tests the handler in isolation from creation — each layer proved through
the layer above, a pattern worth stealing for your own services.

## Task

Fill in `handleGet` in `getone.go`:

```go
func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Parse the path id, answer `400` for a malformed id, `404` when the store has
no such ticket, otherwise the ticket as JSON with `200`. The stub answers
`501`, so the tests fail on status and body.

## Check

```bash
go test ./exercises/15_capstone/02_get_one/ -v
```
