// Package trim teaches trimming.
package trim

import "strings"

// Clean strips surrounding whitespace.
func Clean(s string) string {
	return strings.TrimSpace(s)
}

// TrimExt strips a trailing ".txt".
func TrimExt(name string) string {
	return strings.TrimSuffix(name, ".txt")
}
