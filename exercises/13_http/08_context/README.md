# Contexts and Cancellation

Every request carries a silent passenger: a `context.Context` that dies
when the request does — client disconnected, deadline exceeded, server
shutting down. Work that ignores it keeps computing for ghosts: burning
CPU on answers nobody will read, holding connections nobody owns. This
exercise teaches the polite alternative.

## Listening for goodbye

```go
func Greet(ctx context.Context, name string) (string, error) {
    select {
    case <-time.After(10 * time.Millisecond):
        return "hi " + name, nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}
```

The `select` shape should ring bells from the timeout lesson — two
futures, whichever answers first. Here the racers are asymmetric: real
work (a simulated 10ms lookup) against `ctx.Done()`, the channel that
closes when cancellation arrives. Work wins, greeting returns. Goodbye
wins, `ctx.Err()` explains why (`context.Canceled`, or
`context.DeadlineExceeded` for timeouts) — an error shaped for callers,
not a boolean shaped for guessing.

In handlers, the context arrives on the request (`r.Context()`), already
wired to the connection's fate. Plumb it down to anything slow — database
calls, downstream fetches, simulated lookups — and every layer gets the
same choice: finish promptly, or stop promptly. Contexts flow *down* call
stacks as the first parameter (convention, and linters enforce it);
values flow back *up* as returns. Never store a context in a struct, never
pass `nil` where one belongs.

## Testing both timelines

```go
got, err := Greet(context.Background(), "bob") // never cancels: greeting
```

`context.Background()` is the root that outlives everything — production's
starting point, tests' neutral baseline. Against it, the lookup wins and
`"hi bob"` returns with nil error.

```go
ctx, cancel := context.WithCancel(context.Background())
cancel() // fire immediately: already dead on arrival
Greet(ctx, "bob") // must return the cancellation error
```

The canceled case pre-fires: `cancel()` runs *before* the call, so `Done()`
is already closed and the select takes the goodbye branch without waiting
out the 10ms. Deterministic — no sleeps, no races, no flakiness. Testing
cancellation means controlling time's arrow, and a pre-canceled context is
the simplest time machine available.

## Task

Fill in `Greet` in `greet.go`:

```go
func Greet(ctx context.Context, name string) (string, error) {
    // ...
}
```

Return `"hi "+name` after the simulated 10ms lookup, but return
`ctx.Err()` promptly if `ctx` is canceled first. The stub never watches the
context, so the test fails on the canceled case.

## Check

```bash
go test ./exercises/13_http/08_context/ -v
```
