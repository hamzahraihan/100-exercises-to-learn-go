package calc

// SaturatingAdd adds clamped at 255 instead of wrapping.
func SaturatingAdd(a, b uint8) uint8 {
	sum := a + b
	if sum < a {
		return 255
	}
	return sum
}
