// Package loadst teaches boot-time loading.
package loadst

import "os"

// Load reads path; a missing file means empty state, not an error.
func Load(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}
