# Query Parameters

Paths name the resource; the `?` tail tunes the request. `GET
/tickets?status=open` asks for the collection *filtered* — same endpoint,
narrower answer. This exercise reads that tail, with one ambiguity to
resolve by design.

## The map behind the question mark

```go
func StatusParam(r *http.Request) string {
    if s := r.URL.Query().Get("status"); s != "" {
        return s
    }
    return "all"
}
```

`r.URL.Query()` parses the query string into a `url.Values` — a
`map[string][]string`, since `?tag=a&tag=b` repeats keys legitimately.
`.Get("status")` returns the first value for the key, or `""` when absent.

And there's the ambiguity: `Get` returns `""` both when the key is *missing*
(`/tickets`) and when it's *present-but-empty* (`/tickets?status=`). The
two situations are indistinguishable through `Get` alone (checking key
presence needs the map directly). So handlers pick a policy, and this one
treats both as "not given" — defaulting to `"all"`. Sensible here: an
empty filter and no filter want the same answer. When they wouldn't (say,
clearing a setting vs leaving it), read the map with the comma-ok idiom
instead and branch on presence.

## Tests build URLs, not servers

```go
r := httptest.NewRequest("GET", "/tickets?status=open", nil)
StatusParam(r) // "open"

r = httptest.NewRequest("GET", "/tickets", nil)
StatusParam(r) // "all"
```

The request target carries the query string verbatim — no URL-encoding
ceremony for simple values, no test server, no network. Both branches get
pinned: present value passes through, missing value defaults. Query
handling is pure string logic over a parsed map, and tests treat it that
way — function in, string out.

## Task

Fill in `StatusParam` in `filter.go`:

```go
func StatusParam(r *http.Request) string {
    // ...
}
```

Return the `"status"` query value, or `"all"` when it is missing. The stub
always returns `""`, so the test fails on the returned value.

## Check

```bash
go test ./exercises/13_http/03_query/ -v
```
