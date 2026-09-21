package dayparse

import (
	"testing"
	"time"
)

func TestParseDay(t *testing.T) {
	want := time.Date(2026, 9, 21, 0, 0, 0, 0, time.UTC)
	got, err := ParseDay("2026-09-21")
	if err != nil || !got.Equal(want) {
		t.Fatalf("ParseDay = (%v, %v), want (%v, nil)", got, err, want)
	}
	if _, err := ParseDay("not-a-date"); err == nil {
		t.Fatal("ParseDay(not-a-date) accepted, want error")
	}
}
