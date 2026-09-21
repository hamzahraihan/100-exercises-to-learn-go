// Package ticket teaches designing for useful zero values.
package ticket

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// IsZero reports whether t is the zero value.
func (t Ticket) IsZero() bool {
	return t.Title == "" && t.Description == ""
}
