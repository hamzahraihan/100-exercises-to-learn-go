package concat

import "testing"

func TestConcatN(t *testing.T) {
	if got, want := ConcatN("ab", 3), "ababab"; got != want {
		t.Fatalf("ConcatN = %q, want %q", got, want)
	}
	if got := ConcatN("x", 0); got != "" {
		t.Fatalf("ConcatN(x, 0) = %q, want empty", got)
	}
}
