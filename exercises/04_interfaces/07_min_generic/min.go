// Package min teaches ordered constraints.
package min

import "cmp"

// Min returns the smaller of a and b for any ordered type.
// TODO: compare and return the smaller (the constraint allows <).
func Min[T cmp.Ordered](a, b T) T {
	return a
}
