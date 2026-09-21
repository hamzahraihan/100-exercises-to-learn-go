# Wave 1 Expansion Implementation Plan (24 → 50 exercises)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add 26 exercises to existing sections (50 total), extending the `solutions` branch the same way.

**Architecture:** Same course architecture as v1: single module `learngo`, each exercise dir holds a broken-by-design stub (compiles under `go vet`, test fails on assertion), a stdlib `testing` test, and a teaching `README.md` that never prints full solutions.

**Tech Stack:** Go >= 1.23, stdlib only (`testing`, plus `strings`, `unicode/utf8`, `sort`/`slices`, `errors`, `cmp`, `sync`, `time` as lesson subjects), GitHub Actions (unchanged CI).

## Global Constraints

- Go floor: `go 1.23` in `go.mod`; module path exactly `learngo`.
- Test deps: stdlib `testing` only; no external dependencies.
- Exercise contract: stub `.go` file MUST compile under `go vet ./...`; its test MUST fail on assertion, never on build (test must reference only identifiers that exist in the stub).
- READMEs teach without printing full solutions: signatures + `// ...` placeholders only; naming a stdlib package (e.g. `unicode/utf8`) is allowed, showing the answer body is not.
- Concurrency/timing tests: verify under `-race`; assert outcomes, never elapsed durations; generous margins (timeout thresholds >= 10x the fast path).
- `go vet ./...` + `go build ./...` MUST pass on `main`; `go test ./...` MUST fail on `main` and pass on `solutions`.
- All prose is original Go teaching with attribution to the Rust source already in root README; do not copy Rust chapter text.
- Each exercise dir has exactly: `<name>.go`, `<name>_test.go`, `README.md`.

---

### Task 1: Expand `01_intro` (+2 → 5)

**Files:**
- Create: `exercises/01_intro/03_packages/greet.go`, `greet_test.go`, `README.md`
- Create: `exercises/01_intro/04_tooling/tooling.go`, `tooling_test.go`, `README.md`

**Interfaces:**
- Consumes: module `learngo` (existing).
- Produces: `package greet` with `func Shout(s string) string`; `package tooling` with `func Quote(s string) string`.

- [ ] **Step 1: `03_packages` test + stub + README**

`greet_test.go`:
```go
package greet

import "testing"

func TestShout(t *testing.T) {
	if got, want := Shout("hello"), "HELLO"; got != want {
		t.Fatalf("Shout(%q) = %q, want %q", "hello", got, want)
	}
}
```

`greet.go` (broken: lowercases instead of uppercases; import is used so it compiles):
```go
// Package greet teaches imports and exported names.
package greet

import "strings"

// Shout returns s in upper case.
// TODO: use strings.ToUpper instead of strings.ToLower.
func Shout(s string) string {
	return strings.ToLower(s)
}
```

`README.md`:
```markdown
# Packages and imports

Go code lives in packages. `import "strings"` makes the standard library's
string helpers available, and only capitalized names (`ToUpper`) are visible
outside their package.

## Task

Fix `Shout` in `greet.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/03_packages/ -v
```
```

- [ ] **Step 2: `04_tooling` test + stub + README**

`tooling_test.go`:
```go
package tooling

import "testing"

func TestQuote(t *testing.T) {
	if got, want := Quote("hi"), `"hi"`; got != want {
		t.Fatalf("Quote(%q) = %q, want %q", "hi", got, want)
	}
}
```

`tooling.go` (broken behavior + deliberately flat indentation for gofmt to flag):
```go
// Package tooling teaches gofmt and go vet.
package tooling

// Quote wraps s in double quotes.
// TODO: return "\"" + s + "\"" and run gofmt -w on this file.
func Quote(s string) string {
return s
}
```

`README.md`:
```markdown
# gofmt and go vet

`gofmt -l .` lists files whose formatting differs from Go standard — fix
them with `gofmt -w`. `go vet ./...` catches real bugs, e.g. a `Printf`
call with mismatched verbs:

```go
// vet flags this: %d with a string argument
// fmt.Printf("%d", "not a number")
```

That snippet is only an example. Your exercise has no vet issue — vet must
stay clean while the test fails.

## Task

Fix `Quote` in `tooling.go`, then run `gofmt -w tooling.go`,
`go vet ./exercises/01_intro/04_tooling/` and the check below.

## Check

```bash
go test ./exercises/01_intro/04_tooling/ -v
```
```

- [ ] **Step 3: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS (exit 0, no warnings).
Run: `go test ./exercises/01_intro/...`
Expected: FAIL (only the 2 new packages fail; the 3 v1 packages still fail too — all by design).

