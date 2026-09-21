// Package store teaches slices as growable storage.
package store

// Ticket is a stored item.
type Ticket struct {
	ID    int
	Title string
}

// Store keeps tickets in a slice.
type Store struct {
	tickets []Ticket
	nextID  int
}

// Add appends a ticket and returns its id.
// TODO: use append, assign s.nextID then increment it.
func (s *Store) Add(title string) int {
	return 0
}

// Get finds a ticket by id.
// TODO: range over s.tickets.
func (s *Store) Get(id int) (Ticket, bool) {
	return Ticket{}, false
}
