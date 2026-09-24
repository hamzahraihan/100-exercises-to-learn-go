# Strings

Strings support `==` and `+`, slice like arrays — and then, just as you
reach for them, refuse to be changed. Try it:

```go
s := "hello"
s[0] = 'H' // compile error: cannot assign to s[0]
```

**Strings are immutable.** There is no edit-in-place, no character
replacement, no append-to-same-string. Every "modification" builds a brand
new string from pieces. Once that clicks, the whole standard library's
shape makes sense.

## Surgery with slices

Reading a piece is free — slicing works exactly as on arrays:

```go
s := "hello"
s[:1]  // "h"   — everything before index 1
s[1:]  // "ello" — everything from index 1 on
```

Writing means concatenating the pieces back together around the change:

```go
func ToUpperFirst(s string) string {
    if s == "" {
        return ""
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
```

Two details carry the lesson. First, the **empty-string guard**: `s[:1]`
on `""` panics with index out of range — the emergency brake from the
panics lesson, triggered by slicing past the end. Any function that slices
must first ask whether there's anything to slice. Second, `strings.ToUpper`
does the casing work; the standard library you met in the packages lesson
keeps paying rent.

One honest footnote: `s[:1]` cuts *bytes*, so this recipe assumes ASCII
input. A multibyte first character (the runes lesson's `é`) would split
mid-character. Handling that correctly means decoding to runes first —
`[]rune(s)` — which the strings section will drill properly. For today,
the test speaks ASCII, and the guard-plus-slice pattern is the point.

## Task

Fix `ToUpperFirst` in `calc.go`. Keep `""` mapping to `""`.

## Check

```bash
go test ./exercises/02_calculator/14_strings/ -v
```
