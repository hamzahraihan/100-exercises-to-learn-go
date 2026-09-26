# Surviving Corrupt Files

The loader from two exercises ago trusted the disk. The disk doesn't
deserve it. Users hand-edit state files, disks glitch, crashes interrupt
mid-write (the atomic-save lesson exists because they do) — and every one
of those produces bytes your decoder must survive. This exercise hardens
the load path: no panics on bad data, and errors that locate the damage.

## Disk is untrusted input

Treat file contents exactly like network input: hostile until proven
otherwise. That rules out every shortcut — no `panic` on decode failure
(panics are for programmer bugs, and a corrupt file is the *world's* bug),
no silent defaults that launder garbage into state, no bare errors that
say *what* broke but never *where*.

Two gates stand between the bytes and your program, and both can fail:

```go
func LoadTickets(path string) ([]Ticket, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("read %s: %w", path, err)
    }
    var ts []Ticket
    if err := json.Unmarshal(data, &ts); err != nil {
        return nil, fmt.Errorf("decode %s: %w", path, err)
    }
    return ts, nil
}
```

Read failures (missing file, permissions) and decode failures (`{oops`
isn't JSON, shapes don't match) get *separate* wraps — `read` vs `decode`
tells the operator which stage broke before they open anything. Each wrap
carries two things: the path (which file?) and the chain (`%w`, so
`errors.Is` still sees through to the cause). Context for humans, identity
for programs: the wrapping lesson's dual audience, now with a filename
attached.

## Errors that locate

The test asserts geography, not just failure:

```go
} else if !strings.Contains(err.Error(), bad) {
    t.Fatalf("error = %q, want path inside", err.Error())
}
```

A message without the path forces a hunt across every state file on the
machine; a message with it points at the patient. "Corrupt JSON" is a
complaint; "`decode /var/tickets/bad.json: invalid character 'o'`" is a
work order. When your own loaders fail in production at 3 AM, the
difference between those two messages is the difference between a fix and
an incident — so the test grades the path's presence, not just the error's
existence.

The good file still decodes cleanly through the same code: one ticket in,
one ticket out, title intact. Hardening must never tax the honest case —
validation gates the corrupt, the well-formed flows untouched.

## Task

Fill in `LoadTickets` in `loadtickets.go`:

```go
func LoadTickets(path string) ([]Ticket, error) {
    // ...
}
```

Read the file with `os.ReadFile`, decode it with `json.Unmarshal`, and wrap
every failure with the path (hint: `fmt.Errorf` with `%w`). The stub
returns `(nil, nil)`, so the test fails on both the good and corrupt cases.

## Check

```bash
go test ./exercises/14_persistence/03_corrupt/ -v
```
