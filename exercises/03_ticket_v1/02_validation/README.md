# Validation

Methods on a struct use a receiver:

```go
func (t Ticket) Validate() error {
    // return nil when valid, an error otherwise
}
```

The `error` type is a built-in interface; the standard way to build one
is `errors.New` or `fmt.Errorf`. A common convention is to return `nil`
for success and a non-nil error describing what is wrong.

## Task

Implement `Validate` in `ticket.go`: return an error when `Title` is empty
or `Description` is shorter than 10 characters, and `nil` otherwise, so the
test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/02_validation/ -v
```
