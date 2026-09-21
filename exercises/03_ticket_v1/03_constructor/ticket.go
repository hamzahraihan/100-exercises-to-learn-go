// Package ticket teaches validated constructors.
package ticket

import "errors"

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// NewTicket builds a validated Ticket: non-empty title, description of at
// least 10 characters.
// TODO: return Ticket{...}, nil when valid, else a descriptive error.
func NewTicket(title, description string) (Ticket, error) {
	return Ticket{}, errors.New("TODO")
}
