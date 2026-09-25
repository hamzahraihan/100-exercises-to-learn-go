# Table-Driven Tests

You've *run* dozens of table tests by now. This is the exercise where you
learn to *read* one as a design — because from here on, every test you
write yourself will wear this shape.

## The pattern, named

```go
func TestGrade(t *testing.T) {
    tests := []struct {
        score int
        want  string
    }{
        {95, "A"},
        {85, "B"},
        {75, "C"},
        {59, "F"},
    }
    for _, tt := range tests {
        if got := Grade(tt.score); got != tt.want {
            t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
        }
    }
}
```

Table, loop, verdict. The table lists input/output pairs as data — adding
a case means adding a line, never a function. The loop runs every row. The
verdict compares and reports. Three moving parts, endless reuse: this one
shape covers the overwhelming majority of Go unit tests in the wild.

## Why Errorf, not Fatalf

The verdict calls `t.Errorf` — record and *continue*. With the stub
returning `""` for everything, all four rows fail in a single run, and you
see all four failures at once. `t.Fatalf` would stop at 95 and hide the
rest, turning one run into four debug cycles. Tables collect evidence;
`Errorf` is how. Reserve `Fatalf` for setup that poisons everything after
it — a fixture that didn't load, a server that didn't start — where
continuing only manufactures confusion.

## Order the guards, mind the edges

```go
func Grade(score int) string {
    if score >= 90 {
        return "A"
    } else if score >= 80 {
        return "B"
    } else if score >= 70 {
        return "C"
    }
    return "F"
}
```

Descending cutoffs, each branch handling everything above its line — by
the time `score >= 80` runs, 90+ has already returned, so no upper bound
needs stating. Ascending order would demand compound conditions (`score >=
70 && score < 80`) and invite boundary slips. Let the chain's structure
carry the exclusions.

And the table's rows are chosen, not sampled: 95/85/75 sit mid-band, 59
fails clearly. When you write your own tables later, add the fence-sitters
— exactly 90, exactly 70 — because grading bugs live on boundaries, and
rows are cheap.

## Task

Fill in `Grade` in `grade.go` so each score maps to its letter:

```go
func Grade(score int) string {
    // ...
}
```

Use an if/else chain on the 90/80/70 cutoffs (90+ is "A", 80+ is "B", 70+
is "C", anything below is "F").

## Check

```bash
go test ./exercises/08_testing/01_table/ -v
```
