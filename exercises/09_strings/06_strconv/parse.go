// Package parse teaches string/number conversion.
package parse

import "strconv"

// SumStrings parses a and b and returns their sum.
func SumStrings(a, b string) (int, error) {
	ai, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	bi, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return ai + bi, nil
}
