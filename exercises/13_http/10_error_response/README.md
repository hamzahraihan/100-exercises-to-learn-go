# Structured Errors

Success responses vary per endpoint; failures shouldn't. A client parsing
ten different error shapes writes ten parsers and still misses the
eleventh. The industry answer — and this section's closing discipline — is
one envelope every failure wears: same keys, same content type, only the
status and message changing.

## The recipe, in order

```go
func WriteError(w http.ResponseWriter, code int, msg string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
```

Three steps, and the sequence is load-bearing. **Content type first** — so
clients parse instead of sniff. **Status second** — the freezing law from
the status lesson in its final appearance: headers lock on first body
write, so `WriteHeader(code)` must precede encoding. **Body last** — the
envelope `{"error":"nope"}`, encoded straight into the writer. Reverse any
two steps and something observable breaks: wrong type, frozen 200, or a
body without a verdict.

A `map[string]string` serves for a one-key envelope; a struct with tags
serves when the shape grows (`code`, `details`, trace ids). Either
encodes; the contract is the JSON shape, not the Go type behind it.

## Why envelopes win

Compare the alternatives. Plain-text errors (`http.Error` with a string)
can't carry structure — clients regex the message. Per-endpoint shapes
force per-endpoint parsers. One envelope means one parser, and the *status
code* carries severity while the *body* carries explanation. Machines
branch on numbers, humans read messages; the envelope serves both without
either tripping over the other.

The test asserts all three facets — status, content type, message presence —
because a structured error *is* three promises: the right code, the
parseable type, the human text. Miss any one and some consumer, machine or
human, stumbles. This mirrors the JSON-creation test's triple assertion
from the JSON API lesson: same discipline, failure-flavored.

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
