// Package hours teaches durations.
package hours

import "time"

// HoursBetween returns (b-a) in hours, possibly fractional.
func HoursBetween(a, b time.Time) float64 {
	return b.Sub(a).Hours()
}
