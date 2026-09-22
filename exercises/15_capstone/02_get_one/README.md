# Get One Ticket

Collection routes address the whole list; a single resource hangs off its id
in the path: `GET /tickets/{id}`. The `ServeMux` pattern captures the segment
and the handler reads it back with `r.PathValue("id")`, converts it with
`strconv.Atoi` (a non-number is the client's fault: `400`), looks it up in
the store (`404` when absent), and encodes it as JSON with `200`. The store
already works, the route is registered, the handler is the exercise.

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
