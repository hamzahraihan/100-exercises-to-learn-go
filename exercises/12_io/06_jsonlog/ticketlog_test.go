package ticketlog

import (
	"path/filepath"
	"testing"
)

func TestTicketLog(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tickets.jsonl")
	if err := AppendLog(path, Ticket{ID: 1, Title: "a"}); err != nil {
		t.Fatalf("AppendLog errored: %v", err)
	}
	if err := AppendLog(path, Ticket{ID: 2, Title: "b"}); err != nil {
		t.Fatalf("AppendLog errored: %v", err)
	}
	got, err := ReadLog(path)
	if err != nil {
		t.Fatalf("ReadLog errored: %v", err)
	}
	if len(got) != 2 || got[0].ID != 1 || got[1].ID != 2 {
		t.Fatalf("ReadLog = %+v, want 2 tickets in order", got)
	}
}
