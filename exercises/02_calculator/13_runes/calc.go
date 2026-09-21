package calc

// CountRunes returns the number of characters (runes) in s.
// len(s) counts bytes instead: "é" is 2 bytes but 1 rune.
// TODO: count runes (hint: the unicode/utf8 package helps).
func CountRunes(s string) int {
	return len(s)
}