```bash
git add exercises/01_intro
git commit -m "feat: expand 01_intro with packages and tooling (wave 1)"
```

---

### Task 2: Expand `02_calculator` (+4 → 15)

**Files:** In each new dir (`12_floats`, `13_runes`, `14_strings`, `15_bitwise`): `calc.go`, `calc_test.go`, `README.md`. Every dir: `package calc`.

**Interfaces:**
- Consumes: nothing new.
- Produces: `func Half(n int) float64`, `func CountRunes(s string) int`, `func ToUpperFirst(s string) string`, `type Perm uint8` + consts `Read/Write/Execute` + `func Has(p, flag Perm) bool`.

- [ ] **Step 1: `12_floats`**

`calc_test.go`:
```go
package calc

import "testing"

func TestHalf(t *testing.T) {
	if got, want := Half(5), 2.5; got != want {
		t.Fatalf("Half(5) = %v, want %v", got, want)
	}
	if got, want := Half(4), 2.0; got != want {
		t.Fatalf("Half(4) = %v, want %v", got, want)
	}
}
```

`calc.go`:
```go
package calc

// Half returns n divided by two. Integer division truncates before the
// conversion, so convert first.
// TODO: return float64(n) / 2.
func Half(n int) float64 {
	return float64(n / 2)
}
```

`README.md`:
```markdown
# Floats

`float64` holds fractional numbers. Converting after integer division keeps
the truncation: `float64(5/2)` is `2.0`, while `float64(5)/2` is `2.5`.
The `math` package (`math.Pi`, `math.Sqrt`) covers the rest.

## Task

Fix `Half` in `calc.go` so odd inputs keep their `.5`.

## Check

```bash
go test ./exercises/02_calculator/12_floats/ -v
```
```

- [ ] **Step 2: `13_runes`**

`calc_test.go`:
```go
package calc

import "testing"

func TestCountRunes(t *testing.T) {
	if got, want := CountRunes("go"), 2; got != want {
		t.Fatalf("CountRunes(%q) = %d, want %d", "go", got, want)
	}
	if got, want := CountRunes("héllo"), 5; got != want {
		t.Fatalf("CountRunes(%q) = %d, want %d", "héllo", got, want)
	}
}
```

`calc.go`:
```go
package calc

// CountRunes returns the number of characters (runes) in s.
// len(s) counts bytes instead: "é" is 2 bytes but 1 rune.
// TODO: count runes (hint: the unicode/utf8 package helps).
func CountRunes(s string) int {
	return len(s)
}
```

`README.md`:
```markdown
# Runes

A Go `string` is bytes; a `rune` is one character. `"é"` is 2 bytes
(`len` reports 2) but a single rune. Indexing a string yields bytes —
ranging over it yields runes.

## Task

Fix `CountRunes` in `calc.go` so multibyte characters count once.

## Check

```bash
go test ./exercises/02_calculator/13_runes/ -v
```
```

- [ ] **Step 3: `14_strings`**

`calc_test.go`:
```go
package calc

import "testing"

func TestToUpperFirst(t *testing.T) {
	if got, want := ToUpperFirst("hello"), "Hello"; got != want {
		t.Fatalf("ToUpperFirst(%q) = %q, want %q", "hello", got, want)
	}
	if got := ToUpperFirst(""); got != "" {
		t.Fatalf("ToUpperFirst(%q) = %q, want %q", "", got, want)
	}
}
```

`calc.go`:
```go
package calc

// ToUpperFirst capitalizes the first character of s.
// Strings are immutable: build a new one from pieces.
// TODO: uppercase s[:1] and append the rest (watch the empty string).
func ToUpperFirst(s string) string {
	return s
}
```

`README.md`:
```markdown
# Strings

Strings support `==` and `+`, but you cannot assign to an index:
`s[0] = 'H'` does not compile. Build a new string with slicing and
concatenation instead. The `strings` package has the rest.

## Task

Fix `ToUpperFirst` in `calc.go`. Keep `""` mapping to `""`.

## Check

```bash
go test ./exercises/02_calculator/14_strings/ -v
```
```

- [ ] **Step 4: `15_bitwise`**

`calc_test.go`:
```go
package calc

import "testing"

func TestHas(t *testing.T) {
	if !Has(Read|Write, Write) {
		t.Fatal("Has(Read|Write, Write) = false, want true")
	}
	if Has(Read, Write) {
		t.Fatal("Has(Read, Write) = true, want false")
	}
	if !Has(Execute, Execute) {
		t.Fatal("Has(Execute, Execute) = false, want true")
	}
}
```

