---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Get One Ticket"
weight: 2
draft: false
---

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

---

*Source: `exercises/15_capstone/02_get_one/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
