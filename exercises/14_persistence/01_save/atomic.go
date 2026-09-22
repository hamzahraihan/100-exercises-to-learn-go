// Package atomic teaches crash-safe file writes.
package atomic

import (
	"os"
	"path/filepath"
)

// Save writes data atomically: temp file in the same dir + rename, so a
// crash never leaves a half-written file.
func Save(path string, data []byte) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), "tmp-*")
	if err != nil {
		return err
	}
	name := tmp.Name()
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(name)
		return err
	}
	if err := tmp.Close(); err != nil {
		os.Remove(name)
		return err
	}
	return os.Rename(name, path)
}
