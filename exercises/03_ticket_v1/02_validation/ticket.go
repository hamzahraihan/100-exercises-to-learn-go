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
// TODO: return an error when Title is empty or Description is shorter than 10 chars.
func (t Ticket) Validate() error {
	return errors.New("TODO")
}
