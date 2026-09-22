---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "JSON In, JSON Out"
weight: 5
draft: false
---

# JSON In, JSON Out

A JSON API handler has three moves: decode the request body with
`json.NewDecoder(r.Body).Decode(&v)`, validate what you got (here: the title
must be non-empty, answered with `400 Bad Request`), then encode the reply
with `json.NewEncoder(w).Encode(v)`. Set
`w.Header().Set("Content-Type", "application/json")` and call `WriteHeader`
before encoding — headers and status freeze once the first body bytes flow.

## Task

Fill in `CreateTicket` in `api.go`:

```go
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

Decode a `Ticket`, reject an empty title with `400`, otherwise echo the
ticket back as JSON with `201`. The stub writes nothing, so the tests fail
on status, content type, and body.

## Check

```bash
go test ./exercises/13_http/05_json_api/ -v
```

---

*Source: `exercises/13_http/05_json_api/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
