// Package kv teaches strings.Cut for first-separator splits.
package kv

import "strings"

// SplitKV splits s on the first "=" into key and value, reporting
// ("", "", false) when no separator is present.
func SplitKV(s string) (key, value string, ok bool) {
	if before, after, found := strings.Cut(s, "="); found {
		return before, after, true
	}
	return "", "", false
}
