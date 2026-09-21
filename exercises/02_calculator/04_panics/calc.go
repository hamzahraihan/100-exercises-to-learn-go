package calc

// Divide returns a/b, panicking on division by zero.
func Divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
