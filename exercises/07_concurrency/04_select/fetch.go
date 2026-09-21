// Package fetch teaches select with timeouts.
package fetch

import (
	"errors"
	"time"
)

// Fetch simulates a fetch taking d. It must fail when d exceeds 200ms.
func Fetch(d time.Duration) (string, error) {
	select {
	case <-time.After(d):
		return "data", nil
	case <-time.After(200 * time.Millisecond):
		return "", errors.New("fetch timed out")
	}
}
