// Package ticket models a support ticket.
package ticket

// Ticket is a support request with a title and description.
type Ticket struct {
	Title       string
	Description string
}

// NewTicket builds a Ticket.
func NewTicket(title, description string) Ticket {
	return Ticket{Title: title, Description: description}
}
