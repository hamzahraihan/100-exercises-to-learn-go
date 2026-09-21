// Package expiry teaches testable time via a fake clock.
package expiry

import "time"

// Now is the clock. Tests replace it with a fixed time; production
// leaves it as time.Now.
// TODO: keep this variable; use Now() inside IsExpired.
var Now = time.Now

// IsExpired reports whether expiry is in the past per Now().
// TODO: return Now().After(expiry).
func IsExpired(expiry time.Time) bool {
	return false
}
