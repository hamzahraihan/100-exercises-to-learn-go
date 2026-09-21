// Package store teaches maps as indexed storage.
package store

// Ticket is a stored item.
type Ticket struct {
	ID    int
	Title string
}

// Store keeps tickets in a map.
type Store struct {
	tickets map[int]Ticket
	nextID  int
}

// Add inserts a ticket and returns its id.
// TODO: lazily make(s.tickets) when nil, then store and bump nextID.
func (s *Store) Add(title string) int {
	return 0
}

// Get finds a ticket by id using the comma-ok idiom.
// TODO: look up s.tickets[id].
func (s *Store) Get(id int) (Ticket, bool) {
	return Ticket{}, false
}
