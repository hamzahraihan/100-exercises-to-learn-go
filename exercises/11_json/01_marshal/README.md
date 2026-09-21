# Marshaling Tickets

`encoding/json` turns Go values into JSON bytes with `json.Marshal`, and the
struct tags on `Ticket` control the wire names: `json:"title"` renames a
field, while `json:"description,omitempty"` drops it when empty. The test
decodes your output back into a `Ticket` and compares values, so key order
never matters. The stub returns `""` for everything, so the test fails on
the decode.

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
