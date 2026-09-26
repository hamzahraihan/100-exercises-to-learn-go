# Status Codes

A body without a status is a sentence without punctuation — technically
complete, practically ambiguous. Every HTTP response carries a three-digit
code telling the client *what kind* of answer this is, and Go defaults it
to `200 OK` unless you say otherwise. This exercise says otherwise.

## The number before the words

```go
func Create(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
}
```

Status codes group by century: `2xx` success, `3xx` redirection, `4xx`
client error, `5xx` server failure. Never write raw numbers — the `net/http`
package names them (`StatusCreated` is 201, `StatusNotFound` 404,
`StatusBadRequest` 400), and names document intent where digits invite
typos. `201` vs `200` looks like nothing; `StatusCreated` vs `StatusOK`
reads like the decision it is.

Creation answers `201 Created`, not `200 OK`. The distinction matters to
clients: `200` says "here's your answer," `201` says "something now exists
that didn't before." REST clients, caches, and retry logic all branch on
that difference. Choosing the precise code is part of the handler's job,
not decoration.

## Headers freeze on first write

One law governs the whole response, and this section will repeat it until
it's reflex: **`WriteHeader` must precede the first body write.** Headers
(including the status line) travel ahead of the body on the wire; once
bytes flow, the status is locked and later `WriteHeader` calls are silently
ignored. Set status, then write — never the reverse:

```go
w.WriteHeader(http.StatusCreated) // first: freeze the status...
fmt.Fprint(w, "...")              // ...then the body flows under it
```

The test reads the frozen value through the recorder:

```go
Create(rec, httptest.NewRequest(http.MethodPost, "/tickets", nil))
if rec.Code != http.StatusCreated {
```

Note the method: `POST`, the verb that *creates*. The stub writes nothing,
so the recorder keeps its default `200` — and the assertion fails on the
number alone, body untouched. Status is independently testable, and
independently meaningful.

## Task

Fill in `Create` in `create.go`:

```go
func Create(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Answer `http.StatusCreated`. The stub writes nothing, so the recorder keeps
the default `200` and the test fails on the status.

## Check

```bash
go test ./exercises/13_http/02_status/ -v
```
