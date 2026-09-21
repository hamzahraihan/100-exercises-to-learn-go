package ticket

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestTags(t *testing.T) {
	out, err := MarshalTicket(Ticket{ID: 1, Title: "t", InternalNote: "leak"})
	if err != nil {
		t.Fatalf("MarshalTicket errored: %v", err)
	}
	if strings.Contains(out, "leak") {
		t.Fatalf("secret leaked into JSON: %s", out)
	}
	var back Ticket
	if err := json.Unmarshal([]byte(out), &back); err != nil || back.Title != "t" {
		t.Fatalf("round-trip = (%+v, %v)", back, err)
	}
}
