# Your First Handler

In Go, an HTTP handler is any function with the signature
`func(http.ResponseWriter, *http.Request)` — the `http.HandlerFunc` type
adapts such a function so it can serve traffic. You write the response body
through the `ResponseWriter` (it is an `io.Writer`), and you read everything
about the request from `*http.Request`. Tests never bind a socket: they call
your handler directly with `httptest.NewRequest` to build a request and
`httptest.NewRecorder` to capture what you write.

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
