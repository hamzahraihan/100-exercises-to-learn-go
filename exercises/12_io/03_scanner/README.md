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
