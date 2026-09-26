# Fuzzing

Unit tests check the inputs you imagined. The bugs live in the inputs you
didn't — empty strings, lone surrogates, three-megabyte emoji. **Fuzzing**
automates the unimaginable: the harness mutates seed inputs, runs your
code against thousands of mutants, and files any crash as a reproducible
regression. Go ships it built in; this exercise reverse-engineers a
multibyte bug with it.

## Properties, not examples

```go
func FuzzReverse(f *testing.F) {
    f.Add("abc")    // seeds: where exploration starts
    f.Add("héllo")
    f.Add("")
    f.Fuzz(func(t *testing.T, s string) {
        r := Reverse(s)
        if !utf8.ValidString(r) {
            t.Fatalf("Reverse(%q) = %q, invalid UTF-8", s, r)
        }
        if back := Reverse(r); back != s {
            t.Fatalf("Reverse(Reverse(%q)) = %q, want original", s, back)
        }
    })
}
```

A fuzz test states **properties** — truths for *every* input — instead of
examples. Here: the reversal is always valid UTF-8, and reversing twice
restores the original (an involution, like negation). `f.Add` seeds the
corpus with interesting starts, including the tricky `héllo`; the engine
mutates from there into territory no table would cover.

Plain `go test` runs the seeds as ordinary cases — fast, deterministic,
CI-friendly. The real hunt is opt-in:

```bash
go test -fuzz=FuzzReverse -fuzztime=15s ./exercises/16_appendix/05_fuzz/
```

Fifteen seconds of mutation against your implementation. A failure lands
in `testdata/fuzz/` as a permanent regression seed — the fuzzer writes its
own test case, and every future `go test` replays it. Bugs found once stay
found.

## The bug it would have caught

Reverse bytes naively and `"héllo"` shatters: the two bytes of `é`
(`0xC3 0xA9`) swap into `0xA9 0xC3` — invalid UTF-8, caught by the first
property within seconds of fuzzing. The fix thinks in runes, not bytes:

```go
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

Decode to characters, swap symmetrically, re-encode. The runes lesson's
bytes-vs-characters split, now with a machine hunting violations. Fuzz
targets like this — parsers, codecs, anything transforming untrusted
input — repay the setup a hundredfold; pure arithmetic rarely bothers.

## Scope the domain

Properties hold over a *domain*, and inputs outside it aren't
counterexamples — they're out of scope. A lone `0xA2` byte is invalid
UTF-8; decoding replaces it with `U+FFFD`, re-encoding can't restore the
original, and no implementation could make double-reverse hold there. The
fuzz body says so explicitly:

```go
if !utf8.ValidString(s) {
    t.Skip("out of domain: properties hold over valid UTF-8")
}
```

`t.Skip` (not `t.Fail`, not silent acceptance) marks the boundary: this
input is excluded *by decision*, visibly, with the reason attached.
Scoping is part of specifying — a property without a domain is a wish.
When your own fuzz runs find "failures" in territory you never promised,
reach for `Skip` before reaching for excuses.

## Task

Fill in `Reverse` in `reverse.go` so it flips rune order (multibyte-safe):

```go
func Reverse(s string) string {
    // ...
}
```

Decode to `[]rune`, reverse in place, convert back. The table pins exact
outputs; the fuzzer guards the properties.

## Check

```bash
go test ./exercises/16_appendix/05_fuzz/ -v
go test -fuzz=FuzzReverse -fuzztime=15s ./exercises/16_appendix/05_fuzz/
```
