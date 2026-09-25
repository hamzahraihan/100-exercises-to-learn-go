# Slice Functions (Clone)

Slices feel like values until the day two of them share a secret. Assign a
slice, mutate through one name, and watch the other change — spooky action
at a distance, fully documented, biting everyone exactly once. This
exercise is the vaccination.

## The header is not the array

A slice is a small descriptor — pointer, length, capacity — over a shared
backing array:

```go
a := []int{1, 2, 3}
b := a          // copies the header, NOT the elements
b[0] = 99       // a[0] is now 99 as well — same array underneath
```

Assignment copies three words (pointer, len, cap). Both headers point at
the same memory, so writes through either alias land in one place. `len`
says how many elements you may touch; `cap` says how many fit before
`append` must allocate a fresh array and move house — the moment aliases
quietly *stop* sharing, which is its own surprise in the other direction.

Functions receive the header by value and the array by reference, all at
once. Pass a slice in, mutate elements, and the caller sees it. The ticket
store's methods relied on this silently; today it's explicit.

## Cloning for real

Independence needs a fresh array with the elements copied over. The
standard library names the operation directly:

```go
import "slices"

func CloneTickets(ts []Ticket) []Ticket {
    return slices.Clone(ts)
}
```

`slices.Clone` allocates exactly enough room and copies every element —
mutating the result never touches the original. The manual equivalent says
the same thing longhand: append into nothing, or copy into a made slice:

```go
out := append([]Ticket(nil), ts...) // fresh array via append
// or:
out := make([]Ticket, len(ts))
copy(out, ts)
```

All three sever the alias. Prefer `slices.Clone` when it fits — named
operations beat hand-rolled ones — and reach for the manual forms when the
clone needs reshaping en route (the next exercise does exactly that).

The test proves independence the only convincing way: equal first
(`reflect.DeepEqual`, since slices reject `==` — the comparable lesson
explains why), then mutate the clone and confirm the original stands
unchanged. Equality checks content; the mutation checks sharing. Both
assertions, or the clone isn't proven.

## Task

Complete `CloneTickets` in `store.go` so the result equals the input but
shares no backing array with it, and the test passes:

```go
func CloneTickets(ts []Ticket) []Ticket {
    // ...
}
```

## Check

```bash
go test ./exercises/06_ticket_store/03_slices_fn/ -v
```
