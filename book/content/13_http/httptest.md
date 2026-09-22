---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Testing Handlers with httptest"
weight: 6
draft: false
---

# Testing Handlers with httptest

`net/http/httptest` lets you exercise handlers without a network.
`httptest.NewRequest(method, target, body)` builds a `*http.Request`, and
`httptest.NewRecorder()` returns a `ResponseRecorder` that stands in for the
`ResponseWriter`. After calling your handler, inspect the recording:

| Inspector            | What it shows                  |
|----------------------|--------------------------------|
| `rec.Code`           | captured status code           |
| `rec.Header().Get`   | a response header value        |
| `rec.Body.String()`  | captured body as a string      |

This test file doubles as the pattern reference — every handler test in this
section follows the same recorder/request shape.

## Task

Fill in `Ping` in `ping.go`:

```go
func Ping(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

Write `"pong"` to `w`. The stub writes nothing, so the test fails on the
body.

## Check

```bash
go test ./exercises/13_http/06_httptest/ -v
```

---

*Source: `exercises/13_http/06_httptest/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
