package calc

import "testing"

func TestCountRunes(t *testing.T) {
	if got, want := CountRunes("go"), 2; got != want {
		t.Fatalf("CountRunes(%q) = %d, want %d", "go", got, want)
	}
	if got, want := CountRunes("héllo"), 5; got != want {
		t.Fatalf("CountRunes(%q) = %d, want %d", "héllo", got, want)
	}
}
