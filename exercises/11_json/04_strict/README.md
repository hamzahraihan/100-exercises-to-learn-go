# Strict Decoding

`json.Unmarshal` greets unknown fields with a shrug: not in the struct?
Discarded, no error. That leniency is a deliberate policy — old clients
survive servers that grew new fields — but policies have victims. A typo'd
`"titel"` decodes into nothing, the title stays empty, and validation
blames the *content* for a *spelling* mistake three layers away. This
exercise flips the policy.

## Leniency is a choice, not a default of nature

Default decoding optimizes for evolution: producers may add fields freely,
consumers ignore what they don't know. The cost lands on typos and version
skew — misspellings vanish silently, and a client talking to a newer API
never learns what it's missing. Neither behavior is "correct"; each serves
a different trust relationship:

- **Lenient** (`Unmarshal`): public APIs, long-lived consumers, fields
  that come and go. Be liberal in what you accept.
- **Strict** (decoder + flag): configs, internal contracts, anything where
  a misspelled key means a misconfigured system. Reject what you don't
  recognize, loudly.

Configs are the canonical strict case. A `"titel"` in a config file that
silently becomes empty is an outage wearing a typo's clothes — precisely
the failure strictness exists to convert into a startup error.

## The stricter machine

```go
func UnmarshalStrict(data string) (Ticket, error) {
    var t Ticket
    dec := json.NewDecoder(strings.NewReader(data))
    dec.DisallowUnknownFields()
    if err := dec.Decode(&t); err != nil {
        return Ticket{}, err
    }
    return t, nil
}
```

`json.Unmarshal` is a one-shot convenience; `json.Decoder` is the
streaming machine underneath, and it carries knobs the convenience call
hides. `DisallowUnknownFields` arms the one that matters here: every key
must match the struct (by tag or field name), or decoding fails naming the
offender. Same input shape, same target type — only the tolerance changed,
and tolerance was always a setting, not a law.

The test exercises the seam: valid JSON decodes identically under both
policies (strictness costs nothing when input is clean), while
`{"id":1,"title":"t","zzz":1}` — innocent to `Unmarshal` — errors under
the decoder. One flag, and the typo becomes evidence instead of mystery.

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
