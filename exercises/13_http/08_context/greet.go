// Package greetctx teaches request contexts.
package greetctx

import (
	"context"
	"time"
)

// Greet greets after a 10ms simulated lookup, aborting on cancellation.
func Greet(ctx context.Context, name string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(10 * time.Millisecond):
		return "hi " + name, nil
	}
}
