# Struct Tags: Hiding and Omitting

The marshaling lesson encoded every exported field under a renamed key.
Real wire formats are pickier: secrets must never leave the process,
sparse updates shouldn't ship oceans of empty strings. Both are handled
without touching the Go code that uses the struct — by annotations *on*
the fields.

## Metadata in backticks

```go
type Ticket struct {
    ID           int    `json:"id"`
    Title        string `json:"title"`
    Description  string `json:"description,omitempty"`
    Status       string `json:"status"`
    InternalNote string `json:"-"`
}
```

A **struct tag** is a string literal after the field — conventionally
`key:"value"` pairs inside backticks — read via reflection by libraries.
`encoding/json` claims the `json` key. Its anatomy: a wire name, then
comma-separated options. `json:"description,omitempty"` renames *and*
conditions; `json:"-"` means "never on the wire, in either direction."

## Hide the secrets

`InternalNote` lives only in memory — reviewer comments, internal flags,
anything a client must never see. The `"-"` tag excludes it from encoding
*and* decoding alike: marshal a ticket carrying `"leak"` and the output
contains no trace; unmarshal JSON naming it and the field stays empty.

The stub demonstrates the alternative — a hardcoded string with the secret
baked in — and the test calls it a leak to its face:

```go
if strings.Contains(out, "leak") {
    t.Fatalf("secret leaked into JSON: %s", out)
}
```

Hand-built JSON bypasses every tag, every rule, every safeguard. Real
marshaling enforces the annotations; string surgery ignores them. That gap
is the lesson: serialization logic belongs to the encoder, never to
concatenation, because only the encoder sees the tags.

## Omit the empties

`omitempty` drops a field holding its zero value — `""`, `0`, `false`,
`nil`. An empty `Description` vanishes instead of serializing as `"\"\""`,
keeping payloads lean and PATCH-style updates unambiguous (absent means
untouched, present means set). Reach for it on optional fields; leave it
off where empty-vs-absent carries meaning — a `0` balance differs from an
unknown one, and there a `*int` with `omitempty` (nil drops, zero stays)
says precisely that.

Round-trip closes the test as before: decode, compare, confirm the title
survived while the secret didn't. Tags shape both directions with one
annotation — declare once, enforced everywhere.

## Task

Fill in `MarshalTicket` in `ticket.go`:

```go
func MarshalTicket(t Ticket) (string, error) {
    // ...
}
```

Use `json.Marshal` so the struct tags take effect — the `json:"-"` tag keeps
`InternalNote` out of the output on its own.

## Check

```bash
go test ./exercises/11_json/03_tags/ -v
```
