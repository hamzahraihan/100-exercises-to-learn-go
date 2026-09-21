// Package greet teaches imports and exported names.
package greet

import "strings"

// Shout returns s in upper case.
// TODO: use strings.ToUpper instead of strings.ToLower.
func Shout(s string) string {
	return strings.ToLower(s)
}
