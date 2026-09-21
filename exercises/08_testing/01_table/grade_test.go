package grade

import "testing"

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{59, "F"},
	}
	for _, tt := range tests {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}
