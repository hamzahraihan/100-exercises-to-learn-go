package syntax

import "testing"

func TestCompute(t *testing.T) {
	if got, want := Compute(2, 3), 5; got != want {
		t.Fatalf("Compute(2, 3) = %d, want %d", got, want)
	}
}
