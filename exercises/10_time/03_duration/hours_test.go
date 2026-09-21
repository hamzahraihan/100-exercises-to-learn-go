package hours

import (
	"testing"
	"time"
)

func TestHoursBetween(t *testing.T) {
	a := time.Date(2026, 9, 21, 10, 0, 0, 0, time.UTC)
	b := time.Date(2026, 9, 21, 11, 30, 0, 0, time.UTC)
	if got, want := HoursBetween(a, b), 1.5; got != want {
		t.Fatalf("HoursBetween = %v, want %v", got, want)
	}
	if got := HoursBetween(a, a); got != 0 {
		t.Fatalf("HoursBetween(same) = %v, want 0", got)
	}
}
