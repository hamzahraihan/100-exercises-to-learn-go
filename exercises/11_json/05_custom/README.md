# Custom Marshalers

When the default encoding is not the wire format you need, implement
`json.Marshaler`: a `MarshalJSON() ([]byte, error)` method on your type takes
over how values of that type render. Here `Status` is an `int` underneath,
but the API speaks `"open"`/`"closed"`, so the method maps each constant to
its quoted string and `MarshalTicket` picks it up automatically. The stub
method returns a TODO error and the encoder returns `""`, so the test fails
on the output check.

## Task

Fill in both functions in `ticket.go`:

```go
func (s Status) MarshalJSON() ([]byte, error) {
	// ...
}
```

```go
func MarshalTicket(t Ticket) (string, error) {
	// ...
}
```

Switch on `s` and return the quoted status word (`strconv.Quote` helps), then
encode the ticket with `json.Marshal`.

## Check

```bash
go test ./exercises/11_json/05_custom/ -v
```
