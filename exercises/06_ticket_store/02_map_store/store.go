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
func (s *Store) Add(title string) int {
	if s.tickets == nil {
		s.tickets = make(map[int]Ticket)
	}
	id := s.nextID
	s.tickets[id] = Ticket{ID: id, Title: title}
	s.nextID++
	return id
}

// Get finds a ticket by id using the comma-ok idiom.
func (s *Store) Get(id int) (Ticket, bool) {
	t, ok := s.tickets[id]
	return t, ok
}
