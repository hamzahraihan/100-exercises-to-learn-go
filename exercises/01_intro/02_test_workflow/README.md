# The test workflow

The previous two exercises handed you tests to run. This one makes you
*read* one. Open `add_test.go` before touching `add.go` — for the rest of
this course, the test file is the specification and the stub is just the
blank to fill in.

## Anatomy of a table test

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {1, 2, 3},
        {-1, 1, 0},
        {0, 0, 0},
    }
    for _, tt := range tests {
        if got := Add(tt.a, tt.b); got != tt.want {
            t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
        }
    }
}
```

Three ideas packed into eighteen lines. Let's take them apart.

**A slice of cases, not three separate tests.** The `tests` variable holds
every input/output pair in one table. Adding a case means adding one line —
`{2, 3, 5}` — instead of writing a whole new function. When a bug report
arrives six months from now, it becomes a row in this table first, a fix
second. That's the habit this shape teaches.

**An anonymous struct as the row type.** `struct { a, b, want int }` is
declared inline, used once, never named. Go lets types appear exactly where
they're needed — no ceremony of inventing a `TestCase` type for something
with a three-line lifespan. You'll see this frugality everywhere: slices of
anonymous structs for tables, anonymous functions for callbacks.

**One loop that reports every failure.** `for _, tt := range tests` walks
the table; `_` discards the index nobody needs. Each mismatch calls
`t.Errorf`, which marks the test failed *and keeps going*, so a single run
shows all broken cases at once.

## `Errorf` vs `Fatalf`

Two failure functions, different temperaments:

```go
t.Errorf("...") // record failure, continue with the next case
t.Fatalf("...") // record failure, stop this test function now
```

`Errorf` suits tables — you want the full damage report. `Fatalf` suits
setup — if the fixture didn't load, further assertions would only pile
confusion onto confusion. The welcome exercise used `Fatalf` for a single
irrefutable check; here `Errorf` collects all three verdicts. Match the
function to the situation and your test output reads like a diagnosis
instead of a riddle.

## Running them

One exercise, verbosely:

```bash
go test ./exercises/01_intro/02_test_workflow/ -v
```

The `-v` flag names each test as it runs. Later, when a package holds a
dozen tables, `-run TestAdd` filters to just the one you care about:

```bash
go test ./exercises/01_intro/02_test_workflow/ -run TestAdd -v
```

From here on, this loop is the course: read the table, make every row
green, move on. The tests are strict but fair — they check behavior, never
implementation, so any `Add` returning the right sums passes.

## Task

Fix `Add` in `add.go`.

## Check

```bash
go test ./exercises/01_intro/02_test_workflow/ -v
```
