// Package sqldb teaches database/sql with a SQLite driver.
package sqldb

import (
	"context"
	"database/sql"
	"errors"

	_ "modernc.org/sqlite" // registers the "sqlite" driver
)

// ErrNotFound is returned when a ticket ID does not exist.
var ErrNotFound = errors.New("ticket not found")

// Ticket is a row in the tickets table.
type Ticket struct {
	ID     int64
	Title  string
	Status string
}

// Store wraps a *sql.DB.
type Store struct {
	db *sql.DB
}

// NewStore opens an in-memory SQLite DB and creates the tickets table.
// TODO: sql.Open("sqlite", ":memory:"), db.ExecContext to CREATE TABLE tickets
// (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, status TEXT NOT NULL).
func NewStore(ctx context.Context) (*Store, error) {
	return nil, errors.New("TODO")
}

// Add inserts a ticket and returns its auto-generated ID.
// TODO: db.ExecContext INSERT with ? placeholders; result.LastInsertId().
func (s *Store) Add(ctx context.Context, title, status string) (int64, error) {
	return 0, errors.New("TODO")
}

// Get returns the ticket with the given id, or ErrNotFound.
// TODO: db.QueryRowContext SELECT; Scan into Ticket; map sql.ErrNoRows → ErrNotFound.
func (s *Store) Get(ctx context.Context, id int64) (Ticket, error) {
	return Ticket{}, errors.New("TODO")
}

// List returns all tickets ordered by id.
// TODO: db.QueryContext SELECT; rows.Next loop; rows.Err() check.
func (s *Store) List(ctx context.Context) ([]Ticket, error) {
	return nil, errors.New("TODO")
}
