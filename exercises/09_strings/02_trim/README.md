# Trimming

Dirt collects at the edges. Pasted input trails newlines, filenames carry
extensions, config values arrive padded — and the interesting part sits
untouched in the middle. Trimming is the art of naming exactly which edges
to cut, and reaching past the right tool ruins the middle along with them.

## Three cutters, three promises

```go
strings.TrimSpace("  hi\n")     // "hi" — all surrounding whitespace, gone
strings.TrimSuffix("notes.txt", ".txt") // "notes" — one exact piece, if present
strings.Trim("...hi...", ".")   // "hi" — any of the cutset chars, repeatedly
```

`TrimSpace` is the blunt instrument you'll use most: every Unicode
whitespace character, both ends, no questions. `TrimSuffix` is the scalpel:
it removes one *exact* trailing literal, and only if it's actually there —
`"notes.md"` passes through untouched, no error, no partial cut.

Then the cutset family — `Trim`, `TrimLeft`, `TrimRight` — which trims any
*of a set* of characters, repeatedly, from the ends. Powerful, and the
source of a classic over-trim:

```go
strings.Trim("hello", "ho") // "ell" — not "hell"!
```

The cutset `"ho"` means *h-or-o*, so both ends get eaten past what anyone
intended. Whenever you know the literal (`".txt"`), prefer the exact
`TrimSuffix`/`TrimPrefix` form; save cutsets for character *classes*
(whitespace runs, quote characters) where "any of these" is truly the rule.

## Echoes that fail

Both stubs return their input unchanged — and both tests fail, each on its
own assertion. `Clean("  hi\n")` must lose the padding; `TrimExt` must lose
the extension. Two functions, two tools, one file: the pairing is
deliberate, showing the blunt and precise cutters side by side so the
*choice* between them becomes the lesson. Most string-cleaning bugs aren't
wrong tools but imprecise ones — a cutset where a suffix belonged.

## Task

Fill in `Clean` and `TrimExt` in `trim.go`:

```go
func Clean(s string) string {
    // ...
}
```

```go
func TrimExt(name string) string {
    // ...
}
```

Use `strings.TrimSpace` to strip surrounding whitespace, and
`strings.TrimSuffix` to drop a trailing `".txt"`.

## Check

```bash
go test ./exercises/09_strings/02_trim/ -v
```
