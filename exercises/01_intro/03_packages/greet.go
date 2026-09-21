// Package greet teaches imports and exported names.
package greet

import "strings"

// Shout returns s in upper case.
func Shout(s string) string {
	return strings.ToUpper(s)
}
