// Package ticket teaches struct embedding.
package ticket

// Meta holds fields shared by several domain types.
type Meta struct {
	ID int
}

// Ticket embeds Meta: ID is promoted, so tk.ID works directly.
type Ticket struct {
	Meta
	Title string
}

// NewTicketWithID builds a Ticket with an id.
func NewTicketWithID(id int, title string) Ticket {
	return Ticket{Meta: Meta{ID: id}, Title: title}
}
