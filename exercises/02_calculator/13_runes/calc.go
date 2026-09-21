package calc

import "unicode/utf8"

// CountRunes returns the number of characters (runes) in s.
// len(s) counts bytes instead: "é" is 2 bytes but 1 rune.
func CountRunes(s string) int {
	return utf8.RuneCountInString(s)
}
