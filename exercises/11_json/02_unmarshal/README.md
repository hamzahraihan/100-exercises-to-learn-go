# Unmarshaling and Validating

Encoding trusts your structs. Decoding trusts the *network* — and the
network lies by omission. Missing fields, misspelled keys, half-formed
payloads: `json.Unmarshal` accepts them all without complaint, filling
what it recognizes and zeroing the rest. This exercise adds the skepticism
decoding requires.

## Decode needs an address

```go
var t Ticket
if err := json.Unmarshal([]byte(data), &t); err != nil {
    return Ticket{}, err
}
```

The `&` is load-bearing. Unmarshal *writes into* the value, so it needs
the address — passing `t` by value would decode into a copy the caller
never sees, compiling cleanly and achieving nothing. Pointer-to-struct is
the decode target, always; the compiler can't catch the omission because a
value is a perfectly legal (if useless) argument.

Malformed JSON errors here — truncated input, broken syntax — and those
errors propagate untouched. But well-formed JSON with *wrong content*
sails through: `{"id":1,"status":"open"}` decodes fine, leaving `Title` as
`""`. No error. No warning. A titleless ticket, born valid.

## Distrust, then verify

That silence is why decoding ends with validation — the constructor
discipline from the ticket sections, relocated to the border:

```go
if t.Title == "" {
    return Ticket{}, errors.New("title must not be empty")
}
return t, nil
```

Unmarshal-then-validate, in that order, every time: decode into the
struct, check the fields your domain requires, reject what fails. The test
pins both halves — the valid payload decodes with fields intact, the
titleless one earns its rejection. Never trust a decoded struct before its
check; the wire is the least supervised constructor your types will ever
meet.

Notice what's *not* validated: unknown fields. `{"title":"x","zzz":1}`
decodes happily, `zzz` discarded. Whether that leniency stands is a policy
decision — and the strict-decoding exercise two doors down exists to
revisit exactly it.

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
