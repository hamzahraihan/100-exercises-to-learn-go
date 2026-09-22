---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Struct Tags: Hiding and Omitting"
weight: 3
draft: false
---

# Struct Tags: Hiding and Omitting

Two tags do most of the shaping work: `json:"-"` keeps a field out of the
JSON entirely (handy for secrets like `InternalNote` that live only in
memory), and `,omitempty` drops a field when it holds its zero value, so an
empty `Description` disappears instead of serializing as `""`. The stub
returns a hardcoded string that leaks the secret, so the test fails on the
leak check.

## Task

Fill in `MarshalTicket` in `ticket.go`:

```go
func MarshalTicket(t Ticket) (string, error) {
	// ...
}
```

Use `json.Marshal` so the struct tags take effect — the `json:"-"` tag keeps
`InternalNote` out of the output on its own.

## Check

```bash
go test ./exercises/11_json/03_tags/ -v
```

---

*Source: `exercises/11_json/03_tags/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
