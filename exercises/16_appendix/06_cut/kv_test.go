package kv

import "testing"

func TestSplitKV(t *testing.T) {
	tests := []struct {
		in         string
		key, val  string
		ok        bool
	}{
		{"a=1", "a", "1", true},
		{"novalue", "", "", false},
		{"a=b=c", "a", "b=c", true},
	}
	for _, tt := range tests {
		key, val, ok := SplitKV(tt.in)
		if key != tt.key || val != tt.val || ok != tt.ok {
			t.Errorf("SplitKV(%q) = (%q, %q, %v), want (%q, %q, %v)",
				tt.in, key, val, ok, tt.key, tt.val, tt.ok)
		}
	}
}
