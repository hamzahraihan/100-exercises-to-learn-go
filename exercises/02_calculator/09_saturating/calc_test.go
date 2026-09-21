package calc

import "testing"

func TestSaturatingAdd(t *testing.T) {
	if got := SaturatingAdd(10, 20); got != 30 {
		t.Fatalf("SaturatingAdd(10,20) = %d, want 30", got)
	}
	if got := SaturatingAdd(200, 100); got != 255 {
		t.Fatalf("SaturatingAdd(200,100) = %d, want 255 (saturated)", got)
	}
}
