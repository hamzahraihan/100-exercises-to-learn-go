package calc

// SumRange adds all slice elements.
func SumRange(vals []int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}
