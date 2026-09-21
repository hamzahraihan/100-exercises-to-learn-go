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
func NewTicket(title, description string) (Ticket, error) {
	if title == "" {
		return Ticket{}, errors.New("title must not be empty")
	}
	if len(description) < 10 {
		return Ticket{}, errors.New("description must be at least 10 characters")
	}
	return Ticket{Title: title, Description: description}, nil
}
