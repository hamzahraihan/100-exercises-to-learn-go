---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Structured Errors"
weight: 10
draft: false
---

# Structured Errors

Clients parse errors best when every failure shares one envelope, e.g.
`{"error":"nope"}`. The recipe mirrors the JSON exercise: set
`Content-Type: application/json`, call `WriteHeader(code)` with the caller's
status, then encode the envelope — status before body, always, because the
first body write freezes the code at `200` if you forget.

## Task

Fill in `WriteError` in `errw.go`:

```go
func WriteError(w http.ResponseWriter, code int, msg string) {
	// ...
}
```

Answer `code` with a JSON body carrying `msg` under the `"error"` key. The
stub writes nothing, so the test fails on status, content type, and body.

## Check

```bash
go test ./exercises/13_http/10_error_response/ -v
```

---

*Source: `exercises/13_http/10_error_response/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
