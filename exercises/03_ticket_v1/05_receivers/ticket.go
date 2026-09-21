// Package ticket teaches value receivers that return modified copies.
package ticket

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// Renamed returns a copy of t with a new title, leaving t untouched.
// TODO: set t.Title and return t.
func (t Ticket) Renamed(newTitle string) Ticket {
	return Ticket{}
}
