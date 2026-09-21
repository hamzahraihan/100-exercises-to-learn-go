package calc

import "strings"

// ToUpperFirst capitalizes the first character of s.
// Strings are immutable: build a new one from pieces.
func ToUpperFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
