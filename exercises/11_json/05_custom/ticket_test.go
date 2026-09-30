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

	out, err = MarshalTicket(Ticket{ID: 2, Title: "t", Status: StatusClosed})
	if err != nil {
		t.Fatalf("MarshalTicket(closed) errored: %v", err)
	}
	if !strings.Contains(out, `"status":"closed"`) {
		t.Fatalf("output = %s, want status as \"closed\"", out)
	}
	if strings.Contains(out, `"status":"open"`) {
		t.Fatalf("closed output leaked open status: %s", out)
	}

	if _, err := MarshalTicket(Ticket{ID: 3, Title: "t", Status: Status(99)}); err == nil {
		t.Fatal("unknown status accepted, want error")
	}
}
