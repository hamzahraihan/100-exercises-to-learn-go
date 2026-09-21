package calc

import "testing"

func TestDivide(t *testing.T) {
	if got := Divide(6, 3); got != 2 {
		t.Fatalf("Divide(6, 3) = %d, want 2", got)
	}
}

func TestDivideByZeroPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Divide(1, 0) did not panic")
		}
	}()
	_ = Divide(1, 0)
}
