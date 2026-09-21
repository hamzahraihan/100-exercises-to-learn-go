// Package store teaches sorting and filtering slices.
package store

import (
	"cmp"
	"slices"
)

// Ticket is a stored item.
type Ticket struct {
	ID     int
	Title  string
	Closed bool
}

// SortedByID returns tickets ordered by ID ascending.
func SortedByID(ts []Ticket) []Ticket {
	out := slices.Clone(ts)
	slices.SortFunc(out, func(a, b Ticket) int { return cmp.Compare(a.ID, b.ID) })
	return out
}

// OpenOnly keeps tickets that are not closed.
func OpenOnly(ts []Ticket) []Ticket {
	var out []Ticket
	for _, t := range ts {
		if !t.Closed {
			out = append(out, t)
		}
	}
	return out
}
