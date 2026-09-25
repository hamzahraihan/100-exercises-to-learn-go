# Formatting Time

Every language formats dates with codes: `%Y` for year, `%m` for month, a
table memorized once and forgotten monthly. Go looked at that table and
replaced it with a single moment in history. No codes. Just an example.

## The reference moment

```go
t.Format("2006-01-02") // e.g. "2026-09-21"
```

The layout `"2006-01-02"` isn't a pattern of symbols — it's a *date*: what
you write is how January 2nd, 2006 renders in the shape you want, and Go
maps each component onto the real value. Year where `2006` sits, month
where `01` sits, day where `02` sits. Want slashes? Write `"2006/01/02"`.
Want the month's name? The reference month is January, so `"Jan"` gives
`"Sep"` and `"January"` gives `"September"`.

The full reference moment is `Mon Jan 2 15:04:05 MST 2006` — month, day,
hour, minute, second, year, timezone, each piece usable in a layout:

```go
t.Format("15:04")          // "15:04" — hour:minute, 24-hour clock
t.Format("Jan 2, 2006")    // "Sep 21, 2026"
t.Format(time.RFC3339)     // "2026-09-21T15:04:00Z" — the machine constant
```

The digits run 1-2-3-4-5-6 (month 1, day 2, hour 3... er, 15, minute 4,
second 5, year 6) — a mnemonic hiding in plain sight once someone points
it out. And for machine timestamps, skip bespoke layouts: the `time`
package ships `RFC3339`, `Kitchen`, and friends as constants, so every Go
program on earth writes interoperable timestamps the same way.

## Format is a method, not a function

```go
// Syntax: <time>.Format(<layout>)
func FormatDay(t time.Time) string {
    return t.Format("2006-01-02")
}
```

Formatting hangs off the value — `t.Format`, never `Format(t, ...)`.
Times carry their location with them (the test builds its input with
`time.UTC` explicitly), so the same instant renders differently per zone
without extra arguments. Data and presentation travel together; the layout
only chooses the outfit.

## Task

Fill in `FormatDay` in `day.go`:

```go
func FormatDay(t time.Time) string {
    // ...
}
```

Use `t.Format` with the `2006-01-02` layout to render just the calendar date.

## Check

```bash
go test ./exercises/10_time/01_format/ -v
```
