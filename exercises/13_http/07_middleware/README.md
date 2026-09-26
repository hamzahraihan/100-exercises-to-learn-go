# Middleware

Logging, auth, header-tagging, panic recovery — every request passes
through concerns that aren't any single handler's business. Duplicating
them into each handler rots within weeks. **Middleware** factors them out:
handlers wrapped in handlers, each layer adding behavior around the next.

## Handlers all the way down

```go
// Syntax: a function from handlers to handlers
func WithHeader(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Course", "go")
        next.ServeHTTP(w, r)
    })
}
```

The shape `func(http.Handler) http.Handler` is the whole pattern. Take the
next handler, return a new one that does its work and delegates. The
returned closure *is* a handler (via the `HandlerFunc` adapter from the
first lesson), so it chains arbitrarily:

```go
WithHeader(logging(mux))
```

Read inside out: requests hit `WithHeader` first, then `logging`, then the
mux. Each layer runs its pre-work, calls `next`, and optionally does
post-work after it returns. Forget the `next` call and the chain stops
dead — requests answered (or more often, *unanswered*) mid-stack. A
middleware that never delegates isn't layered behavior; it's a wall.

## Headers before delegation

This exercise's layer tags every response:

```go
w.Header().Set("X-Course", "go")
next.ServeHTTP(w, r)
```

Set-then-delegate, in that order — headers mutate freely before the inner
handler writes, and the freezing law (status lesson) still governs the
final bytes. The test asserts both halves independently: the header exists
*and* the inner body `"ok"` survived, proving the layer added without
obstructing. A middleware test always checks both directions — what it
changed, and what it preserved.

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
