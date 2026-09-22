# Structured Errors

Clients parse errors best when every failure shares one envelope, e.g.
`{"error":"nope"}`. The recipe mirrors the JSON exercise: set
`Content-Type: application/json`, call `WriteHeader(code)` with the caller's
status, then encode the envelope — status before body, always, because the
first body write freezes the code at `200` if you forget.

## Task

Fill in `WriteError` in `errw.go`:

```go
func WriteError(w http.ResponseWriter, code int, msg string) {
	// ...
}
```

Answer `code` with a JSON body carrying `msg` under the `"error"` key. The
stub writes nothing, so the test fails on status, content type, and body.

## Check

```bash
go test ./exercises/13_http/10_error_response/ -v
```
