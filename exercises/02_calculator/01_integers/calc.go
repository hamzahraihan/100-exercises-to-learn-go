package calc

// Compute returns a + b*multiplier where multiplier is 4.
func Compute(a, b uint32) uint32 {
	var multiplier uint8 = 4
	return a + b*uint32(multiplier)
}
