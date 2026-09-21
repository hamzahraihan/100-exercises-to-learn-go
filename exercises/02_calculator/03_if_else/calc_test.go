package calc

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{3, 7, 7},
		{7, 3, 7},
		{4, 4, 4},
		{-1, -5, -1},
	}
	for _, tt := range tests {
		if got := Max(tt.a, tt.b); got != tt.want {
			t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
