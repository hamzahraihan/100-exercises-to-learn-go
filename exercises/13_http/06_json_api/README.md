# JSON In, JSON Out

Handlers so far wrote plain text. Real APIs trade JSON both ways — and a
creation endpoint performs the same three moves every time: decode the
request, judge it, encode the answer. This exercise is that choreography,
with every lesson of the section inside it.

## Three moves

```go
func CreateTicket(w http.ResponseWriter, r *http.Request) {
    var t Ticket
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    if t.Title == "" {
        http.Error(w, "title required", http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(t)
}
```

**Decode.** `json.NewDecoder(r.Body).Decode(&t)` streams the request body
straight into the struct — no intermediate bytes, no manual reading. The
pointer is load-bearing (the unmarshal lesson's `&`, unchanged). Malformed
JSON fails here, answered `400`: the client's payload, the client's fault.

**Judge.** Empty title fails validation — the constructor discipline from
the ticket sections, relocated to the HTTP border. Same `400`: well-formed
but unacceptable is still the client's mistake. Two different rejections,
one status: the code distinguishes them in logs, the client sees one
verdict.

**Encode.** Content type first, status second, body last — the freezing law
from the status lesson, now with all three steps visible. `Set` the header
(`application/json`, so clients parse instead of sniff), `WriteHeader`
(`201`, something now exists), then `Encode` the ticket straight into the
writer. Reverse the order and the status locks at 200 while the body
already flows — correct JSON, wrong verdict.

## Two tests, two stories

```go
// Happy: 201 + JSON content type + title echoed in the body
CreateTicket(rec, httptest.NewRequest(http.MethodPost, "/tickets", body))
// Sad: missing title → 400
```

The request body arrives as a `strings.Reader` — test doubles again, the
reader lesson's `strings.NewReader` feeding a handler instead of a
function. The happy test asserts all three response facets (status,
content type, body content); the sad test asserts the rejection code.
Three moves in, three facets out: this is the complete JSON-handler
shape, and every endpoint in the capstone will echo it.

## Task

Fill in `CreateTicket` in `api.go`:

```go
func CreateTicket(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Decode a `Ticket`, reject an empty title with `400`, otherwise echo the
ticket back as JSON with `201`. The stub writes nothing, so the tests fail
on status, content type, and body.

## Check

```bash
go test ./exercises/13_http/06_json_api/ -v
```
