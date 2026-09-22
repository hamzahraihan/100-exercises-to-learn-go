---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Middleware"
weight: 7
draft: false
---

# Middleware

Middleware is a function with the shape `func(http.Handler) http.Handler`:
it takes the next handler, returns a new one that does some work (logging,
headers, auth) and then calls `next.ServeHTTP(w, r)`. Handlers chain by
wrapping — `WithHeader(logging(mux))` runs the header step first, the
logging step second, and the mux last — so each layer must call `next` or
the chain stops there.

## Task

Fill in `WithHeader` in `mw.go`:

```go
func WithHeader(next http.Handler) http.Handler {
	// ...
}
```

Set the `X-Course: go` response header before delegating to `next`. The stub
delegates without setting it, so the test fails on the header.

## Check

```bash
go test ./exercises/13_http/07_middleware/ -v
```

---

*Source: `exercises/13_http/07_middleware/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
