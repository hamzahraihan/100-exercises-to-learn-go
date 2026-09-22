// Package loadj teaches graceful corrupt-file handling.
package loadj

import (
	"encoding/json"
	"fmt"
	"os"
)

// Ticket is a stored support request.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// LoadTickets decodes path; failures name the path (never panic).
func LoadTickets(path string) ([]Ticket, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("load %s: %w", path, err)
	}
	var ts []Ticket
	if err := json.Unmarshal(raw, &ts); err != nil {
		return nil, fmt.Errorf("load %s: %w", path, err)
	}
	return ts, nil
}
