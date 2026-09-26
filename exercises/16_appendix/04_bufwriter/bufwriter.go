// Package bufwriter teaches buffered output with explicit flush.
package bufwriter

import "io"

// WriteLines writes each line plus "\n" through a buffer.
// TODO: bufio.NewWriter, WriteString per line, Flush (import bufio).
func WriteLines(w io.Writer, lines []string) error {
	return nil
}
