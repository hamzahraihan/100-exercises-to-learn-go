# Marshaling Tickets

Somewhere your ticket must become bytes — to cross HTTP, to rest on disk.
`encoding/json` performs that translation, and for structs it needs almost
no instruction: point at the value, collect the bytes. The subtleties hide
in what gets encoded and how the test proves it.

## Values out, bytes back

```go
import "encoding/json"

func MarshalTicket(t Ticket) (string, error) {
    out, err := json.Marshal(t)
    if err != nil {
        return "", err
    }
    return string(out), nil
}
```

`json.Marshal` returns `[]byte`, not a string — JSON is bytes on the wire,
and Go keeps the distinction honest. `string(out)` converts for callers
who traffic in text. Errors are rare here (only unencodable values like
channels or cyclic structures fail) but never impossible, so the signature
carries `(string, error)` and failures propagate like any other.

Two silent rules govern what appears. **Only exported fields encode** —
unexported ones vanish without warning, the capital-letter rule from the
structs lesson now operating at the serialization boundary. And **tags
rename**: `json:"title"` puts `Title` on the wire as `"title"`,
`json:"description,omitempty"` additionally drops the field when empty.
The test input carries a full description, so `omitempty` stays quiet this
round; the tags lesson ahead makes it sing.

## Round-trip testing

The test never compares your bytes to a golden string:

```go
var got Ticket
if err := json.Unmarshal([]byte(out), &got); err != nil {
    t.Fatalf("output is not valid JSON: %v (output = %q)", err, out)
}
if !reflect.DeepEqual(got, in) {
    t.Fatalf("round-trip = %+v, want %+v", got, in)
}
```

Decode your output back into a `Ticket` and compare *values*. Key order,
spacing, escaping choices — all irrelevant, as they should be: the test
asserts meaning, not formatting. First it demands validity (malformed JSON
fails with the offending output quoted for inspection), then equivalence.
Round-trip assertions are the standard shape for serialization tests —
steal it for every encoder you write, in either direction.

## Task

Fill in `MarshalTicket` in `ticket.go`:

```go
func MarshalTicket(t Ticket) (string, error) {
    // ...
}
```

Use `json.Marshal` on `t` and return the bytes as a string.

## Check

```bash
go test ./exercises/11_json/01_marshal/ -v
```