`calc.go`:
```go
package calc

// Perm is a set of permission flags.
type Perm uint8

const (
	Read Perm = 1 << iota
	Write
	Execute
)

// Has reports whether flag is set in p.
// TODO: return p&flag == flag.
func Has(p, flag Perm) bool {
	return false
}
```

`README.md`:
```markdown
# Bitwise operators

`<<` shifts bits, `&`/`|`/`^` combine them. With `iota`, each constant in
a block gets a successive value, so `1 << iota` builds power-of-two flags
that combine with `|` and test with `&`.

## Task

Fix `Has` in `calc.go` using `&`.

## Check

```bash
go test ./exercises/02_calculator/15_bitwise/ -v
```
```

- [ ] **Step 5: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/02_calculator/...`
Expected: FAIL (at least the 4 new packages fail on assertions).

```bash
git add exercises/02_calculator
git commit -m "feat: expand 02_calculator with floats, runes, strings, bits (wave 1)"
```

---

### Task 3: Expand `03_ticket_v1` (+6 → 8)

**Files:** In each new dir (`03_constructor`, `04_setters`, `05_receivers`, `06_zero_value`, `07_string`, `08_embed`): `ticket.go`, `ticket_test.go`, `README.md`. Every dir: `package ticket`, self-contained `Ticket` type (dirs never import each other).

**Interfaces:**
- Consumes: nothing new.
- Produces: `func NewTicket(title, description string) (Ticket, error)` (constructor dir), `func (t *Ticket) SetTitle(title string)`, `func (t Ticket) Renamed(newTitle string) Ticket`, `func (t Ticket) IsZero() bool`, `func (t Ticket) String() string`, `type Meta struct{ ID int }` + embedded `Ticket` + `func NewTicketWithID(id int, title string) Ticket`.

- [ ] **Step 1: `03_constructor`**

`ticket_test.go`:
```go
package ticket

import "testing"

func TestNewTicketValid(t *testing.T) {
	tk, err := NewTicket("Fix bug", "Crash on login screen")
	if err != nil {
		t.Fatalf("valid ticket rejected: %v", err)
	}
	if tk.Title != "Fix bug" {
		t.Fatalf("Title = %q", tk.Title)
	}
}

func TestNewTicketInvalid(t *testing.T) {
	if _, err := NewTicket("", "long enough description"); err == nil {
		t.Fatal("empty title accepted")
	}
	if _, err := NewTicket("ok title", "short"); err == nil {
		t.Fatal("short description accepted")
	}
}
```

`ticket.go`:
```go
// Package ticket teaches validated constructors.
package ticket

import "errors"

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// NewTicket builds a validated Ticket: non-empty title, description of at
// least 10 characters.
// TODO: return Ticket{...}, nil when valid, else a descriptive error.
func NewTicket(title, description string) (Ticket, error) {
	return Ticket{}, errors.New("TODO")
}
```

`README.md`: constructors (`New...` returning values + error), why validation lives at construction.

- [ ] **Step 2: `04_setters`**

`ticket_test.go`:
```go
package ticket

import "testing"

func TestSetTitle(t *testing.T) {
	tk := Ticket{Title: "old", Description: "long enough description"}
	tk.SetTitle("new")
	if tk.Title != "new" {
		t.Fatalf("Title = %q, want %q", tk.Title, "new")
	}
}
```

`ticket.go`:
```go
// Package ticket teaches pointer-receiver setters.
package ticket

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// SetTitle changes the title. A pointer receiver mutates the caller's copy;
// a value receiver would mutate a throwaway copy.
// TODO: assign t.Title = title.
func (t *Ticket) SetTitle(title string) {
}
```

`README.md`: pointer vs value receivers for mutation.

- [ ] **Step 3: `05_receivers`**

`ticket_test.go`:
```go
package ticket

import "testing"

func TestRenamed(t *testing.T) {
	orig := Ticket{Title: "a", Description: "long enough description"}
	got := orig.Renamed("b")
	if got.Title != "b" || got.Description != orig.Description {
		t.Fatalf("Renamed = %+v, want title b with description kept", got)
	}
	if orig.Title != "a" {
		t.Fatalf("original mutated: %+v", orig)
	}
}
```

`ticket.go`:
```go
// Package ticket teaches value receivers that return modified copies.
package ticket

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// Renamed returns a copy of t with a new title, leaving t untouched.
// TODO: set t.Title and return t.
func (t Ticket) Renamed(newTitle string) Ticket {
	return Ticket{}
}
```

`README.md`: value semantics, when to return copies.

- [ ] **Step 4: `06_zero_value`**

`ticket_test.go`:
```go
package ticket

import "testing"

