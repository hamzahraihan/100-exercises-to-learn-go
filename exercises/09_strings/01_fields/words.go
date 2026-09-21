// Package words teaches splitting and joining.
package words

import "strings"

// Words splits s on whitespace runs.
func Words(s string) []string {
	return strings.Fields(s)
}

// JoinWords joins with single spaces.
func JoinWords(w []string) string {
	return strings.Join(w, " ")
}
