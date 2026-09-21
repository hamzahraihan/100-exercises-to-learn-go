package calc

// Factorial returns n! with Factorial(0) == 1.
func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
