// Package show teaches the any alias for interface{}.
package show

import "fmt"

// SprintAny formats any value.
func SprintAny(v any) string {
	return fmt.Sprint(v)
}
