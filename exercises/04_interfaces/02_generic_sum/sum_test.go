package gsum

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int64{1, 2, 3}); got != 6 {
		t.Fatalf("Sum ints = %d, want 6", got)
	}
	if got := Sum([]float64{1.5, 2.5}); got != 4.0 {
		t.Fatalf("Sum floats = %v, want 4.0", got)
	}
}
