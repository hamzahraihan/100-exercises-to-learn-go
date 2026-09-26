// Package bufwriter teaches buffered output with explicit flush.
package bufwriter

import (
	"bufio"
	"io"
)

// WriteLines writes each line plus "\n" through a buffer, flushing once.
func WriteLines(w io.Writer, lines []string) error {
	bw := bufio.NewWriter(w)
	for _, s := range lines {
		if _, err := bw.WriteString(s + "\n"); err != nil {
			return err
		}
	}
	return bw.Flush()
}
