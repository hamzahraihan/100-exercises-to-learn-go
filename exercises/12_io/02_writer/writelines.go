// Package writelines teaches the io.Writer contract.
package writelines

import (
	"fmt"
	"io"
)

// WriteLines writes each line plus "\n".
func WriteLines(w io.Writer, lines []string) error {
	for _, s := range lines {
		if _, err := fmt.Fprintf(w, "%s\n", s); err != nil {
			return err
		}
	}
	return nil
}
