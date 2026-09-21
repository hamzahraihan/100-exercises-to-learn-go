package mean

import "testing"

func mustMean(t *testing.T, xs []float64, want float64) {
	t.Helper()
	if got := Mean(xs); got != want {
		t.Fatalf("Mean(%v) = %v, want %v", xs, got, want)
	}
}

func TestMean(t *testing.T) {
	mustMean(t, []float64{1, 2, 3}, 2)
	mustMean(t, nil, 0)
}
