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
func (s *Store) Add(title string) int {
	id := s.nextID
	s.tickets = append(s.tickets, Ticket{ID: id, Title: title})
	s.nextID++
	return id
}

// Get finds a ticket by id.
func (s *Store) Get(id int) (Ticket, bool) {
	for _, t := range s.tickets {
		if t.ID == id {
			return t, true
		}
	}
	return Ticket{}, false
}
