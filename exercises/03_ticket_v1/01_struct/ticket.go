// Package ticket models a support ticket.
package ticket

// Ticket is a support request with a title and description.
// TODO: populate and return the struct in NewTicket.
type Ticket struct {
	Title       string
	Description string
}

// NewTicket builds a Ticket.
// TODO: return Ticket{Title: title, Description: description} instead of Ticket{}.
func NewTicket(title, description string) Ticket {
	return Ticket{}
}
