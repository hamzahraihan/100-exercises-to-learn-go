package calc

import "testing"

func TestHalf(t *testing.T) {
	if got, want := Half(5), 2.5; got != want {
		t.Fatalf("Half(5) = %v, want %v", got, want)
	}
	if got, want := Half(4), 2.0; got != want {
		t.Fatalf("Half(4) = %v, want %v", got, want)
	}
}
