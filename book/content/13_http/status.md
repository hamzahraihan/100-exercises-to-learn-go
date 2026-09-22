---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Status Codes"
weight: 2
draft: false
---

# Status Codes

Every response carries a numeric status code; the default is `200 OK`.
When your handler creates a resource, answer `201 Created` with
`w.WriteHeader(http.StatusCreated)`. The one rule to remember:
`WriteHeader` must come before the first body write — once bytes hit the
wire, the status is locked in and later calls are ignored.

## Task

Fill in `Create` in `create.go`:

```go
func Create(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

Answer `http.StatusCreated`. The stub writes nothing, so the recorder keeps
the default `200` and the test fails on the status.

## Check

```bash
go test ./exercises/13_http/02_status/ -v
```

---

*Source: `exercises/13_http/02_status/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
