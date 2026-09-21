package mail

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"a@b.co", true},
		{"nope", false},
		{"@", true},
	}
	for _, tt := range tests {
		if got := IsEmail(tt.in); got != tt.want {
			t.Errorf("IsEmail(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
