// Package poem teaches testdata fixtures.
package poem

import (
	"os"
	"strings"
)

// FirstLine returns the first line of the file at path.
func FirstLine(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	parts := strings.SplitN(string(data), "\n", 2)
	return strings.TrimSuffix(parts[0], "\r"), nil
}
