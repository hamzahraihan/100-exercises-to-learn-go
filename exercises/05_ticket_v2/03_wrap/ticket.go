// Package ticket teaches wrapping errors with %w.
package ticket

import "errors"

// ErrNotFound is returned when an id is absent.
var ErrNotFound = errors.New("not found")

// FindTicket returns the index of id, or a wrapped ErrNotFound.
// TODO: loop; on miss return -1 and fmt.Errorf("ticket %d: %w", id, ErrNotFound).
func FindTicket(ids []int, id int) (int, error) {
	return -1, errors.New("TODO")
}
