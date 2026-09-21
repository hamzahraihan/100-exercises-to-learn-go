// Package dayfmt teaches time formatting.
package dayfmt

import "time"

// FormatDay renders t as YYYY-MM-DD.
func FormatDay(t time.Time) string {
	return t.Format("2006-01-02")
}
