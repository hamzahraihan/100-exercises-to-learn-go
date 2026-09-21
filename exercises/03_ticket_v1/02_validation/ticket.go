// Package ticket models a support ticket with validation.
package ticket

import "errors"

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// NewTicket builds a Ticket.
func NewTicket(title, description string) Ticket {
	return Ticket{Title: title, Description: description}
}

// Validate checks the ticket.
func (t Ticket) Validate() error {
	if t.Title == "" {
		return errors.New("title must not be empty")
	}
	if len(t.Description) < 10 {
		return errors.New("description must be at least 10 characters")
	}
	return nil
}
