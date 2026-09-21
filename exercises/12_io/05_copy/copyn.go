// Package copyn teaches io.CopyN.
package copyn

import "io"

// CopyN streams exactly n bytes from src to dst.
// TODO: use io.CopyN (no need to load everything).
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
	return 0, nil
}