func TestIsZero(t *testing.T) {
	if !(Ticket{}).IsZero() {
		t.Fatal("Ticket{}.IsZero() = false, want true")
	}
	if (Ticket{Title: "x"}).IsZero() {
		t.Fatal("non-empty ticket reported zero")
	}
}
```

`ticket.go`:
```go
// Package ticket teaches designing for useful zero values.
package ticket

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// IsZero reports whether t is the zero value.
// TODO: return t.Title == "" && t.Description == "".
func (t Ticket) IsZero() bool {
	return false
}
```

`README.md`: Go zero values, designing structs usable without constructors.

- [ ] **Step 5: `07_string`**

`ticket_test.go`:
```go
package ticket

import (
	"fmt"
	"strings"
	"testing"
)

func TestTicketString(t *testing.T) {
	tk := Ticket{Title: "bug", Description: "long enough description"}
	if got := fmt.Sprint(tk); !strings.Contains(got, "bug") {
		t.Fatalf("String() = %q, want title inside", got)
	}
}
```

`ticket.go`:
```go
// Package ticket teaches the String method (fmt.Stringer, implicit).
package ticket

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// String describes the ticket for humans and fmt printing.
// TODO: include the title (hint: fmt.Sprintf is allowed here).
func (t Ticket) String() string {
	return ""
}
```

`README.md`: `fmt.Stringer`, implicit satisfaction (no `implements` keyword).

- [ ] **Step 6: `08_embed`**

`ticket_test.go`:
```go
package ticket

import "testing"

func TestNewTicketWithID(t *testing.T) {
	tk := NewTicketWithID(7, "bug")
	if tk.ID != 7 || tk.Title != "bug" {
		t.Fatalf("got %+v, want ID 7 title bug", tk)
	}
}
```

`ticket.go`:
```go
// Package ticket teaches struct embedding.
package ticket

// Meta holds fields shared by several domain types.
type Meta struct {
	ID int
}

// Ticket embeds Meta: ID is promoted, so tk.ID works directly.
// TODO: return Ticket{Meta: Meta{ID: id}, Title: title}.
type Ticket struct {
	Meta
	Title string
}

// NewTicketWithID builds a Ticket with an id.
func NewTicketWithID(id int, title string) Ticket {
	return Ticket{}
}
```

`README.md`: embedding vs inheritance, promoted fields.

- [ ] **Step 7: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/03_ticket_v1/...`
Expected: FAIL (6 new packages fail on assertions; `var _` style checks absent here — plain tests).

```bash
git add exercises/03_ticket_v1
git commit -m "feat: expand 03_ticket_v1 with constructors, setters, embedding (wave 1)"
```

---

### Task 4: Expand `04_interfaces` (+5 → 7)

**Files:** New dirs `03_any` (`any.go`, `any_test.go`), `04_type_switch` (`describe.go`, `describe_test.go`), `05_embed_iface` (`bucket.go`, `bucket_test.go`), `06_comparable` (`keys.go`, `keys_test.go`), `07_min_generic` (`min.go`, `min_test.go`), each + `README.md`. Package names: `show`, `kind`, `bucket`, `keys`, `min`.

**Interfaces:**
- Consumes: nothing new.
- Produces: `func SprintAny(v any) string`, `func Describe(v any) string`, `type ReadWriter interface` + `Bucket` implementing it, `func Keys[K comparable, V any](m map[K]V) []K`, `func Min[T cmp.Ordered](a, b T) T`.

- [ ] **Step 1: `03_any`**

`any_test.go`:
```go
package show

import "testing"

func TestSprintAny(t *testing.T) {
	if got, want := SprintAny(42), "42"; got != want {
		t.Fatalf("SprintAny(42) = %q, want %q", got, want)
	}
	if got, want := SprintAny("hi"), "hi"; got != want {
		t.Fatalf("SprintAny(%q) = %q, want %q", "hi", got, want)
	}
}
```

Stub (`any.go`):
```go
// Package show teaches the any alias for interface{}.
package show

// SprintAny formats any value.
// TODO: return fmt.Sprint(v) (import fmt).
func SprintAny(v any) string {
	return ""
}
```
README: `any`, formatting verbs, when `any` is (not) appropriate.

- [ ] **Step 2: `04_type_switch`**

Test:
```go
package kind

import "testing"

func TestDescribe(t *testing.T) {
	tests := []struct {
		in   any
		want string
	}{
		{42, "int: 42"},
		{"hi", "string: hi"},
		{1.5, "unknown"},
		{nil, "unknown"},
	}
	for _, tt := range tests {
		if got := Describe(tt.in); got != tt.want {
			t.Errorf("Describe(%v) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
```
Stub:
```go
// Package kind teaches type switches.
package kind

// Describe names supported dynamic types.
// TODO: switch v := v.(type) for int and string cases.
func Describe(v any) string {
	return "unknown"
}
```
README: type switches, comma-ok assertion (mentioned, not solved).

