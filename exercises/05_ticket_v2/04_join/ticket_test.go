package ticket

import (
	"errors"
	"testing"
)

func TestValidateAll(t *testing.T) {
	if err := ValidateAll("Ticket title", "long enough description"); err != nil {
		t.Fatalf("valid input rejected: %v", err)
	}
	err := ValidateAll("", "short")
	if !errors.Is(err, ErrBadTitle) || !errors.Is(err, ErrBadDesc) {
		t.Fatalf("err = %v, want both ErrBadTitle and ErrBadDesc", err)
	}
	err = ValidateAll("", "long enough description")
	if !errors.Is(err, ErrBadTitle) || errors.Is(err, ErrBadDesc) {
		t.Fatalf("err = %v, want only ErrBadTitle", err)
	}
}
