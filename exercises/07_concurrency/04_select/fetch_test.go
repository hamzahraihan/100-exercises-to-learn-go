package fetch

import (
	"testing"
	"time"
)

func TestFetchFast(t *testing.T) {
	got, err := Fetch(10 * time.Millisecond)
	if err != nil || got != "data" {
		t.Fatalf("Fetch(fast) = (%q, %v), want (data, nil)", got, err)
	}
}

func TestFetchTimeout(t *testing.T) {
	got, err := Fetch(500 * time.Millisecond)
	if err == nil {
		t.Fatalf("Fetch(slow) = (%q, nil), want timeout error", got)
	}
}
