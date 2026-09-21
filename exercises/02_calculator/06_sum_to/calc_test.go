package calc

import "testing"

func TestSumTo(t *testing.T) {
	if got, want := SumTo(100), 5050; got != want {
		t.Fatalf("SumTo(100) = %d, want %d", got, want)
	}
	if got := SumTo(0); got != 0 {
		t.Fatalf("SumTo(0) = %d, want 0", got)
	}
}
