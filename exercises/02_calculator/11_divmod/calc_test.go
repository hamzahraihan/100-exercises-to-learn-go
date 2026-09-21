package calc

import "testing"

func TestDivMod(t *testing.T) {
	q, r := DivMod(7, 3)
	if q != 2 || r != 1 {
		t.Fatalf("DivMod(7,3) = (%d,%d), want (2,1)", q, r)
	}
}
