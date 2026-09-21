package ticket

import (
	"errors"
	"testing"
)

func TestFindTicket(t *testing.T) {
	idx, err := FindTicket([]int{10, 20, 30}, 20)
	if err != nil || idx != 1 {
		t.Fatalf("FindTicket = (%d, %v), want (1, nil)", idx, err)
	}
	_, err = FindTicket([]int{10, 20}, 99)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want errors.Is(err, ErrNotFound)", err)
	}
}
