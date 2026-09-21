package calc

import "testing"

func TestToInt64(t *testing.T) {
	if got := ToInt64(-42); got != -42 {
		t.Fatalf("ToInt64(-42) = %d, want -42", got)
	}
}
