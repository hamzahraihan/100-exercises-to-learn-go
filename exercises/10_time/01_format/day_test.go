package dayfmt

import (
	"testing"
	"time"
)

func TestFormatDay(t *testing.T) {
	in := time.Date(2026, 9, 21, 15, 4, 0, 0, time.UTC)
	if got, want := FormatDay(in), "2026-09-21"; got != want {
		t.Fatalf("FormatDay = %q, want %q", got, want)
	}
}
