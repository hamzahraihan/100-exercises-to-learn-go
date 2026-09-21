package welcome

import "testing"

func TestMessage(t *testing.T) {
	got := Message()
	if got == "" {
		t.Fatal("Message() returned empty string, expected a greeting")
	}
}
