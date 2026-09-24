// Package gsum teaches generics with type constraints.
package gsum

// Sum adds all values. The constraint permits int64 and float64.
// TODO: range over vals and accumulate.
func Sum[T int64 | float64](vals []T) T {
	var total T
	return total
}
