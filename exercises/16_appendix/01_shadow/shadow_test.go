package shadow

import "testing"

func TestApplyBonus(t *testing.T) {
	tests := []struct {
		total, bonus, want int
	}{
		{100, 20, 120},
		{100, 0, 100},
		{0, 5, 5},
	}
	for _, tt := range tests {
		if got := ApplyBonus(tt.total, tt.bonus); got != tt.want {
			t.Errorf("ApplyBonus(%d, %d) = %d, want %d", tt.total, tt.bonus, got, tt.want)
		}
	}
}
