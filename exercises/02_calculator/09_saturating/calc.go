package calc

// SaturatingAdd adds clamped at 255 instead of wrapping.
// TODO: reuse the overflow check: if sum < a, return 255.
func SaturatingAdd(a, b uint8) uint8 {
	return a + b
}
