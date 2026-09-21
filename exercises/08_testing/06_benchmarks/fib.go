// Package fib teaches benchmarks.
package fib

// Fib returns the n-th Fibonacci number (Fib(0) == 0, Fib(1) == 1).
func Fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
