package ticket

import (
	"errors"
	"testing"
)

func TestNewTicketWithStatusRejectsUnknown(t *testing.T) {
	_, err := NewTicketWithStatus("t", "long enough description", 99)
	if !errors.Is(err, ErrUnknownStatus) {
		t.Fatalf("err = %v, want errors.Is(err, ErrUnknownStatus)", err)
	}
}

func TestNewTicketWithStatusOK(t *testing.T) {
	tk, err := NewTicketWithStatus("t", "long enough description", StatusOpen)
	if err != nil || tk.Status != StatusOpen {
		t.Fatalf("got %+v, %v", tk, err)
	}
}
