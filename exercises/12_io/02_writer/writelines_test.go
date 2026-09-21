package writelines

import (
	"bytes"
	"testing"
)

func TestWriteLines(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteLines(&buf, []string{"a", "b"}); err != nil {
		t.Fatalf("WriteLines errored: %v", err)
	}
	if buf.String() != "a\nb\n" {
		t.Fatalf("buffer = %q, want %q", buf.String(), "a\nb\n")
	}
}
