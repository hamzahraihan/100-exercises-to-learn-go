# Graceful Shutdown

Starting servers is covered everywhere; *stopping* them is the neglected
art. `Close` drops every connection mid-sentence — in-flight requests die,
responses truncate, clients see errors for work that nearly finished. In
production, restarts happen constantly: deploys, scaling, crashes upstream.
`Shutdown` is how a server lands the plane instead of ejecting the crew.

## Drain, don't drop

```go
func ShutdownGracefully(srv *http.Server, timeout time.Duration) error {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    return srv.Shutdown(ctx)
}
```

`Shutdown` stops accepting *new* connections, waits for in-flight requests
to complete, then returns. In-flight work lands safely; only the
yet-unstarted is refused. But "wait" needs a bound — one stuck handler
must not hold a deploy hostage forever — so the wait travels with a
timeout context: `context.WithTimeout` derives a context that cancels
itself after the duration, `defer cancel()` releases its resources, and
`Shutdown` abandons stragglers when the fuse burns out. Bounded patience,
stated up front.

The context lesson's vocabulary, now aimed at the server itself: same
`Done` mechanics, same `WithTimeout` derivation, one level up the stack.
Contexts started per-request; this one spans the shutdown. The pattern
transfers unchanged because it was designed to.

## Testing the drain with nothing in it

```go
srv := &http.Server{}
ShutdownGracefully(srv, 2*time.Second) // want nil
```

A server with no listeners, no connections, nothing in flight — shutdown
returns at once, nil error. The test proves the machinery runs, not that
it drains under load (deterministic beats realistic for unit scope). Real
drain behavior gets exercised in integration, with live requests and
measured timing; here the contract under test is narrower and exact: the
function completes, returns nil, honors its signature.

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
