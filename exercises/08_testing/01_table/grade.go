// Package grade teaches table-driven tests.
package grade

// Grade maps a score to A/B/C/F.
func Grade(score int) string {
	if score >= 90 {
		return "A"
	}
	if score >= 80 {
		return "B"
	}
	if score >= 70 {
		return "C"
	}
	return "F"
}
