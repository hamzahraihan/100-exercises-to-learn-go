# Calling Other Servers

`http.Get(url)` performs an outgoing GET and returns a `*http.Response`
whose `Body` must always be closed — `defer resp.Body.Close()` right after
the error check, or connections leak. Read the bytes with `io.ReadAll` and
convert to a string. For requests that need timeouts or cancellation, build
them with `http.NewRequestWithContext` and run them with `http.DefaultClient.Do`.
The test serves a local `httptest.NewServer`, so no real network is involved.

## Task

Fill in `GetBody` in `getbody.go`:

```go
func GetBody(url string) (string, error) {
	// ...
}
```

Fetch `url` and return its body as a string. The stub returns `"", nil`, so
the test fails on the body.

## Check

```bash
go test ./exercises/13_http/09_client/ -v
```
