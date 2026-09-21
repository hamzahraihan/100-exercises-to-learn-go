package ticket

import (
	"strings"
	"testing"
)

func TestCustomStatus(t *testing.T) {
	out, err := MarshalTicket(Ticket{ID: 1, Title: "t", Status: StatusOpen})
	if err != nil {
		t.Fatalf("MarshalTicket errored: %v", err)
	}
	if !strings.Contains(out, `"status":"open"`) {
		t.Fatalf("output = %s, want status as \"open\"", out)
	}
}