- [ ] **Step 3: `05_embed_iface`**

Test:
```go
package bucket

import "testing"

func TestReadWriter(t *testing.T) {
	var b Bucket
	b.Write("hi")
	if got := b.Read(); got != "hi" {
		t.Fatalf("Read() = %q, want %q", got, "hi")
	}
}
```
Stub:
```go
// Package bucket teaches interface embedding.
package bucket

// Reader reads stored data.
type Reader interface {
	Read() string
}

// Writer stores data.
type Writer interface {
	Write(s string)
}

// ReadWriter embeds both: one interface, two capabilities.
type ReadWriter interface {
	Reader
	Writer
}

// Bucket is an in-memory store.
type Bucket struct {
	data string
}

// Write stores s.
// TODO: assign b.data = s.
func (b *Bucket) Write(s string) {
}

// Read returns stored data.
// TODO: return b.data.
func (b *Bucket) Read() string {
	return ""
}

var _ ReadWriter = (*Bucket)(nil)
```
README: embedding interfaces (like `io.ReadWriter`), compile-time `var _` checks.

- [ ] **Step 4: `06_comparable`**

Test:
```go
package keys

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	if got, want := Keys(map[string]int{"a": 1}), []string{"a"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Keys = %v, want %v", got, want)
	}
	if got := Keys(map[int]string{}); len(got) != 0 {
		t.Fatalf("Keys of empty map = %v, want empty", got)
	}
}
```
Stub:
```go
// Package keys teaches the comparable constraint.
package keys

// Keys returns all map keys. K must be comparable (usable as a map key).
// TODO: range over m and append each k.
func Keys[K comparable, V any](m map[K]V) []K {
	return nil
}
```
README: constraints, why `comparable` exists.

- [ ] **Step 5: `07_min_generic`**

`min_test.go`:
```go
package min

import "testing"

func TestMin(t *testing.T) {
	if got, want := Min(3, 7), 3; got != want {
		t.Fatalf("Min(3, 7) = %v, want %v", got, want)
	}
	if got, want := Min(7, 3), 3; got != want {
		t.Fatalf("Min(7, 3) = %v, want %v", got, want)
	}
	if got, want := Min("b", "a"), "a"; got != want {
		t.Fatalf("Min(b, a) = %q, want %q", got, want)
	}
}
```

Stub (`min.go`):
```go
// Package min teaches ordered constraints.
package min

import "cmp"

// Min returns the smaller of a and b for any ordered type.
// TODO: compare and return the smaller (the constraint allows <).
func Min[T cmp.Ordered](a, b T) T {
	return a
}
```
README: `cmp.Ordered`, generics vs overloads (describe the comparison, elide the if).

- [ ] **Step 6: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/04_interfaces/...`
Expected: FAIL (5 new packages on assertions).

```bash
git add exercises/04_interfaces
git commit -m "feat: expand 04_interfaces with any, switches, embedding, generics (wave 1)"
```

---

### Task 5: Expand `05_ticket_v2` (+3 → 5)

**Files:** New dirs `03_wrap`, `04_join`, `05_panic_vs_error`, each `ticket.go`/`must.go` + `*_test.go` + `README.md` as below. `03_wrap`/`04_join`: `package ticket`. `05_panic_vs_error`: `package must` (`must.go`, `must_test.go`).

**Interfaces:**
- Consumes: nothing new.
- Produces: `var ErrNotFound` + `func FindTicket(ids []int, id int) (int, error)`; `var ErrBadTitle/ErrBadDesc` + `func ValidateAll(title, desc string) error`; `func MustParse(s string) (n int)`.

- [ ] **Step 1: `03_wrap`**

Test:
```go
package ticket

import (
	"errors"
	"testing"
)

func TestFindTicket(t *testing.T) {
	idx, err := FindTicket([]int{10, 20, 30}, 20)
	if err != nil || idx != 1 {
		t.Fatalf("FindTicket = (%d, %v), want (1, nil)", idx, err)
	}
	_, err = FindTicket([]int{10, 20}, 99)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want errors.Is(err, ErrNotFound)", err)
	}
}
```
Stub:
```go
// Package ticket teaches wrapping errors with %w.
package ticket

import "errors"

// ErrNotFound is returned when an id is absent.
var ErrNotFound = errors.New("not found")

