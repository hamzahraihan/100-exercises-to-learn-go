// Package ticketfile teaches file reads and writes.
package ticketfile

import "os"

// SaveTicket writes data to path.
func SaveTicket(path string, data []byte) error {
	return os.WriteFile(path, data, 0o644)
}

// LoadTicket reads path back (already correct).
func LoadTicket(path string) ([]byte, error) {
	return os.ReadFile(path)
}
