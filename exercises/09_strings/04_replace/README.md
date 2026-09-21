# Replacement

Scrubbing a word out of text means swapping every occurrence, not just the
first. `strings.ReplaceAll` rewrites all matches; its sibling `Replace` takes
a count argument so you can limit the swap to the first `n` hits (negative
means all). The stub returns its input untouched, so the first test case fails
while the no-match case already passes — the suite as a whole still fails.

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
