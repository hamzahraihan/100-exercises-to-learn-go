// Package store teaches map idioms.
package store

// Ticket is a stored item.
type Ticket struct {
	Title string
}

// Registry indexes tickets by title.
type Registry struct {
	byTitle map[string]Ticket
}

// Put stores t. The zero Registry has a nil map: make it on first use,
// since assigning into a nil map panics.
func (r *Registry) Put(t Ticket) {
	if r.byTitle == nil {
		r.byTitle = make(map[string]Ticket)
	}
	r.byTitle[t.Title] = t
}

// Lookup finds a ticket by title with the comma-ok idiom.
func (r *Registry) Lookup(title string) (Ticket, bool) {
	t, ok := r.byTitle[title]
	return t, ok
}
