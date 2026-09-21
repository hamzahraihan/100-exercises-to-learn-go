// Package ticket teaches the String method (fmt.Stringer, implicit).
package ticket

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// String describes the ticket for humans and fmt printing.
// TODO: include the title (hint: fmt.Sprintf is allowed here).
func (t Ticket) String() string {
	return ""
}
