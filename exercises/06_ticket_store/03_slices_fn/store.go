// Package store teaches slice cloning.
package store

import "slices"

// Ticket is a stored item.
type Ticket struct {
	ID     int
	Title  string
	Closed bool
}

// CloneTickets returns an independent copy: mutating the clone must not
// affect the original.
func CloneTickets(ts []Ticket) []Ticket {
	return slices.Clone(ts)
}
