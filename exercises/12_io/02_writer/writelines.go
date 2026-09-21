// Package writelines teaches the io.Writer contract.
package writelines

import "io"

// WriteLines writes each line plus "\n".
// TODO: range over lines writing s + "\n" (Fprintf helps).
func WriteLines(w io.Writer, lines []string) error {
	return nil
}
