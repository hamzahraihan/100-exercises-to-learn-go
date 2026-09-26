package reverse

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	tests := []struct{ in, want string }{
		{"abc", "cba"},
		{"héllo", "olléh"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := Reverse(tt.in); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	f.Add("abc")
	f.Add("héllo")
	f.Add("")
	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip("out of domain: properties hold over valid UTF-8")
		}
		r := Reverse(s)
		if !utf8.ValidString(r) {
			t.Fatalf("Reverse(%q) = %q, invalid UTF-8", s, r)
		}
		if back := Reverse(r); back != s {
			t.Fatalf("Reverse(Reverse(%q)) = %q, want original", s, back)
		}
	})
}
