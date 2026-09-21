package divide

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	if got, err := Divide(6, 3); err != nil || got != 2 {
		t.Fatalf("Divide(6, 3) = (%d, %v), want (2, nil)", got, err)
	}
	if _, err := Divide(1, 0); !errors.Is(err, ErrZeroDivisor) {
		t.Fatalf("Divide(1, 0) err = %v, want ErrZeroDivisor", err)
	}
}
