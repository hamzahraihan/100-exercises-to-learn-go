# SQL Database

`database/sql` is Go's stdlib database layer — a thin interface over any relational database. Switch drivers (SQLite → Postgres → MySQL) by changing one import; the rest of your code stays put. This exercise wires up the essentials: open, create, insert, query.

## Open is not connect

```go
db, err := sql.Open("sqlite", ":memory:")
```

`Open` validates the driver name and DSN format; the actual connection is lazy. Call `db.Ping()` or issue a query to force it. `:memory:` gives an in-process SQLite database — no file, no cleanup, perfect for tests.

## Schema and writes with ExecContext

```go
_, err = db.ExecContext(ctx, `CREATE TABLE tickets (
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    title   TEXT NOT NULL,
    status  TEXT NOT NULL
)`)
_, err = db.ExecContext(ctx, `INSERT INTO tickets (title, status) VALUES (?, ?)`, title, status)
```

`ExecContext` runs statements that return no rows. Always use `?` placeholders — never string-concatenate values (SQL injection). The `Context` variants respect cancellation; prefer them over the bare `Exec`.

## Single-row query with QueryRowContext

```go
var t Ticket
err = db.QueryRowContext(ctx, `SELECT id, title, status FROM tickets WHERE id = ?`, id).
    Scan(&t.ID, &t.Title, &t.Status)
if errors.Is(err, sql.ErrNoRows) {
    return Ticket{}, ErrNotFound
}
```

`QueryRowContext` returns exactly one row; `Scan` fills the variables in column order. `sql.ErrNoRows` is the sentinel for empty results — check it explicitly before treating the error as fatal.

## Multi-row query with rows.Next

```go
rows, err := db.QueryContext(ctx, `SELECT id, title, status FROM tickets`)
if err != nil { return nil, err }
defer rows.Close()
var tickets []Ticket
for rows.Next() {
    var t Ticket
    if err := rows.Scan(&t.ID, &t.Title, &t.Status); err != nil {
        return nil, err
    }
    tickets = append(tickets, t)
}
return tickets, rows.Err() // rows.Err() surfaces iterator errors
```

Always `defer rows.Close()` and check `rows.Err()` after the loop — the iterator can fail mid-stream.

<details><summary>Hint</summary>

Open with `sql.Open("sqlite", ":memory:")`, use `db.ExecContext` for DDL and writes, `db.QueryRowContext(...).Scan` for single-row reads, `db.QueryContext` + `rows.Next` loop for lists.

For `Add`, retrieve the auto-generated ID with `result.LastInsertId()` from the return value of `ExecContext`.

For `List`, return an empty `[]Ticket{}` (not `nil`) when there are no rows so callers can safely range over it.

</details>

## Setup

Install the pure-Go SQLite driver (no CGO required):

```bash
go get modernc.org/sqlite
```

## Task

In `store.go`, implement the four functions:

1. **`NewStore(ctx context.Context) (*Store, error)`** — open an in-memory SQLite DB with `sql.Open("sqlite", ":memory:")`, create the `tickets` table, return a `*Store`.
2. **`(s *Store) Add(ctx context.Context, title, status string) (int64, error)`** — insert a ticket using `?` placeholders, return the new row's auto-generated ID.
3. **`(s *Store) Get(ctx context.Context, id int64) (Ticket, error)`** — query a single ticket by id; return `ErrNotFound` when no row matches.
4. **`(s *Store) List(ctx context.Context) ([]Ticket, error)`** — return all tickets ordered by id; return an empty (non-nil) slice when the table is empty.

## Check

```bash
go test ./exercises/18_database/01_sql/ -v
```
