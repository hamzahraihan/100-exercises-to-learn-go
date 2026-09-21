// Package store teaches slice cloning.
package store

// Ticket is a stored item.
type Ticket struct {
	ID     int
	Title  string
	Closed bool
}

// CloneTickets returns an independent copy: mutating the clone must not
// affect the original.
// TODO: append into a fresh slice (or use the slices package Clone).
func CloneTickets(ts []Ticket) []Ticket {
	return nil
}
