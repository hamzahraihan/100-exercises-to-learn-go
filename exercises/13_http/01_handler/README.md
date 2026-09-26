# Your First Handler

Every web framework in every language funnels HTTP into a function with
two arguments: *answer here, question there*. Go strips the idea to its
minimum — a plain function with a fixed signature — and builds the entire
`net/http` ecosystem on top of it.

## A function that serves

```go
// Syntax: every handler takes these two parameters
func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello")
}
```

`http.ResponseWriter` is where the answer goes — and it's an `io.Writer`,
the contract from the io section, so `Fprint`, `WriteString`, and every
other writer tool works unchanged. `*http.Request` is where the question
lives: method, path, headers, body. One side produces, the other
describes; the handler sits between them, and that's the whole model.

The adapter that turns such functions into servable values is
`http.HandlerFunc` — a type with a `ServeHTTP` method that just calls the
function. Interface satisfaction again (the Stringer lesson's shape):
anything with `ServeHTTP(w, r)` *is* an `http.Handler`, and a bare
function qualifies through the adapter. You'll meet this direction —
function to interface — in routing and middleware next.

## No sockets in tests

```go
rec := httptest.NewRecorder()
Hello(rec, httptest.NewRequest(http.MethodGet, "/", nil))
```

Tests never bind a port. `httptest.NewRequest` builds a request value
(method, target, optional body) and `httptest.NewRecorder` stands in for
the network's end of the writer, capturing status, headers, and body for
inspection. The handler can't tell the difference — same signature, same
behavior, zero networking. `rec.Body.String()` then reads the captured
answer, and the assertion compares text, not sockets.

Note the unused parameter: `Hello` ignores `r` entirely, and Go doesn't
mind — unlike unused *variables*, unused *parameters* compile fine. (The
stub's comment warns that imports are stricter: import `fmt` and never
call it, and the build fails. Use what you import, starting now.)

## Task

Fill in `Hello` in `hello.go`:

```go
func Hello(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Write `"hello"` to `w`. The stub writes nothing, so the test fails on the
body.

## Check

```bash
go test ./exercises/13_http/01_handler/ -v
```
