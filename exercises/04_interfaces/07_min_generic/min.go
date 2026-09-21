// Package min teaches ordered constraints.
package min

import "cmp"

// Min returns the smaller of a and b for any ordered type.
func Min[T cmp.Ordered](a, b T) T {
	if b < a {
		return b
	}
	return a
}
