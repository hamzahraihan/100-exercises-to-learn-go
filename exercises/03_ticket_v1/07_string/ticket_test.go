package ticket

import (
	"fmt"
	"strings"
	"testing"
)

func TestTicketString(t *testing.T) {
	tk := Ticket{Title: "bug", Description: "long enough description"}
	if got := fmt.Sprint(tk); !strings.Contains(got, "bug") {
		t.Fatalf("String() = %q, want title inside", got)
	}
}
