// Package greetctx teaches request contexts.
package greetctx

import (
	"context"
	"time"
)

// Greet greets after a 10ms simulated lookup, aborting on cancellation.
// TODO: select on ctx.Done() (return ctx.Err()) vs time.After result.
func Greet(ctx context.Context, name string) (string, error) {
	<-time.After(10 * time.Millisecond)
	return "hi " + name, nil
}
