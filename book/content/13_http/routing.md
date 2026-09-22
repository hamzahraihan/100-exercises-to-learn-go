---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Routing with ServeMux"
weight: 4
draft: false
---

# Routing with ServeMux

`http.NewServeMux` routes requests to handlers by pattern. Since Go 1.22,
patterns can include the method (`"GET /tickets"`) and named path segments
(`"GET /tickets/{id}"`); inside the handler, `r.PathValue("id")` returns the
segment the client actually sent. A request that matches no pattern gets a
`404` automatically.

## Task

Fill in `NewMux` in `router.go`:

```go
func NewMux() *http.ServeMux {
	// ...
}
```

Register `"GET /tickets"` to write `"list"` and `"GET /tickets/{id}"` to
write `"one:"` plus the path value. The stub registers nothing, so the test
fails on the bodies.

## Check

```bash
go test ./exercises/13_http/04_routing/ -v
```

---

*Source: `exercises/13_http/04_routing/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
