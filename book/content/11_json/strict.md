---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Strict Decoding"
weight: 4
draft: false
---

# Strict Decoding

By default `json.Unmarshal` ignores unknown fields, which is friendly to
forward compatibility — old code keeps working when a server adds a field —
but it also hides typos like `"titel"`. A `json.Decoder` with
`DisallowUnknownFields` flips the tradeoff: every field must match the
struct, so misspellings and version skew surface as errors. The stub uses
plain `Unmarshal`, so the unknown-field case passes silently and the test
fails.

## Task

Fill in `UnmarshalStrict` in `ticket.go`:

```go
func UnmarshalStrict(data string) (Ticket, error) {
	// ...
}
```

Decode with a `json.Decoder` over `data` with `DisallowUnknownFields`
enabled instead of `json.Unmarshal`.

## Check

```bash
go test ./exercises/11_json/04_strict/ -v
```

---

*Source: `exercises/11_json/04_strict/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
