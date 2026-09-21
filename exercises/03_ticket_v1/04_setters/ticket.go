// Package ticket teaches pointer-receiver setters.
package ticket

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// SetTitle changes the title. A pointer receiver mutates the caller's copy;
// a value receiver would mutate a throwaway copy.
func (t *Ticket) SetTitle(title string) {
	t.Title = title
}
