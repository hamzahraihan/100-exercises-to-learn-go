# Join

Every validator so far stopped at the first complaint. One bad title *or*
one short description — never both, even when both are wrong. That's
fail-fast, and it's rude: the user fixes the title, resubmits, and only
then learns about the description. Forms, configs, and APIs all deserve
better: the full list, in one round trip.

## Collect, then combine

```go
var (
    ErrBadTitle = errors.New("bad title")
    ErrBadDesc  = errors.New("bad description")
)

func ValidateAll(title, desc string) error {
    var errs []error
    if title == "" {
        errs = append(errs, ErrBadTitle)
    }
    if len(desc) < 10 {
        errs = append(errs, ErrBadDesc)
    }
    return errors.Join(errs...)
}
```

Each rule appends its sentinel independently — no `else`, no early return,
no rule aware of the others. Then `errors.Join` fuses the collection into
one error. The fusion preserves every identity: `errors.Is(joined,
ErrBadTitle)` and `errors.Is(joined, ErrBadDesc)` both hold, so callers
test parts exactly as if each had arrived alone.

## The empty join is nil

The elegant half of the design: `errors.Join()` with no arguments returns
`nil`. Valid input appends nothing, joins nothing, and the function returns
`nil` — success, with no special case written:

```go
if err := ValidateAll("Ticket title", "long enough description"); err != nil {
```

Zero code paths for the happy path. The test's first assertion pins this:
valid input must yield a `nil` error, not an empty-but-non-nil wrapper.
Sentinel-per-field plus join-on-exit makes "report everything" and "report
nothing" the same code.

## Fail-fast versus report-all

Two philosophies, different homes. The constructor from two exercises ago
returns the *first* failure — right for creation, where one verdict gates
everything downstream. `ValidateAll` reports *every* failure — right for
user-facing validation, where each round trip costs patience. Neither is
universally better; the question is always who pays for another round.

Note the test's granularity: both-bad asserts both identities present,
title-only-bad asserts `ErrBadTitle` present *and `ErrBadDesc` absent*.
Partial failure must not hallucinate the other half. Independent appends
give this for free — each sentinel appears if and only if its own rule
fired.

## Task

Complete `ValidateAll` in `ticket.go` so empty titles report `ErrBadTitle`,
short descriptions report `ErrBadDesc`, both problems are reported together,
and valid input returns `nil`.

## Check

```bash
go test ./exercises/05_ticket_v2/04_join/ -v
```
