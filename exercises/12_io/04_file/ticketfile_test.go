package ticketfile

import (
	"path/filepath"
	"testing"
)

func TestSaveLoadTicket(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ticket.json")
	if err := SaveTicket(path, []byte(`{"id":1}`)); err != nil {
		t.Fatalf("SaveTicket errored: %v", err)
	}
	got, err := LoadTicket(path)
	if err != nil {
		t.Fatalf("LoadTicket errored: %v", err)
	}
	if string(got) != `{"id":1}` {
		t.Fatalf("LoadTicket = %q, want %q", got, `{"id":1}`)
	}
}
