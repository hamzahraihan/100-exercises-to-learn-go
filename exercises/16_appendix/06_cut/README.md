# Cut

Splitting `"key=value"` on `"="` sounds like a job for `Index` plus
slicing: find the position, cut around it, mind the `-1` when absent.
Three steps, two off-by-one opportunities, one missing-separator branch —
for the most common string operation in config parsing. `strings.Cut`
collapses all of it into a single call that cannot misindex.

## One call, three answers

```go
key, value, ok := strings.Cut("a=1", "=")
// key=="a", value=="1", ok==true

_, _, ok := strings.Cut("novalue", "=")
// ok==false — nothing to split, no slicing attempted

key, value, _ := strings.Cut("a=b=c", "=")
// key=="a", value=="b=c" — splits on the FIRST separator only
```

Before, after, found — the complete verdict in one return. Absence
arrives as `false` rather than `-1`, so there's no index to misread and
no `s[:i]`/`s[i+1:]` arithmetic to fumble. First-separator semantics fall
out naturally: everything past the first `=` stays intact inside `value`,
which is precisely what `key=value=with=equals` requires.

Compare the old dance it replaces:

```go
// The Index pas de deux, retired:
i := strings.Index(s, "=")
if i < 0 { /* absent */ }
key, value := s[:i], s[i+1:] // two slices, one fencepost each
```

Every line there is a chance to err — `< 0` vs `== -1`, `i` vs `i+1`,
empty-key edge cases. `Cut` deletes the dance instead of teaching it.

## The found-flag over ambiguity

```go
func SplitKV(s string) (key, value string, ok bool) {
    if before, after, found := strings.Cut(s, "="); found {
        return before, after, true
    }
    return "", "", false
}
```

`Cut` does the splitting; the wrapper normalizes absence to zero values.
An input with no `=` isn't a key with an empty value — it's *not a pair
at all* — so `("", "", false)` says exactly that. Callers checking `ok`
first never touch the empties; callers ignoring `ok` get nothing
misleading either. One branch, and the absent case can never masquerade
as data.

The signature returns the split *and* its success, because `""` alone is
ambiguous — is `value` empty, or was there no `=` at all? The query lesson
faced the same fog with `Get` returning `""` for missing-or-empty, and
resolved it by policy. `Cut` resolves it structurally: `ok` separates the
cases, no policy needed. When an API can distinguish, it should — boolean
verdicts beat empty-string mind-reading, the paired-returns rhythm in yet
another costume.

Siblings cover the edges: `CutPrefix`/`CutSuffix` trim one known affix
with the same found-flag honesty (replacing `TrimPrefix` guesswork about
whether anything was cut), and `Cut` itself generalizes to any separator.
Reach for the family whenever a split has exactly two interesting sides.

## Task

Fill in `SplitKV` in `kv.go` so `"a=1"` splits and bare words report
absent:

```go
func SplitKV(s string) (key, value string, ok bool) {
    // ...
}
```

Use `strings.Cut` on `"="` — first separator wins, absence reports `false`.

## Check

```bash
go test ./exercises/16_appendix/06_cut/ -v
```
