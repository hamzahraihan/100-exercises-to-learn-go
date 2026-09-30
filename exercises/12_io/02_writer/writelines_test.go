package writelines

import (
	"bytes"
	"errors"
	"testing"
)

type failWriter struct{ err error }

func (f failWriter) Write([]byte) (int, error) { return 0, f.err }

func TestWriteLines(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteLines(&buf, []string{"a", "b"}); err != nil {
		t.Fatalf("WriteLines errored: %v", err)
	}
	if buf.String() != "a\nb\n" {
		t.Fatalf("buffer = %q, want %q", buf.String(), "a\nb\n")
	}

	var empty bytes.Buffer
	if err := WriteLines(&empty, nil); err != nil {
		t.Fatalf("WriteLines(nil) errored: %v", err)
	}
	if empty.String() != "" {
		t.Fatalf("WriteLines(nil) = %q, want empty", empty.String())
	}

	sentinel := errors.New("write boom")
	if err := WriteLines(failWriter{err: sentinel}, []string{"a"}); !errors.Is(err, sentinel) {
		t.Fatalf("WriteLines(failing writer) = %v, want %v", err, sentinel)
	}
}
