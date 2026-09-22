// Package loadst teaches boot-time loading.
package loadst

import "os"

// Load reads path; a missing file means empty state, not an error.
// TODO: os.ReadFile, map os.IsNotExist(err) to (nil, nil)
// (import errors only if needed for errors.Is style).
func Load(path string) ([]byte, error) {
	return os.ReadFile(path)
}
