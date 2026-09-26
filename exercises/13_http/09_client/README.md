# Calling Other Servers

Handlers answer; clients ask. Sooner or later your server calls another —
microservices, webhooks, third-party APIs — and the standard library's
client side is as unadorned as its server side: one function for the
simple case, visible machinery for everything else.

## The four-line fetch

```go
func GetBody(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
```

`http.Get` performs the request and returns the response — headers
arrived, body *streaming*. That streaming is why the next line is
non-negotiable: `defer resp.Body.Close()` immediately after the error
check. An unclosed body leaks the connection; the transport can't reuse
it, and under load the client starves on its own leftovers. Close-first
is the closest thing HTTP clients have to a sacrament — defer it before
reading a single byte, every time, no exceptions.

Then `io.ReadAll` drains the stream (the reader lesson, over a socket
instead of a string) and the bytes cross to text. Each error returns as-is:
network failures are the caller's business, reported, not hidden.

## Timeouts belong on requests

`http.Get` has no timeout — a hung server hangs you forever. Production
code sets deadlines, and the context lesson provides the vocabulary:

```go
req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := http.DefaultClient.Do(req)
```

Same four-line shape downstream; the deadline rides in via context, and
cancellation propagates through the socket automatically. Default to this
form the moment a call leaves your laptop: unbounded waits are outages
wearing patience's clothes.

## A server that isn't

```go
srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("data"))
}))
defer srv.Close()
```

Unlike the recorder (which calls handlers directly), `httptest.NewServer`
binds a real local port and serves real HTTP — full stack, real sockets,
still no internet. `srv.URL` is the address; `defer srv.Close()` tears it
down. Deterministic, hermetic, fast: client tests get a server Orson Welles
would envy — real enough to believe, fake enough to control.

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
