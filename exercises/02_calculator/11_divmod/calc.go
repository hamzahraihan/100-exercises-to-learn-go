package calc

// DivMod returns quotient and remainder. Integer division
// truncates toward zero: -5/2 == -2.
func DivMod(a, b int) (quotient, remainder int) {
	return a / b, a % b
}
