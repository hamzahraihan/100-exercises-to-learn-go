# Substring Checks

Most string questions are yes-or-no: does it contain this, start with
that, end with the other? The `strings` package answers each with a
function named like the question itself — and this exercise asks the
simplest one through a deliberately naive email check.

## Questions, not answers

```go
strings.Contains("a@b.co", "@")  // true — piece anywhere inside
strings.HasPrefix("go.mod", "go") // true — piece at the start
strings.HasSuffix("notes.txt", ".txt") // true — piece at the end
```

`Contains`, `HasPrefix`, `HasSuffix`: the call reads as the question,
which means call sites rarely need comments. When the check grows teeth —
case-insensitivity (`EqualFold`), counting occurrences (`Count`), locating
positions (`Index`) — the same package holds them, same naming rhythm.

## Honest naivety

```go
// Deliberately naive: an address contains "@". Nothing more.
func IsEmail(s string) bool {
    return strings.Contains(s, "@")
}
```

A real validator checks structure on both sides, rejects spaces, handles
quoted oddities — a small parser wearing a regex costume. This exercise
claims none of that. The comment says *naive* up front, and the table
proves the limits are known: `"a@b.co"` passes, `"nope"` fails, and `"@"`
alone passes too — a single at-sign is technically "containing @", and the
table enshrines that absurdity rather than hiding it.

That honesty is the actual lesson. Every codebase holds checks like this:
good enough for the current caller, wrong for the general case. Marking
the boundary in a comment (`deliberately naive`) plus a table row showing
the edge (`"@" → true`) turns a future bug report into a planned upgrade.
Unmarked shortcuts rot silently; marked ones wait their turn visibly.

The stub returns `false` for everything, so the table fails on its
positive rows while `"nope"` passes vacuously — partial red, the signature
of a boolean stub. One `Contains` call flips the whole table green.

## Task

Fill in `IsEmail` in `mail.go`:

```go
func IsEmail(s string) bool {
    // ...
}
```

Use `strings.Contains` to report whether the address holds an `"@"`.

## Check

```bash
go test ./exercises/09_strings/03_contains/ -v
```
