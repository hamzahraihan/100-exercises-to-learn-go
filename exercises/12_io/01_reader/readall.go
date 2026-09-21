// Package readall teaches the io.Reader contract.
package readall

import "io"

// ReadAll drains r to a string.
func ReadAll(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
