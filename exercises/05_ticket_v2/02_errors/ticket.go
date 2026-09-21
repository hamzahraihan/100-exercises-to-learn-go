// Package ticket teaches sentinel + wrapped errors (errors.Is/As).
package ticket

import (
	"errors"
	"fmt"
)

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
func NewTicketWithStatus(title, description string, status Status) (Ticket, error) {
	if title == "" {
		return Ticket{}, errors.New("title must not be empty")
	}
	if len(description) < 10 {
		return Ticket{}, errors.New("description must be at least 10 characters")
	}
	switch status {
	case StatusOpen, StatusInProgress, StatusClosed:
		return Ticket{Title: title, Description: description, Status: status}, nil
	default:
		return Ticket{}, fmt.Errorf("unknown status %d: %w", int(status), ErrUnknownStatus)
	}
}
