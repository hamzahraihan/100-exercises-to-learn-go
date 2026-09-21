// Package ticketlog teaches JSON-lines logs.
package ticketlog

import (
	"bufio"
	"encoding/json"
	"os"
)

// Ticket is a logged support request.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// AppendLog appends one ticket as a JSON line (creating the file).
// TODO: open with O_APPEND|O_CREATE|O_WRONLY, encode with json.Encoder.
func AppendLog(path string, t Ticket) error {
	return nil
}

// ReadLog decodes every JSON line, in order.
func ReadLog(path string) ([]Ticket, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []Ticket
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var t Ticket
		if err := json.Unmarshal(sc.Bytes(), &t); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, sc.Err()
}
