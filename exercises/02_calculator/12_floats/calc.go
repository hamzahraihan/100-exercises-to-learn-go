package calc

// Half returns n divided by two. Integer division truncates before the
// conversion, so convert first.
// TODO: return float64(n) / 2.
func Half(n int) float64 {
	return float64(n / 2)
}
