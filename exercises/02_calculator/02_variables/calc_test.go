package calc

import "testing"

func TestDouble(t *testing.T) {
	if got, want := Double(21), 42; got != want {
		t.Fatalf("Double(21) = %d, want %d", got, want)
	}
}
