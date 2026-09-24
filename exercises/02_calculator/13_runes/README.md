# Runes

Count the characters in `"héllo"`. Five — obviously. Now ask Go:

```go
len("héllo") // 6?!
```

Both answers are right; they're answering different questions. `len`
counts **bytes**, and `é` needs two bytes in UTF-8. The string holds six
bytes that render as five characters. Everything in this exercise follows
from that split.

## Bytes are storage, runes are characters

A Go `string` is a read-only sequence of bytes. A **`rune`** is a Unicode
code point — what you mean by "character" — and it's an alias for `int32`:

```go
'é'  // a rune literal (single quotes!): the number 233 with manners
"é"  // a string of length 2: the bytes 0xC3 0xA9
```

Single quotes make runes, double quotes make strings. Indexing a string
yields a byte (`s[1]` is a `uint8`), which is why byte-indexing into
non-English text returns beautiful nonsense. The string doesn't know where
its characters are — UTF-8 is a variable-width encoding, and only decoding
reveals the boundaries.

## Counting what humans count

The fix reaches for the standard library's decoder:

```go
import "unicode/utf8"

func CountRunes(s string) int {
    return utf8.RuneCountInString(s) // "héllo" -> 5
}
```

`utf8.RuneCountInString` walks the bytes and counts decoded characters.
Its sibling `utf8.DecodeRuneInString` decodes one step at a time for manual
iteration — but manual iteration has a better vehicle. Ranging over a string
decodes as it goes:

```go
n := 0
for range s { // index and rune available; both skippable
    n++
}
```

That's the same `range` from the slice lesson, now doing double duty: over
slices it yields indexes and elements, over strings it yields byte offsets
and runes. The keyword didn't change; the collection taught it a new trick.

So when is `len(s)` correct? When you want bytes — buffer sizes, wire
lengths, capacity checks. When you want characters — display widths, input
limits, "hello" counts — decode. Mismatching the two is an evergreen source
of sliced-in-half accented characters; the strings section ahead will show
the safe slicing patterns.

## Task

Fix `CountRunes` in `calc.go` so multibyte characters count once.

## Check

```bash
go test ./exercises/02_calculator/13_runes/ -v
```
