# Splitting and Joining

Real input never arrives neatly. Users double-space, paste tabs, leave
trailing newlines — and somewhere inside the mess sit the words you
actually want. This exercise splits mess into words and joins words back
into order, with the standard library doing both directions.

## Two splitters, different promises

```go
strings.Fields("  go  is fun ") // ["go" "is" "fun"]
strings.Split("  go  is fun ", " ") // ["" "" "go" "" "is" "fun" ""]
```

Same input, wildly different output. `strings.Fields` splits on *runs* of
whitespace — spaces, tabs, newlines, any length — and drops the empties.
`strings.Split` cuts on one *exact* separator and keeps every piece,
including the empty strings between adjacent separators.

Which to reach for is a question about your input, not your taste.
Human-typed text with unpredictable spacing wants `Fields`: the test's
`"  go  is fun "` becomes exactly `["go" "is" "fun"]`, edges and doubles
forgiven. Machine-formatted text with a strict delimiter (`"a,b,c"`,
`"key=value"`) wants `Split`, where empties carry meaning — a missing
field between commas is data, not dirt. Choosing wrong here doesn't error;
it silently returns the wrong pieces, which is worse.

## Reassembly

```go
strings.Join([]string{"a", "b"}, " ") // "a b"
```

`strings.Join` is the inverse: one separator placed *between* each
element, never trailing. Together the pair round-trips with a side effect
worth knowing — `Join(Words(s))` normalizes any spacing mess into
single-spaced words. Split dirty, join clean: a legitimate one-line
canonicalizer you'll reuse on user input for the rest of your career.

Both stubs fail at once (`nil` words, `""` join), so the suite starts
fully red. Fix `Words` first — `JoinWords` has nothing sensible to join
until splitting works — and watch the tests flip in dependency order.

## Task

Fill in `Words` and `JoinWords` in `words.go`:

```go
func Words(s string) []string {
    // ...
}
```

```go
func JoinWords(w []string) string {
    // ...
}
```

Use `strings.Fields` to split on whitespace runs, and `strings.Join` with a
single-space separator to put words back together.

## Check

```bash
go test ./exercises/09_strings/01_fields/ -v
```
