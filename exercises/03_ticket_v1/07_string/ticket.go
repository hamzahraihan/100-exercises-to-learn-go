// Package ticket teaches the String method (fmt.Stringer, implicit).
package ticket

import "fmt"

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// String describes the ticket for humans and fmt printing.
func (t Ticket) String() string {
	return fmt.Sprintf("Ticket: %s", t.Title)
}
