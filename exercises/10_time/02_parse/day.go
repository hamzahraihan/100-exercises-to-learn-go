// Package dayparse teaches time parsing.
package dayparse

import "time"

// ParseDay parses YYYY-MM-DD.
func ParseDay(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}
