// Package getbody teaches outgoing requests.
package getbody

import (
	"io"
	"net/http"
)

// GetBody fetches url and returns its body as a string.
func GetBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
