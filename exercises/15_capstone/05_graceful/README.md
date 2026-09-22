# Graceful Shutdown

`Close` drops connections immediately; `Shutdown` drains them: it stops
accepting new requests, waits for in-flight ones to finish, then returns —
bounded by a timeout via `context.WithTimeout` so a stuck handler cannot hang
your deploy forever. A `&http.Server{}` with no listeners shuts down at once,
which is what the test exercises.

## Task

Fill in `ShutdownGracefully` in `shutdown.go`:

```go
func ShutdownGracefully(srv *http.Server, timeout time.Duration) error {
	// ...
}
```

Derive a timeout context and pass it to `srv.Shutdown`. The stub returns an
error, so the test fails on the return value.

## Check

```bash
go test ./exercises/15_capstone/05_graceful/ -v
```
