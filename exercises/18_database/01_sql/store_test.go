package sqldb

import (
	"context"
	"errors"
	"testing"
)

func TestNewStore(t *testing.T) {
	ctx := context.Background()
	store, err := NewStore(ctx)
	if err != nil {
		t.Fatalf("got error %v, want nil (hint: sql.Open then db.ExecContext CREATE TABLE)", err)
	}
	if store == nil {
		t.Fatal("got nil store, want non-nil *Store")
	}
}

func TestAddAndGet(t *testing.T) {
	ctx := context.Background()
	store, err := NewStore(ctx)
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}

	id, err := store.Add(ctx, "Bug", "open")
	if err != nil {
		t.Fatalf("Add: got error %v, want nil (hint: ExecContext INSERT, then result.LastInsertId())", err)
	}
	if id <= 0 {
		t.Fatalf("Add: got id %d, want id > 0 (hint: AUTOINCREMENT starts at 1)", id)
	}

	got, err := store.Get(ctx, id)
	if err != nil {
		t.Fatalf("Get(%d): got error %v, want nil (hint: QueryRowContext then Scan)", id, err)
	}
	want := Ticket{ID: id, Title: "Bug", Status: "open"}
	if got != want {
		t.Fatalf("Get(%d): got %+v, want %+v (hint: Scan columns in SELECT order: id, title, status)", id, got, want)
	}
}

func TestGetNotFound(t *testing.T) {
	ctx := context.Background()
	store, err := NewStore(ctx)
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}

	_, err = store.Get(ctx, 9999)
	if err == nil {
		t.Fatal("Get(9999): got nil error, want ErrNotFound (hint: check errors.Is(err, sql.ErrNoRows) and return ErrNotFound)")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Get(9999): got error %v, want ErrNotFound (hint: wrap or return ErrNotFound when sql.ErrNoRows)", err)
	}
}

func TestList(t *testing.T) {
	ctx := context.Background()
	store, err := NewStore(ctx)
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}

	if _, err := store.Add(ctx, "Alpha", "open"); err != nil {
		t.Fatalf("Add Alpha: %v", err)
	}
	if _, err := store.Add(ctx, "Beta", "closed"); err != nil {
		t.Fatalf("Add Beta: %v", err)
	}

	tickets, err := store.List(ctx)
	if err != nil {
		t.Fatalf("List: got error %v, want nil (hint: QueryContext + rows.Next loop + rows.Err())", err)
	}
	if len(tickets) != 2 {
		t.Fatalf("List: got %d tickets, want 2", len(tickets))
	}

	titles := map[string]bool{}
	for _, tk := range tickets {
		titles[tk.Title] = true
	}
	for _, want := range []string{"Alpha", "Beta"} {
		if !titles[want] {
			t.Fatalf("List: missing ticket with title %q (hint: Scan all three columns in loop)", want)
		}
	}
}

func TestListEmpty(t *testing.T) {
	ctx := context.Background()
	store, err := NewStore(ctx)
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}

	tickets, err := store.List(ctx)
	if err != nil {
		t.Fatalf("List on empty store: got error %v, want nil", err)
	}
	if len(tickets) != 0 {
		t.Fatalf("List on empty store: got %d tickets, want 0", len(tickets))
	}
	// Ensure callers can safely range over an empty result (non-nil slice preferred,
	// but len==0 is the contract; this loop simply must not panic).
	for _, tk := range tickets {
		_ = tk
	}
}