// FindTicket returns the index of id, or a wrapped ErrNotFound.
// TODO: loop; on miss return -1 and fmt.Errorf("ticket %d: %w", id, ErrNotFound).
func FindTicket(ids []int, id int) (int, error) {
	return -1, errors.New("TODO")
}
```
README: sentinel errors, `%w`, `errors.Is` (describe, don't solve).

- [ ] **Step 2: `04_join`**

`ticket_test.go`:
```go
package ticket

import (
	"errors"
	"testing"
)

func TestValidateAll(t *testing.T) {
	if err := ValidateAll("Ticket title", "long enough description"); err != nil {
		t.Fatalf("valid input rejected: %v", err)
	}
	err := ValidateAll("", "short")
	if !errors.Is(err, ErrBadTitle) || !errors.Is(err, ErrBadDesc) {
		t.Fatalf("err = %v, want both ErrBadTitle and ErrBadDesc", err)
	}
	err = ValidateAll("", "long enough description")
	if !errors.Is(err, ErrBadTitle) || errors.Is(err, ErrBadDesc) {
		t.Fatalf("err = %v, want only ErrBadTitle", err)
	}
}
```

Stub (`ticket.go`):
```go
// Package ticket teaches joining multiple errors.
package ticket

import "errors"

var (
	// ErrBadTitle is returned for empty titles.
	ErrBadTitle = errors.New("bad title")
	// ErrBadDesc is returned for short descriptions.
	ErrBadDesc = errors.New("bad description")
)

// ValidateAll checks title and description, reporting every problem.
// TODO: collect per-field errors and combine with errors.Join (nil when valid).
func ValidateAll(title, desc string) error {
	return errors.New("TODO")
}
```
README: `errors.Join`, multi-error validation (describe, don't solve).

- [ ] **Step 3: `05_panic_vs_error`**

Test:
```go
package must

import "testing"

func TestMustParse(t *testing.T) {
	if got := MustParse("hi"); got != 2 {
		t.Fatalf("MustParse(%q) = %d, want 2", "hi", got)
	}
	if got := MustParse(""); got != -1 {
		t.Fatalf("MustParse(%q) = %d, want -1 (no panic must escape)", "", got)
	}
}
```
Stub:
```go
// Package must teaches converting panics to errors with recover.
package must

func parseOrPanic(s string) int {
	if s == "" {
		panic("empty input")
	}
	return len(s)
}

// MustParse returns len(s), or -1 when s is empty. No panic escapes.
// TODO: defer a func that recovers and sets the named return n to -1.
func MustParse(s string) (n int) {
	return parseOrPanic(s)
}
```
README: panic vs error convention, `Must` helpers, recover mechanics (describe the defer pattern's shape, elide exact code).

- [ ] **Step 4: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/05_ticket_v2/...`
Expected: FAIL (3 new packages; `05_panic_vs_error` fails via escaping panic — still a test failure, asserted in the report).

```bash
git add exercises/05_ticket_v2
git commit -m "feat: expand 05_ticket_v2 with wrap, join, recover (wave 1)"
```

---

### Task 6: Expand `06_ticket_store` (+4 → 6)

**Files:** New dirs `03_slices_fn`, `04_sort_filter`, `05_pagination` (`store.go`, `store_test.go`, self-contained `type Ticket struct{ ID int; Title string; Closed bool }`, `package store`), `06_map_idioms` (`store.go`, `store_test.go`, `type Ticket struct{ Title string }` + `type Registry struct{ byTitle map[string]Ticket }`, `package store`). Each + `README.md`.

**Interfaces:**
- Consumes: nothing new.
- Produces: `func CloneTickets(ts []Ticket) []Ticket`, `func SortedByID(ts []Ticket) []Ticket` + `func OpenOnly(ts []Ticket) []Ticket`, `func Page(ts []Ticket, offset, limit int) []Ticket`, `func (r *Registry) Put(t Ticket)` + `func (r *Registry) Lookup(title string) (Ticket, bool)`.

- [ ] **Step 1: `03_slices_fn`**

`store_test.go`:
```go
package store

import (
	"reflect"
	"testing"
)

func TestCloneTickets(t *testing.T) {
	in := []Ticket{{ID: 1, Title: "a"}, {ID: 2, Title: "b"}}
	out := CloneTickets(in)
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("CloneTickets = %v, want %v", out, in)
	}
	out[0].Title = "changed"
	if in[0].Title != "a" {
		t.Fatalf("clone aliases original: %v", in)
	}
}
```

Stub (`store.go`):
```go
// Package store teaches slice cloning.
package store

// Ticket is a stored item.
type Ticket struct {
	ID     int
	Title  string
	Closed bool
}

// CloneTickets returns an independent copy: mutating the clone must not
// affect the original.
// TODO: append into a fresh slice (or use the slices package Clone).
func CloneTickets(ts []Ticket) []Ticket {
	return nil
}
```
README: len/cap, copy vs aliasing, `slices.Clone` mentioned by name only.

