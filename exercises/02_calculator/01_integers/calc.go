package calc

// Compute returns a + b*multiplier where multiplier is 4.
// TODO: multiply b by the uint8 multiplier (convert it to uint32 first).
func Compute(a, b uint32) uint32 {
	var multiplier uint8 = 4
	_ = multiplier
	return a + b
}
