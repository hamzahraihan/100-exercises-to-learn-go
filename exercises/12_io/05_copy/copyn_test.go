package copyn

import (
	"bytes"
	"strings"
	"testing"
)

func TestCopyN(t *testing.T) {
	var buf bytes.Buffer
	n, err := CopyN(&buf, strings.NewReader("hello world"), 5)
	if err != nil || n != 5 || buf.String() != "hello" {
		t.Fatalf("CopyN = (%d, %q, %v), want (5, hello, nil)", n, buf.String(), err)
	}
}
