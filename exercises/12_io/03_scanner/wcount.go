// Package wcount teaches bufio.Scanner.
package wcount

import (
	"bufio"
	"io"
)

// WordCount counts whitespace-separated words in r.
func WordCount(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := 0
	for sc.Scan() {
		n++
	}
	return n, sc.Err()
}
