package ticket

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	if err := NewTicket("ok title", "long enough description").Validate(); err != nil {
		t.Fatalf("valid ticket rejected: %v", err)
	}
	if err := NewTicket("", "long enough description").Validate(); err == nil {
		t.Fatal("empty title accepted")
	}
	if err := NewTicket("ok", "x").Validate(); err == nil || !strings.Contains(err.Error(), "description") {
		t.Fatalf("short description should error mentioning description, got %v", err)
	}
}
