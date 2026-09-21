package calc

// SumTo returns 1+2+...+n (0 for n <= 0). Go has only `for`:
// it covers while-style loops too.
func SumTo(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}
