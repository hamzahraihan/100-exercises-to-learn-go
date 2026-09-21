// Package ticket teaches JSON struct tags.
package ticket

// Ticket is a support request. InternalNote never leaves the process.
type Ticket struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	Status       string `json:"status"`
	InternalNote string `json:"-"`
}

// MarshalTicket encodes t as JSON.
// TODO: use json.Marshal so the struct tags take effect. The hardcoded
// string below leaks InternalNote — real marshaling with json:"-" won't.
func MarshalTicket(t Ticket) (string, error) {
	return `{"id":1,"title":"t","internal_note":"leak"}`, nil
}
