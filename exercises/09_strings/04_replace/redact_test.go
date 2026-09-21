package redact

import "testing"

func TestRedact(t *testing.T) {
	if got, want := Redact("my secret here"), "my [redacted] here"; got != want {
		t.Fatalf("Redact = %q, want %q", got, want)
	}
	if got, want := Redact("nothing to hide"), "nothing to hide"; got != want {
		t.Fatalf("Redact = %q, want %q", got, want)
	}
}
