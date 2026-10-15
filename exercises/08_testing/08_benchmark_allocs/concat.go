// Package benchallocs teaches b.ReportAllocs and allocation-aware benchmarking.
package benchallocs

import "strings"

// Ensure "strings" is imported; remove this line when ConcatFast uses strings.Builder.
var _ strings.Builder

// ConcatSlow repeats s n times using += (O(n) allocations).
func ConcatSlow(s string, n int) string {
	out := ""
	for range n {
		out += s
	}
	return out
}

// ConcatFast repeats s n times using strings.Builder (O(1) allocations).
// TODO: declare a strings.Builder variable (e.g. var b strings.Builder).
// TODO: call b.Grow(len(s) * n) to pre-allocate the exact buffer size.
// TODO: loop n times calling b.WriteString(s).
// TODO: return b.String().
func ConcatFast(s string, n int) string {
	return ConcatSlow(s, n) // TODO: replace with strings.Builder implementation
}
