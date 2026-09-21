package kind

import "testing"

func TestDescribe(t *testing.T) {
	tests := []struct {
		in   any
		want string
	}{
		{42, "int: 42"},
		{"hi", "string: hi"},
		{1.5, "unknown"},
		{nil, "unknown"},
	}
	for _, tt := range tests {
		if got := Describe(tt.in); got != tt.want {
			t.Errorf("Describe(%v) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
