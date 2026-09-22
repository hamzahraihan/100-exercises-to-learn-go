---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Delete a Ticket"
weight: 4
draft: false
---

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

---

*Source: `exercises/15_capstone/04_delete/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
