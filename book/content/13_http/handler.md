---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Your First Handler"
weight: 1
draft: false
---

# Your First Handler

In Go, an HTTP handler is any function with the signature
`func(http.ResponseWriter, *http.Request)` — the `http.HandlerFunc` type
adapts such a function so it can serve traffic. You write the response body
through the `ResponseWriter` (it is an `io.Writer`), and you read everything
about the request from `*http.Request`. Tests never bind a socket: they call
your handler directly with `httptest.NewRequest` to build a request and
`httptest.NewRecorder` to capture what you write.

## Task

Fill in `Hello` in `hello.go`:

```go
func Hello(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

Write `"hello"` to `w`. The stub writes nothing, so the test fails on the
body.

## Check

```bash
go test ./exercises/13_http/01_handler/ -v
```

---

*Source: `exercises/13_http/01_handler/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
