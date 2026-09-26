# Custom Marshalers

Tags rename and omit, but they can't *translate*. This exercise's `Status`
is an `int` underneath — `0` and `1` — while the API speaks `"open"` and
`"closed"`. No tag bridges that gap. So the type takes over its own
rendering, and the encoder steps aside.

## Taking over the rendering

```go
// Syntax: method on the type, exact signature required
func (s Status) MarshalJSON() ([]byte, error) {
    switch s {
    case StatusOpen:
        return []byte(strconv.Quote("open")), nil
    case StatusClosed:
        return []byte(strconv.Quote("closed")), nil
    default:
        return nil, fmt.Errorf("unknown status %d", int(s))
    }
}
```

`json.Marshaler` is an interface with one method — `MarshalJSON() ([]byte,
error)` — and the encoder checks for it on *every* value it touches
(another implicit satisfaction, in the Stringer tradition: define the
method, join the club, no registration). Once `Status` implements it,
`json.Marshal` on the whole `Ticket` calls it automatically for the status
field. Nothing at the call site changes; `MarshalTicket` stays a plain
`json.Marshal`, and the custom rendering flows through.

`strconv.Quote` produces the quoted form with escaping handled — `"open"`
with real quote bytes, safe for any future word containing specials. And
the `default` returns an error rather than guessing: an unknown status is
a programming bug (the `Valid()` gatekeeper from the status lesson should
have caught it), and inventing a wire word for it would launder the bug
into someone else's parser.

## Symmetry is a choice

Notice what's missing: no `UnmarshalJSON`. This exercise encodes custom
words but decodes nothing — the reverse direction is a separate method
(`UnmarshalJSON` on `*Status`) with its own edge cases, and asymmetric
formats are common in practice (write `"open"`, accept `"open"` *and*
legacy `0`). Custom encoding without custom decoding is a legitimate
endpoint, not half a job. When the API later demands `"closed"` parsed
back, that's a new method, a new exercise — one direction at a time, as
always.

The test asserts through the wire format, not the method: marshal a
ticket, demand `"status":"open"` in the output. It doesn't call
`MarshalJSON` directly — because callers never do either. The encoder is
the only client, and the output is the only contract.

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
