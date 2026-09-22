---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Query Parameters"
weight: 3
draft: false
---

# Query Parameters

The `?status=open` tail of a URL carries query parameters. Parse them with
`r.URL.Query()`, which returns a `url.Values` map, and read one key with its
`.Get("status")` method. `Get` returns `""` both when the key is absent and
when it is present-but-empty, so handlers usually treat `""` as "not given"
and substitute a default.

## Task

Fill in `StatusParam` in `filter.go`:

```go
func StatusParam(r *http.Request) string {
	// ...
}
```

Return the `"status"` query value, or `"all"` when it is missing. The stub
always returns `""`, so the test fails on the returned value.

## Check

```bash
go test ./exercises/13_http/03_query/ -v
```

---

*Source: `exercises/13_http/03_query/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
