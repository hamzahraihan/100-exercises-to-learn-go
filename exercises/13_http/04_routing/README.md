# Routing with ServeMux

One handler answers one thing. Real servers answer dozens — collections,
single items, health checks — and something must look at each request and
choose. That something is the **router**: patterns in, handlers out, with
a `404` for everything unmatched.

## Patterns with verbs and holes

```go
func NewMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /tickets", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "list")
    })
    mux.HandleFunc("GET /tickets/{id}", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "one:"+r.PathValue("id"))
    })
    return mux
}
```

Since Go 1.22, `ServeMux` patterns speak HTTP natively: `"GET /tickets"`
matches the method *and* the path (a `POST` there falls through to 404),
and `{id}` captures a path segment into a name the handler reads back with
`r.PathValue("id")`. No third-party router, no regexes, no code
generation — the standard library routes modern REST shapes out of the box.

Unmatched requests get `404 Not Found` automatically. That's not laziness;
it's the only sane default. An explicit catch-all would conflate "no such
endpoint" with application errors, and clients distinguish those cases to
decide between "fix the URL" and "retry later."

## The test's own helper

```go
func get(t *testing.T, mux http.Handler, path string) *httptest.ResponseRecorder {
    t.Helper()
    rec := httptest.NewRecorder()
    mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
    return rec
}
```

Testing a router means calling `ServeHTTP` directly — the mux *is* an
`http.Handler`, so the test drives it like any handler, no socket, no
server process. And the helper wears `t.Helper()` (the testing section's
mark of infrastructure), so failures point at the calling assertion, not
the shared plumbing. Two requests, two bodies — `"list"` for the
collection, `"one:7"` with the captured segment substituted. Routing
assertions read the bodies; the mux did the choosing.

## Task

Fill in `NewMux` in `router.go`:

```go
func NewMux() *http.ServeMux {
    // ...
}
```

Register `"GET /tickets"` to write `"list"` and `"GET /tickets/{id}"` to
write `"one:"` plus the path value. The stub registers nothing, so the test
fails on the bodies.

## Check

```bash
go test ./exercises/13_http/04_routing/ -v
```
