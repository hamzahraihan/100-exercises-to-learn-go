// Package ticket teaches JSON unmarshaling with validation.
package ticket

import (
	"encoding/json"
	"errors"
)

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// UnmarshalTicket decodes JSON and validates: title must be non-empty.
func UnmarshalTicket(data string) (Ticket, error) {
	var t Ticket
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		return Ticket{}, err
	}
	if t.Title == "" {
		return Ticket{}, errors.New("title must be non-empty")
	}
	return t, nil
}
