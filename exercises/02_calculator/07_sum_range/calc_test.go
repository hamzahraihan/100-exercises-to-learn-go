package calc

import "testing"

func TestSumRange(t *testing.T) {
	if got, want := SumRange([]int{1, 2, 3, 4}), 10; got != want {
		t.Fatalf("SumRange = %d, want %d", got, want)
	}
	if got := SumRange(nil); got != 0 {
		t.Fatalf("SumRange(nil) = %d, want 0", got)
	}
}
