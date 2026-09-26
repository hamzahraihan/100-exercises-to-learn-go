# Custom Unmarshal

The custom-marshaler lesson ended with an IOU: encoding learned to speak
`"open"`, decoding never did. An API that writes words but only reads
numbers is half a contract — so here comes the other half. Same pattern
as encoding, mirrored: a method the `encoding/json` package discovers on
its own.

## The mirror method

```go
// Syntax: pointer receiver — decoding MUTATES, so it must be *Status
func (s *Status) UnmarshalJSON(data []byte) error {
    var v string
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch v {
    case "open":
        *s = StatusOpen
    case "closed":
        *s = StatusClosed
    default:
        return fmt.Errorf("unknown status %q", v)
    }
    return nil
}
```

Three differences from `MarshalJSON`, each load-bearing. The receiver is a
**pointer**: decoding writes into the value, and a value receiver would
mutate a copy the caller never sees (the setters lesson, at the JSON
boundary). The input is raw `data` — unquote it first, here by decoding
into a plain string and letting malformed JSON fail naturally. And unknown
words **error** instead of guessing: inventing a status for `"bogus"`
would launder someone else's bug into your store, the same refusal the
encoder practices.

Once defined, the decoder calls it automatically for every `Status` field
— implicit satisfaction in the Stringer tradition, now on the way in. The
test exercises it through `json.Unmarshal` into a `Ticket`, never calling
the method directly: the decoder is the only client, the decoded value the
only contract.

## Symmetry, audited

With both methods defined, check the round trip: every word `MarshalJSON`
emits, `UnmarshalJSON` accepts — and nothing else. `"open"` ↔ `StatusOpen`,
`"closed"` ↔ `StatusClosed`, everything else an error in at least one
direction. Asymmetric pairs (`"pending"` accepted but never emitted, say)
are legitimate when deliberate and bugs when accidental; auditing the pair
together is the habit. Two methods, one vocabulary, zero drift — the
encoder and decoder now speak the same language because they share the
same switch, conceptually if not literally.

## Task

Fill in `UnmarshalJSON` in `status.go` so `"open"`/`"closed"` decode and
anything else errors:

```go
func (s *Status) UnmarshalJSON(data []byte) error {
    // ...
}
```

Unquote `data`, switch on the word, assign through `*s` or return an
error.

## Check

```bash
go test ./exercises/16_appendix/03_unmarshal/ -v
```
