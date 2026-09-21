// Package mail teaches substring checks.
package mail

import "strings"

// IsEmail is a deliberately naive check: an address contains "@".
func IsEmail(s string) bool {
	return strings.Contains(s, "@")
}
