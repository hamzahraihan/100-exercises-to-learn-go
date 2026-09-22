// Package importtk teaches validated JSON import.
package importtk

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

// Import reads a JSON backup, validating every title.
func Import(path string) ([]Ticket, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var ts []Ticket
	if err := json.Unmarshal(raw, &ts); err != nil {
		return nil, err
	}
	for i, t := range ts {
		if t.Title == "" {
			return nil, fmt.Errorf("ticket %d: empty title", i)
		}
	}
	return ts, nil
}
