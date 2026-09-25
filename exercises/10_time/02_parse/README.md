# Parsing Time

Formatting turns a moment into text. Parsing turns text back into a
moment — same layout language, opposite direction, and a strictness that
surprises everyone exactly once.

## The layout must match, exactly

```go
got, err := time.Parse("2006-01-02", "2026-09-21")
// got == 2026-09-21 00:00:00 UTC, err == nil
```

`time.Parse` reads the input *through* the layout: where the layout shows
`2006` it expects a year, where `01` a month, where `02` a day. Feed it
`"21/09/2026"` against `"2006-01-02"` and it returns an error instead of a
time — no guessing, no partial credit, no "close enough." That strictness
is the feature: a mismatched layout fails loudly at the border rather than
smuggling a misread date into the database.

So the implementation passes errors straight through:

```go
func ParseDay(s string) (time.Time, error) {
    return time.Parse("2006-01-02", s)
}
```

No wrapping, no translation — the `Parse` error already names the problem
(`cannot parse "not-a-date" as "2006"`), and this layer adds no context
worth attaching. The strconv lesson's restraint, applied to calendars: pass
through where there's nothing to add.

## Compare instants with Equal

The test checks the result with a method, not an operator:

```go
if err != nil || !got.Equal(want) {
```

`time.Time` values carry more than wall-clock readings — a monotonic
component for measuring elapsed time, plus location metadata. `==`
compares the whole struct, so two readings of the *same instant* from
different sources can differ. `Equal` compares the instants themselves,
which is what "same time" means. Rule of thumb: `Before`, `After`, and
`Equal` for moments; `==` never. The comparison exercise ahead leans on
this hard — consider this paragraph its early warning.

## Both halves get tested

```go
got, err := ParseDay("2026-09-21") // success: instant AND nil error
_, err = ParseDay("not-a-date")    // failure: error present, value ignored
```

The two-branch assertion habit from the error-assertion lesson, intact:
happy path checks value *and* nil error, failure path checks error
presence with `_` discarding the meaningless time. Parse functions always
have both branches worth pinning — valid input converts, junk complains —
and the table here covers each exactly once.

## Task

Fill in `ParseDay` in `day.go`:

```go
func ParseDay(s string) (time.Time, error) {
    // ...
}
```

Use `time.Parse` with the `2006-01-02` layout; return the parse error
unchanged on bad input.

## Check

```bash
go test ./exercises/10_time/02_parse/ -v
```
