# Unmarshaling and Validating

Decoding with `json.Unmarshal` never guarantees a sane value: missing fields
just stay at their zero values, so a missing `"title"` arrives as `""`
without an error. The usual pattern is unmarshal-then-validate — decode into
the struct, then check the fields your domain requires and return an error
for bad input. The stub skips both steps and returns an empty `Ticket`, so
the test fails on the assertion.

## Task

Fill in `UnmarshalTicket` in `ticket.go`:

```go
func UnmarshalTicket(data string) (Ticket, error) {
	// ...
}
```

Use `json.Unmarshal` to decode `data`, then reject an empty `Title` with an
error.

## Check

```bash
go test ./exercises/11_json/02_unmarshal/ -v
```
