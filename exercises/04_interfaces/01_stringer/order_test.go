package order

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringer(t *testing.T) {
	o := Order{ID: 7, Item: "book"}
	if got := fmt.Sprint(o); !strings.Contains(got, "7") || !strings.Contains(got, "book") {
		t.Fatalf("String() = %q, want id and item inside", got)
	}
}
