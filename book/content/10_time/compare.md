---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Testable Time"
weight: 4
draft: false
---

# Testable Time

Code that calls `time.Now` directly cannot be tested deterministically: the
answer changes every run. The fix is a fake clock — a package-level `Now`
variable holding a function. Production leaves it as `time.Now`; tests swap
in a function returning a fixed `time.Date` fixture and restore the original
with `defer`. The stub returns `false` for everything, so the past-expiry
case fails on the assertion.

## Task

Fill in `IsExpired` in `expiry.go` (keep the `Now` variable as the clock):

```go
func IsExpired(expiry time.Time) bool {
	// ...
}
```

Compare the fake clock's reading against `expiry` — expiry counts as expired
once the current time is past it.

## Check

```bash
go test ./exercises/10_time/04_compare/ -v
```

---

*Source: `exercises/10_time/04_compare/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
