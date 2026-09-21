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
// TODO: lazily make r.byTitle, then store under t.Title.
func (r *Registry) Put(t Ticket) {
}

// Lookup finds a ticket by title with the comma-ok idiom.
// TODO: return r.byTitle[title].
func (r *Registry) Lookup(title string) (Ticket, bool) {
	return Ticket{}, false
}
