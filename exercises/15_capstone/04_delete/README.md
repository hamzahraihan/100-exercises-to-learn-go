# Delete a Ticket

Deletion is the smallest handler with the sharpest contract: `DELETE
/tickets/{id}` parses the id (`400` when malformed), removes the row (`404`
when absent), and answers `204 No Content` — a status that by definition
carries no body, so the handler writes the header and returns. The proof is a
read-after-delete: `GET /tickets/{id}` must flip to `404`, which is why this
dir's `NewServer` also wires the working `handleGet` alongside the broken
`handleDelete` — study the working getter as your model.

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
