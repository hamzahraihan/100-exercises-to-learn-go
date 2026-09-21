package min

import "testing"

func TestMin(t *testing.T) {
	if got, want := Min(3, 7), 3; got != want {
		t.Fatalf("Min(3, 7) = %v, want %v", got, want)
	}
	if got, want := Min(7, 3), 3; got != want {
		t.Fatalf("Min(7, 3) = %v, want %v", got, want)
	}
	if got, want := Min("b", "a"), "a"; got != want {
		t.Fatalf("Min(b, a) = %q, want %q", got, want)
	}
}
