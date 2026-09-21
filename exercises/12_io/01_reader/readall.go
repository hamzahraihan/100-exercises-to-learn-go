// Package readall teaches the io.Reader contract.
package readall

import "io"

// ReadAll drains r to a string.
// TODO: use io.ReadAll (it reads until EOF).
func ReadAll(r io.Reader) (string, error) {
	return "", nil
}
