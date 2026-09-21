// Package redact teaches replacement.
package redact

import "strings"

// Redact replaces every "secret" with "[redacted]".
func Redact(s string) string {
	return strings.ReplaceAll(s, "secret", "[redacted]")
}
