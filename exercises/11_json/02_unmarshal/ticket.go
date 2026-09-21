// Package ticket teaches JSON unmarshaling with validation.
package ticket

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// UnmarshalTicket decodes JSON and validates: title must be non-empty.
// TODO: json.Unmarshal into Ticket, then check Title (import encoding/json).
func UnmarshalTicket(data string) (Ticket, error) {
	return Ticket{}, nil
}
