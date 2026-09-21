// Package divide teaches asserting errors.
package divide

import "errors"

// ErrZeroDivisor is returned for division by zero.
var ErrZeroDivisor = errors.New("division by zero")

// Divide returns a/b, or ErrZeroDivisor when b is 0.
// TODO: branch on b == 0.
func Divide(a, b int) (int, error) {
	return 0, nil
}
