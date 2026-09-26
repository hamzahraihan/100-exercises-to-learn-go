// Package reverse teaches property testing with native fuzzing.
package reverse

// Reverse returns s with rune order flipped (multibyte-safe).
// TODO: decode to []rune, reverse in place, convert back to string.
func Reverse(s string) string {
	return s
}
