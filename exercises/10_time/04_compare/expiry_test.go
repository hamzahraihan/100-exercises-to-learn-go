package expiry

import (
	"testing"
	"time"
)

func TestIsExpired(t *testing.T) {
	old := Now
	Now = func() time.Time { return time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC) }
	defer func() { Now = old }()
	if !IsExpired(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("past expiry not expired")
	}
	if IsExpired(time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("future expiry reported expired")
	}
}
