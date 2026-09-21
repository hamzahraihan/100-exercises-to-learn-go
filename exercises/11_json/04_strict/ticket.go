// Package ticket teaches strict JSON decoding.
package ticket

import (
	"encoding/json"
	"strings"
)

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// UnmarshalStrict decodes, rejecting unknown fields.
func UnmarshalStrict(data string) (Ticket, error) {
	var t Ticket
	dec := json.NewDecoder(strings.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&t); err != nil {
		return Ticket{}, err
	}
	return t, nil
}
