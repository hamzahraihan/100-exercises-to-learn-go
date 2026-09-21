package calc

import "testing"

func TestCompute(t *testing.T) {
	if got, want := Compute(1, 2), uint32(9); got != want {
		t.Fatalf("Compute(1, 2) = %d, want %d", got, want)
	}
}
