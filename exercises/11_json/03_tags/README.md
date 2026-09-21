# Struct Tags: Hiding and Omitting

Two tags do most of the shaping work: `json:"-"` keeps a field out of the
JSON entirely (handy for secrets like `InternalNote` that live only in
memory), and `,omitempty` drops a field when it holds its zero value, so an
empty `Description` disappears instead of serializing as `""`. The stub
returns a hardcoded string that leaks the secret, so the test fails on the
leak check.

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
