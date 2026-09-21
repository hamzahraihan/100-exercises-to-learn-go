// Package fetch teaches select with timeouts.
package fetch

import "time"

// Fetch simulates a fetch taking d. It must fail when d exceeds 200ms.
// TODO: select on time.After(d) for success vs time.After(200ms) for timeout.
func Fetch(d time.Duration) (string, error) {
	<-time.After(d)
	return "data", nil
}
