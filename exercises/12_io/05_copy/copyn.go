// Package copyn teaches io.CopyN.
package copyn

import "io"

// CopyN streams exactly n bytes from src to dst.
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
	return io.CopyN(dst, src, n)
}
