# Ticket Errors

Validation errors so far were one-off strings, created where needed and
compared nowhere. That stops working the moment a caller must *react*
differently to different failures — retry on one, abort on another. This
exercise gives failures identities callers can test for, without losing the
human-readable story.

## Name the failure

```go
var ErrUnknownStatus = errors.New("unknown status")
```

A **sentinel error** is a package-level variable naming one specific thing
that can go wrong. Capital E, `Err` prefix, created once with
`errors.New` — the convention is rigid because its power is social: every
caller recognizes the name. Comparing with `==` works when the exact value
travels untouched, but messages get wrapped along the way (next section),
so the robust test unwraps:

```go
if !errors.Is(err, ErrUnknownStatus) {
```

`errors.Is` reports whether the sentinel appears *anywhere* in the chain —
directly or wrapped. Compare identities with `Is`, never with `==` or
string matching. Strings get reworded; sentinels persist.

## Wrap without losing identity

The constructor validates three things now — title, description (the old
rules from the first ticket section), and status. Each failure deserves its
own message *and* its testable identity:

```go
if !status.Valid() {
    return Ticket{}, fmt.Errorf("status %d: %w", status, ErrUnknownStatus)
}
```

The `%w` verb embeds the sentinel inside a richer message: `"status 99:
unknown status"`. Human readers get context (which status?); programmatic
readers get `errors.Is(err, ErrUnknownStatus) == true`. One value serves
both audiences — that's the entire point of wrapping.

Its evil twin is `%v`, which formats the sentinel into plain text and
*severs* the chain: the message looks identical, but `errors.Is` returns
false. Same output, dead identity. When an error crosses a layer boundary,
reach for `%w` by reflex and treat `%v` on errors as a bug until proven
otherwise.

## The constructor assembles

`NewTicketWithStatus` layers all three validations in one function: title
and description checked like before, status checked via `Valid()`, each
rejection wrapped with its context. The tests read both halves — a valid
ticket comes home with its status set, an unknown one arrives as an
`errors.Is`-detectable failure. Flat branches, named failures, no half-built
tickets escaping: the constructor discipline from the first ticket section,
now speaking in identities instead of strings.

## Task

Complete `NewTicketWithStatus` in `ticket.go` so unknown statuses wrap
`ErrUnknownStatus`, valid tickets are returned with their status set, and
the tests pass.

## Check

```bash
go test ./exercises/05_ticket_v2/02_errors/ -v
```
