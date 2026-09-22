package greetctx

import (
	"context"
	"testing"
)

func TestGreet(t *testing.T) {
	got, err := Greet(context.Background(), "bob")
	if err != nil || got != "hi bob" {
		t.Fatalf("Greet = (%q, %v), want (hi bob, nil)", got, err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Greet(ctx, "bob"); err == nil {
		t.Fatal("Greet(canceled) = nil error, want cancellation")
	}
}
