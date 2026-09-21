package readall

import (
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	got, err := ReadAll(strings.NewReader("hello"))
	if err != nil || got != "hello" {
		t.Fatalf("ReadAll = (%q, %v), want (hello, nil)", got, err)
	}
}
