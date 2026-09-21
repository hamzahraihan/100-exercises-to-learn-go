package ticket

import "testing"

func TestNewTicketValid(t *testing.T) {
	tk, err := NewTicket("Fix bug", "Crash on login screen")
	if err != nil {
		t.Fatalf("valid ticket rejected: %v", err)
	}
	if tk.Title != "Fix bug" {
		t.Fatalf("Title = %q", tk.Title)
	}
}

func TestNewTicketInvalid(t *testing.T) {
	if _, err := NewTicket("", "long enough description"); err == nil {
		t.Fatal("empty title accepted")
	}
	if _, err := NewTicket("ok title", "short"); err == nil {
		t.Fatal("short description accepted")
	}
}
