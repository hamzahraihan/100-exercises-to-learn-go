---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Create and List"
weight: 1
draft: false
---

# Create and List

A tiny ticket API starts with two moves against one shared store: `POST
/tickets` decodes a JSON ticket, validates it (here: the title must be
non-empty, answered with `400 Bad Request`), stores it, and echoes it back
with `201 Created`; `GET /tickets` encodes every stored ticket with `200
OK`. Both handlers share the same `*Store`, and `NewServer` wires each
method+path pattern to its handler — the store already works, the wiring is
done, the two handlers are the exercise.

## Task

Fill in `handleCreate` and `handleList` in `server.go`:

```go
func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	// ...
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

Decode with `json.NewDecoder(r.Body).Decode(&v)`, reject an empty title with
`400`, otherwise `store.Add` and reply `201` with the stored ticket as JSON.
For the list, encode `s.store.List()` as JSON with `200`. Set
`w.Header().Set("Content-Type", "application/json")` and call `WriteHeader`
before encoding. The stubs answer `501`, so the tests fail on status.

## Check

```bash
go test ./exercises/15_capstone/01_store_server/ -v
```

---

*Source: `exercises/15_capstone/01_store_server/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
