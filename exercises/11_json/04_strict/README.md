# Strict Decoding

By default `json.Unmarshal` ignores unknown fields, which is friendly to
forward compatibility — old code keeps working when a server adds a field —
but it also hides typos like `"titel"`. A `json.Decoder` with
`DisallowUnknownFields` flips the tradeoff: every field must match the
struct, so misspellings and version skew surface as errors. The stub uses
plain `Unmarshal`, so the unknown-field case passes silently and the test
fails.

## Task

Fill in `UnmarshalStrict` in `ticket.go`:

```go
func UnmarshalStrict(data string) (Ticket, error) {
	// ...
}
```

Decode with a `json.Decoder` over `data` with `DisallowUnknownFields`
enabled instead of `json.Unmarshal`.

## Check

```bash
go test ./exercises/11_json/04_strict/ -v
```
