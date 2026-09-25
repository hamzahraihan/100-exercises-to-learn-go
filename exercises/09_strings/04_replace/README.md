# Replacement

Scrubbing a word from text sounds like deletion. It's actually
*substitution*: every occurrence swapped for something safer, the rest of
the string untouched. And the emphasis falls on *every* — replacing only
the first match is the classic almost-fix, passing the eyeball test while
leaking secrets further down the page.

## All of them, not the first

```go
strings.ReplaceAll("my secret here", "secret", "[redacted]")
// "my [redacted] here"
```

`strings.ReplaceAll` rewrites every match. Its sibling `Replace` takes a
count — `Replace(s, old, new, 1)` swaps only the first hit, negative means
all (making it `ReplaceAll`'s verbose twin). When you truly want one
replacement, say `1` explicitly; reaching for `ReplaceAll` by default keeps
the common case — scrub everything — impossible to under-apply.

Remember the immutability contract from the strings-calculator lesson: the
input is never modified. `ReplaceAll` *returns* the rewritten copy; the
original survives for logging, auditing, or second thoughts. Redaction that
mutated in place would destroy the evidence trail — here the language's
constraint and the domain's need agree completely.

## The test that already passes

```go
Redact("nothing to hide") // "nothing to hide" — no match, unchanged
```

The second assertion passes against the untouched stub: no occurrences,
nothing to swap, input returned as-is. Replacement is a no-op without
matches — no error, no flag, just the string back. Suites where one case
passes pre-fix aren't broken; they're showing that the function degrades
gracefully. The suite still fails overall (the first case demands the
swap), which is exactly the red you want: one failing assertion pointing
at one missing behavior.

## Task

Fill in `Redact` in `redact.go`:

```go
func Redact(s string) string {
    // ...
}
```

Use `strings.ReplaceAll` to swap every `"secret"` for `"[redacted]"`.

## Check

```bash
go test ./exercises/09_strings/04_replace/ -v
```
