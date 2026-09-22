---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Scanning Words"
weight: 3
draft: false
---

# Scanning Words

Reading line-by-line (or word-by-word) by hand means managing buffers and
partial tokens yourself. `bufio.Scanner` hides that bookkeeping: you pick a
split function such as `ScanWords`, call `Scan` until it returns false, and
read each token out. When the loop ends, `scanner.Err()` tells you whether
it stopped at a clean `EOF` or on a real failure — always check it. The
stub counts nothing, so the test fails on the assertion.

## Task

Fill in `WordCount` in `wcount.go`:

```go
func WordCount(r io.Reader) (int, error) {
	// ...
}
```

Scan `r` with the words split function, count the tokens, and return the
count along with any scan error.

## Check

```bash
go test ./exercises/12_io/03_scanner/ -v
```

---

*Source: `exercises/12_io/03_scanner/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
