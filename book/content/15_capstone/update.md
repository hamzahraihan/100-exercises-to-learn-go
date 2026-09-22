---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Update a Ticket"
weight: 3
draft: false
---

# Update a Ticket

`PUT /tickets/{id}` means full replace in this API: the client sends the new
representation and the server swaps it in, keeping the path id. The handler
parses the id (`400` when malformed), decodes the body, validates it (empty
title: `400`), then calls `store.Update` (`404` when absent) and encodes the
stored ticket with `200`. Order matters — validate before touching the store
so bad input never corrupts good data. The store already works, the route is
registered, the handler is the exercise.

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

---

*Source: `exercises/15_capstone/03_update/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
