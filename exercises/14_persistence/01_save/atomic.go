// Package atomic teaches crash-safe file writes.
package atomic

// Save writes data atomically: temp file in the same dir + rename, so a
// crash never leaves a half-written file.
// TODO: os.CreateTemp(dir...), Write, Close, os.Rename (import os, path/filepath).
func Save(path string, data []byte) error {
	return nil
}