- [ ] **Step 2: `04_sort_filter`**

`store_test.go`:
```go
package store

import (
	"reflect"
	"testing"
)

func TestSortedByID(t *testing.T) {
	in := []Ticket{{ID: 3}, {ID: 1}, {ID: 2}}
	want := []Ticket{{ID: 1}, {ID: 2}, {ID: 3}}
	if got := SortedByID(in); !reflect.DeepEqual(got, want) {
		t.Fatalf("SortedByID = %v, want %v", got, want)
	}
}

func TestOpenOnly(t *testing.T) {
	in := []Ticket{{ID: 1, Closed: true}, {ID: 2}, {ID: 3, Closed: true}}
	want := []Ticket{{ID: 2}}
	if got := OpenOnly(in); !reflect.DeepEqual(got, want) {
		t.Fatalf("OpenOnly = %v, want %v", got, want)
	}
}
```

Stubs (`store.go`, same `Ticket` shape as `03_slices_fn`):
```go
// SortedByID returns tickets ordered by ID ascending.
// TODO: sort a copy (hint: slices.SortFunc or sort.Slice).
func SortedByID(ts []Ticket) []Ticket {
	return nil
}

// OpenOnly keeps tickets that are not closed.
// TODO: append matches into a fresh slice.
func OpenOnly(ts []Ticket) []Ticket {
	return nil
}
```
README: `slices.SortFunc`/`sort.Slice` mentioned by name, filter-by-append pattern described without code.

- [ ] **Step 3: `05_pagination`**

`store_test.go` (all assertions via `len()`, so nil counts as empty):
```go
package store

import "testing"

func fiveTickets() []Ticket {
	return []Ticket{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}}
}

func TestPage(t *testing.T) {
	ts := fiveTickets()
	got := Page(ts, 1, 2)
	if len(got) != 2 || got[0].ID != 2 || got[1].ID != 3 {
		t.Fatalf("Page(offset 1, limit 2) = %v, want IDs [2 3]", got)
	}
	if got := Page(ts, 99, 10); len(got) != 0 {
		t.Fatalf("Page past end = %v, want empty", got)
	}
	if got := Page(ts, 0, 0); len(got) != 0 {
		t.Fatalf("Page limit 0 = %v, want empty", got)
	}
	if got := Page(ts, -1, -5); len(got) != 0 {
		t.Fatalf("Page negative = %v, want empty (clamped, no panic)", got)
	}
}
```

Stub (`store.go`, same `Ticket` shape as `03_slices_fn`):
```go
// Page returns up to limit tickets starting at offset. Out-of-range or
// non-positive inputs yield an empty (len 0) result, never a panic.
// TODO: clamp offset/limit into range, then slice.
func Page(ts []Ticket, offset, limit int) []Ticket {
	return nil
}
```
README: offset/limit, clamping, why empty means len 0 not a specific nil.

- [ ] **Step 4: `06_map_idioms`**

`store_test.go`:
```go
package store

import "testing"

func TestRegistry(t *testing.T) {
	var r Registry
	r.Put(Ticket{Title: "a"})
	tk, ok := r.Lookup("a")
	if !ok || tk.Title != "a" {
		t.Fatalf("Lookup(%q) = (%+v, %v), want found", "a", tk, ok)
	}
	if _, ok := r.Lookup("missing"); ok {
		t.Fatal("Lookup(missing) found, want false")
	}
}
```

Stub (`store.go`):
```go
// Package store teaches map idioms.
package store

// Ticket is a stored item.
type Ticket struct {
	Title string
}

// Registry indexes tickets by title.
type Registry struct {
	byTitle map[string]Ticket
}

// Put stores t. The zero Registry has a nil map: make it on first use,
// since assigning into a nil map panics.
// TODO: lazily make r.byTitle, then store under t.Title.
func (r *Registry) Put(t Ticket) {
}

// Lookup finds a ticket by title with the comma-ok idiom.
// TODO: return r.byTitle[title].
func (r *Registry) Lookup(title string) (Ticket, bool) {
	return Ticket{}, false
}
```
README: comma-ok, delete, lazy `make` for nil maps (describe, don't solve).

- [ ] **Step 5: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/06_ticket_store/...`
Expected: FAIL (4 new packages on assertions).

```bash
git add exercises/06_ticket_store
git commit -m "feat: expand 06_ticket_store with slices, sort, pages, maps (wave 1)"
```

---

### Task 7: Expand `07_concurrency` (+2 → 4)

**Files:**
- Create: `exercises/07_concurrency/03_buffered/collect.go`, `collect_test.go`, `README.md` (`package collect`)
- Create: `exercises/07_concurrency/04_select/fetch.go`, `fetch_test.go`, `README.md` (`package fetch`)

**Interfaces:**
- Consumes: nothing new.
- Produces: `func Collect(n int) []int` (order-preserving double via goroutines + buffered channel), `func Fetch(d time.Duration) (string, error)` (succeeds fast, times out after 200ms).

- [ ] **Step 1: `03_buffered`**

Test:
```go
package collect

import (
	"reflect"
	"testing"
)

func TestCollect(t *testing.T) {
	if got, want := Collect(5), []int{0, 2, 4, 6, 8}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Collect(5) = %v, want %v", got, want)
	}
	if got, want := Collect(1), []int{0}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Collect(1) = %v, want %v", got, want)
	}
}
```
Stub:
```go
// Package collect teaches buffered channels for fan-out collection.
package collect

