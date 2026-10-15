# Custom Errors and errors.As

`errors.Is` matches identity — is this error *that* sentinel? `errors.As` matches structure — is this error (or anything it wraps) *of this type*? The difference matters when you need data out of the error, not just a yes/no.

## The limit of sentinels

```go
var ErrNotFound = errors.New("not found")
// caller knows it's not found, but which ID?
return ErrNotFound // data lost
```

Sentinels are identity tags. They carry no payload. When the error needs context — which field failed, what value was rejected, which ID was missing — a sentinel is the wrong tool.

## Custom error types

```go
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

Any type with an `Error() string` method satisfies `error`. Adding fields makes the error a structured value, not a string.

## errors.As — type extraction through the chain

```go
var ve *ValidationError
if errors.As(err, &ve) {
    fmt.Println("bad field:", ve.Field)
}
```

`errors.As` walks the unwrap chain looking for an error assignable to `*ValidationError`. On match it fills `ve` and returns `true`. It works through any number of `%w` wraps:

```go
err := fmt.Errorf("create ticket: %w", &ValidationError{"title", "empty"})
var ve *ValidationError
errors.As(err, &ve) // true — unwraps one layer
```

## Is vs As — decision rule

- Known sentinel, branch on which one → `errors.Is`
- Structured error, need fields → `errors.As`
- Both in the same error chain: both work simultaneously

<details>
<summary>Hint</summary>

Define `ValidationError` with `Field` and `Message` string fields. Make `Error()` format them with `fmt.Sprintf`. In `NewTicket`, return a `*ValidationError` (wrapped with `%w` in `fmt.Errorf`) when title or description is empty. In `ExtractField`, call `errors.As` targeting `*ValidationError` and return its `Field`.

</details>

## Task

In `validate.go`:

1. Add `Field string` and `Message string` fields to `ValidationError`; implement `Error() string` to return `"<Field>: <Message>"`.
2. Complete `NewTicket(title, description string) (Ticket, error)` — return a `*ValidationError` (wrapped with `fmt.Errorf("new ticket: %w", ...)`) when `title` is empty (`field="title"`) or `description` is empty (`field="description"`). Otherwise return a valid `Ticket`.
3. Complete `ExtractField(err error) (string, bool)` — use `errors.As` to extract the `Field` from a `*ValidationError` anywhere in the error chain.

## Check

```bash
go test ./exercises/05_ticket_v2/06_errors_as/ -v
```
