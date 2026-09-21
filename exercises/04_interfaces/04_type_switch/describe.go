// Package kind teaches type switches.
package kind

import "fmt"

// Describe names supported dynamic types.
func Describe(v any) string {
	switch v := v.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return fmt.Sprintf("string: %s", v)
	default:
		return "unknown"
	}
}
