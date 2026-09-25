# Slice Store

Tickets so far lived in variables. Now they need a home that grows: a
store you can add to and read back from, by id. The first engine is the
humble slice — and it teaches three lessons disguised as one exercise.

## Growing with append

```go
type Store struct {
    tickets []Ticket
    nextID  int
}

func (s *Store) Add(title string) int {
    id := s.nextID
    s.nextID++
    s.tickets = append(s.tickets, Ticket{ID: id, Title: title})
    return id
}
```

Two mechanisms cooperate. `nextID` hands out fresh identities — assigned,
then bumped, so no id repeats even as tickets accumulate. And `append`
grows the slice: note the assignment back into `s.tickets`. `append` may
return a *new* backing array when capacity runs out, so discarding its
result silently loses elements. `s.tickets = append(...)` isn't style;
it's correctness.

Both methods hang off `*Store`, pointer receivers — the setters lesson
with the stakes raised. `Add` mutates two fields; through a value receiver
both mutations would evaporate.

## Usable before birth

The test declares its store nakedly:

```go
var s Store
```

No constructor, no `NewStore`, no setup call. This works because the zero
value is already valid: `tickets` is a `nil` slice, and appending to a
`nil` slice behaves exactly like appending to an empty one. `nextID` starts
at 0, a fine first id. The zero-value design habit from the ticket section
pays off at container scale — every user of `Store` gets a working store
for free.

## Finding by scanning

```go
func (s *Store) Get(id int) (Ticket, bool) {
    for _, tk := range s.tickets {
        if tk.ID == id {
            return tk, true
        }
    }
    return Ticket{}, false
}
```

No index, no shortcut: walk every ticket and compare. The `(Ticket, bool)`
pair is the paired-returns rhythm once more — value plus verdict, with the
zero `Ticket{}` standing in when nothing matches. Callers check `ok`
before trusting the ticket, the same way they check `err` before trusting
a result.

Honest about cost: this scan touches every element, so lookups slow as the
store grows. That's not a flaw in the exercise — it's the motivation for
the next one, where the same `Add`/`Get` shape gets a faster engine.

## Task

Complete `Add` and `Get` in `store.go` so added tickets are retrievable by
id, unknown ids report missing, and the test passes.

## Check

```bash
go test ./exercises/06_ticket_store/01_slice_store/ -v
```
