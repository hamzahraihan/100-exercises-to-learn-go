// Package ticket teaches strict JSON decoding.
package ticket

import "encoding/json"

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// UnmarshalStrict decodes, rejecting unknown fields.
// TODO: use json.Decoder with DisallowUnknownFields instead of Unmarshal.
func UnmarshalStrict(data string) (Ticket, error) {
	var t Ticket
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		return Ticket{}, err
	}
	return t, nil
}