// Collect returns [0, 2, ..., 2*(n-1)] computed by n goroutines sending
// over a buffered channel, collected in order.
// TODO: make(chan int, n), one goroutine per i sending 2*i, collect all.
func Collect(n int) []int {
	return make([]int, n)
}
```
README: buffered vs unbuffered, fan-out/collect pattern (describe shape, elide code).

- [ ] **Step 2: `04_select`**

Test:
```go
package fetch

import (
	"testing"
	"time"
)

func TestFetchFast(t *testing.T) {
	got, err := Fetch(10 * time.Millisecond)
	if err != nil || got != "data" {
		t.Fatalf("Fetch(fast) = (%q, %v), want (data, nil)", got, err)
	}
}

func TestFetchTimeout(t *testing.T) {
	got, err := Fetch(500 * time.Millisecond)
	if err == nil {
		t.Fatalf("Fetch(slow) = (%q, nil), want timeout error", got)
	}
}
```
Stub (slow path takes 500ms then wrongly succeeds — outcome assert fails, margins 25x):
```go
// Package fetch teaches select with timeouts.
package fetch

import "time"

// Fetch simulates a fetch taking d. It must fail when d exceeds 200ms.
// TODO: select on time.After(d) for success vs time.After(200ms) for timeout.
func Fetch(d time.Duration) (string, error) {
	<-time.After(d)
	return "data", nil
}
```
README: `select`, `time.After`, generous margins (assert outcome, never duration).

- [ ] **Step 3: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test -race ./exercises/07_concurrency/03_buffered/ ./exercises/07_concurrency/04_select/`
Expected: FAIL (both packages; capture output; note wall time ~0.6s from the 500ms timeout case).

```bash
git add exercises/07_concurrency
git commit -m "feat: expand 07_concurrency with buffered channels and select (wave 1)"
```

---

### Task 8: Extend `solutions` branch and release-check Wave 1

**Files:** none new (branch operation + stub fixes only; never edit tests or READMEs).

**Interfaces:**
- Consumes: all Wave 1 stubs + tests from Tasks 1–7.
- Produces: `solutions` branch where `go test ./...` passes 50/50.

- [ ] **Step 1: Apply fixes on `solutions`**

```bash
git checkout solutions
```
Apply the minimal inverse-of-TODO fix per new exercise (e.g. `Half`: `return float64(n) / 2`; `CountRunes`: `return utf8.RuneCountInString(s)`; `Has`: `return p&flag == flag`; `Min`: `if b < a { return b }; return a`; `Fetch`: select skeleton completed). One commit per touched section:
```bash
git add exercises/01_intro && git commit -m "solution: 01_intro wave 1"
git add exercises/02_calculator && git commit -m "solution: 02_calculator wave 1"
git add exercises/03_ticket_v1 && git commit -m "solution: 03_ticket_v1 wave 1"
git add exercises/04_interfaces && git commit -m "solution: 04_interfaces wave 1"
git add exercises/05_ticket_v2 && git commit -m "solution: 05_ticket_v2 wave 1"
git add exercises/06_ticket_store && git commit -m "solution: 06_ticket_store wave 1"
git add exercises/07_concurrency && git commit -m "solution: 07_concurrency wave 1"
```

- [ ] **Step 2: Verify `solutions` is green**

Run (on `solutions`): `go vet ./... && go build ./... && go test ./...`
Expected: PASS, 50/50 packages ok.
Run: `go test -race -count=1 ./exercises/07_concurrency/...`
Expected: PASS.

- [ ] **Step 3: Verify `main` is red-green-correct and finish on `main`**

```bash
git checkout main
```
Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./...`
Expected: FAIL (26 new + 24 old packages fail; proves unsolved-by-design).
Do NOT push (no remote push without explicit confirmation).
