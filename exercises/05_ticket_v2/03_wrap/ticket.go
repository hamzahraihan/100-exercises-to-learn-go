// Package ticket teaches wrapping errors with %w.
package ticket

import (
	"errors"
	"fmt"
)

// ErrNotFound is returned when an id is absent.
var ErrNotFound = errors.New("not found")

// FindTicket returns the index of id, or a wrapped ErrNotFound.
func FindTicket(ids []int, id int) (int, error) {
	for i, v := range ids {
		if v == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("ticket %d: %w", id, ErrNotFound)
}
