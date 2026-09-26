// Package status teaches custom JSON decoding.
package status

import (
	"encoding/json"
	"fmt"
)

// Status is the ticket state.
type Status int

const (
	StatusOpen Status = iota
	StatusClosed
)

// Ticket carries a typed status.
type Ticket struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}

// UnmarshalJSON decodes "open"/"closed", rejecting anything else.
func (s *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open":
		*s = StatusOpen
	case "closed":
		*s = StatusClosed
	default:
		return fmt.Errorf("unknown status %q", v)
	}
	return nil
}
