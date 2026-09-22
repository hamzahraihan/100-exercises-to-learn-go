// Package exporttk teaches JSON backup export.
package exporttk

import (
	"encoding/json"
	"os"
)

// Ticket is a stored support request.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Export writes ts as a JSON array to path.
func Export(path string, ts []Ticket) error {
	out, err := json.Marshal(ts)
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0o644)
}
