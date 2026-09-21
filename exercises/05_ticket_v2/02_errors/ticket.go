// Package ticket teaches sentinel + wrapped errors (errors.Is/As).
package ticket

import "errors"

// Status is a ticket state.
type Status int

const (
	StatusOpen Status = iota
	StatusInProgress
	StatusClosed
)

// ErrUnknownStatus is returned for bad status values.
var ErrUnknownStatus = errors.New("unknown status")

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
	Status      Status
}

// NewTicketWithStatus builds a Ticket or fails.
// TODO: validate title/description like 03_ticket_v1 and wrap
// ErrUnknownStatus with %w when status is unknown.
func NewTicketWithStatus(title, description string, status Status) (Ticket, error) {
	return Ticket{}, errors.New("TODO")
}
