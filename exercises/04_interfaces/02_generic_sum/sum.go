// Package gsum teaches generics (the Go answer to Rust trait bounds).
package gsum

// Sum adds all values. The constraint permits int64 and float64.
func Sum[T int64 | float64](vals []T) T {
	var total T
	for _, v := range vals {
		total += v
	}
	return total
}
