# Testing Handlers with httptest

Every handler test so far followed the same ritual — recorder, request,
call, inspect — without anyone naming it. This exercise names it. The
logic under test is deliberately trivial (`"ping"` in, `"pong"` out) so
that one hundred percent of your attention lands on the *machinery*.

## The whole ritual, labeled

```go
rec := httptest.NewRecorder()                            // fake writer
req := httptest.NewRequest(http.MethodGet, "/ping", nil) // fake request
Ping(rec, req)                                           // direct call, no network
```

`httptest.NewRequest(method, target, body)` builds a `*http.Request` —
method constant (never a raw string when the constant exists), path, and
an optional body reader (`nil` for GETs, `strings.Reader` for POSTs, as
the JSON lesson showed). `httptest.NewRecorder()` returns a
`ResponseRecorder` standing in for the connection. The handler runs as an
ordinary function call. Then the inspection:

| Inspector           | What it shows                  |
|---------------------|--------------------------------|
| `rec.Code`          | captured status code           |
| `rec.Header().Get`  | a response header value        |
| `rec.Body.String()` | captured body as a string      |

Status, headers, body — the complete observable surface of a response, all
three assertable without a socket. Every handler test in this section uses
some subset of this table; keep it as the pattern reference and write new
tests by choosing rows.

## What the recorder forgives (and doesn't)

The recorder is lenient where the network is strict: it captures multiple
`WriteHeader` calls instead of ignoring all but the first, and it never
times out, disconnects, or half-closes. That leniency is exactly right for
unit tests — deterministic, fast, focused on *your* logic — and exactly why
integration tests with real servers still exist elsewhere. Know which game
you're playing: recorder for handler logic, `httptest.NewServer` (which the
client lesson introduces) for full-stack behavior.

`Ping` itself is one line — write `"pong"` — because this exercise grades
fluency with the ritual, not creativity with handlers. Liveness probes
like this anchor real deployments; here it anchors your testing vocabulary.

## Task

Fill in `Ping` in `ping.go`:

```go
func Ping(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Write `"pong"` to `w`. The stub writes nothing, so the test fails on the
body.

## Check

```bash
go test ./exercises/13_http/05_httptest/ -v
```
