# Wave 2 Implementation Plan (50 → 78 exercises)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add 28 exercises in 5 new sections (08_testing, 09_strings, 10_time, 11_json, 12_io), extending `solutions` the same way.

**Architecture:** Same course architecture: single module `learngo`, per-exercise triplet (broken-by-design stub compiling under `go vet` + failing test + teaching README). 11_json/12_io use Ticket types shaped for the Wave 3 server; other sections use generic examples. Sections never import each other.

**Tech Stack:** Go >= 1.23, stdlib only (`testing`, `strings`, `strconv`, `time`, `encoding/json`, `io`, `bufio`, `bytes`, `os`, `unicode/utf8` as lesson subjects), unchanged CI.

## Global Constraints

- Go floor: `go 1.23` in `go.mod`; module path exactly `learngo`.
- Test deps: stdlib `testing` only; no external dependencies.
- Exercise contract: stub `.go` MUST compile under `go vet ./...`; its test MUST fail on assertion, never on build (every identifier the test references must exist in the stub).
- READMEs teach without printing full solutions: signatures + `// ...` only; naming a stdlib package is allowed, showing the answer body is not.
- Determinism: fixed `time.Date(...)` fixtures, no sleeps, no wall-clock assertions; IO uses `t.TempDir()`; JSON asserts decode-and-compare, never key order.
- `*.go` LF enforced by `.gitattributes`; `go vet`+`go build` green on `main`, `go test` red on `main` by design, green on `solutions`.
- All prose original Go teaching; attribution to the Rust source already in root README.

---

### Task 1: New section `08_testing` (+7 → 57)

**Files:** New dirs `01_table` (grade.go, grade_test.go), `02_subtests` (clamp.go, clamp_test.go), `03_helpers` (mean.go, mean_test.go), `04_testdata` (poem.go, poem_test.go, testdata/poem.txt fixture), `05_examples` (double.go, double_test.go), `06_benchmarks` (fib.go, fib_test.go), `07_errassert` (divide.go, divide_test.go), each + `README.md`. Packages: `grade`, `clamp`, `mean`, `poem`, `double`, `fib`, `divide`.

**Interfaces:**
- Consumes: nothing.
- Produces: `func Grade(score int) string`, `func Clamp(n, lo, hi int) int`, `func Mean(xs []float64) float64`, `func FirstLine(path string) (string, error)`, `func Double(n int) int` + `ExampleDouble`, `func Fib(n int) int` + `BenchmarkFib`, `var ErrZeroDivisor` + `func Divide(a, b int) (int, error)`.

- [ ] **Step 1: `01_table`**

`grade_test.go`:
```go
package grade

import "testing"

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

`grade.go`:
```go
// Package grade teaches table-driven tests.
package grade

// Grade maps a score to A/B/C/F.
// TODO: if/else chain on 90/80/70 cutoffs.
func Grade(score int) string {
	return ""
}
```

`README.md`: table pattern (slice of in/want, loop, `t.Errorf` so all cases report); Task + Check.

- [ ] **Step 2: `02_subtests`**

`clamp_test.go`:
```go
package clamp

import "testing"

func TestClamp(t *testing.T) {
	t.Run("above", func(t *testing.T) {
		if got, want := Clamp(99, 0, 10), 10; got != want {
			t.Fatalf("Clamp(99, 0, 10) = %d, want %d", got, want)
		}
	})
	t.Run("below", func(t *testing.T) {
		if got, want := Clamp(-5, 0, 10), 0; got != want {
			t.Fatalf("Clamp(-5, 0, 10) = %d, want %d", got, want)
		}
	})
	t.Run("inside", func(t *testing.T) {
		if got, want := Clamp(4, 0, 10), 4; got != want {
			t.Fatalf("Clamp(4, 0, 10) = %d, want %d", got, want)
		}
	})
}
```

Stub (`clamp.go`):
```go
// Package clamp teaches subtests.
package clamp

// Clamp constrains n to [lo, hi].
// TODO: if n < lo return lo; if n > hi return hi; else return n.
func Clamp(n, lo, hi int) int {
	return n
}
```
README: subtests isolate failures (each `t.Run` reports separately).

- [ ] **Step 3: `03_helpers`**

`mean_test.go` (full):
```go
package mean

import "testing"

func mustMean(t *testing.T, xs []float64, want float64) {
	t.Helper()
	if got := Mean(xs); got != want {
		t.Fatalf("Mean(%v) = %v, want %v", xs, got, want)
	}
}

func TestMean(t *testing.T) {
	mustMean(t, []float64{1, 2, 3}, 2)
	mustMean(t, nil, 0)
}
```

Stub (`mean.go`):
```go
// Package mean teaches test helpers.
package mean

// Mean averages xs (0 for empty).
// TODO: sum and divide; guard len(xs) == 0.
func Mean(xs []float64) float64 {
	return 0
}
```
README: `t.Helper()` marks helpers so failures point at callers.

- [ ] **Step 4: `04_testdata` + fixture**

`testdata/poem.txt` (exact, 4 lines):
```
Roses are red
Violets are blue
Go tests are green
When fixtures help you
```
Test (`poem_test.go`):
```go
package poem

import "testing"

func TestFirstLine(t *testing.T) {
	got, err := FirstLine("testdata/poem.txt")
	if err != nil {
		t.Fatalf("FirstLine errored: %v", err)
	}
	if got != "Roses are red" {
		t.Fatalf("FirstLine = %q, want %q", got, "Roses are red")
	}
}
```

Stub (`poem.go`):
```go
// Package poem teaches testdata fixtures.
package poem

// FirstLine returns the first line of the file at path.
// TODO: os.ReadFile + split on newline (hint: strings.SplitN).
func FirstLine(path string) (string, error) {
	return "", nil
}
```
README: `testdata/` convention, relative paths run from the package dir.

- [ ] **Step 5: `05_examples`**

`double_test.go`:
```go
package double

import "fmt"

func ExampleDouble() {
	fmt.Println(Double(2))
	// Output: 4
}
```
Stub (`double.go`):
```go
// Package double teaches Example tests.
package double

// Double returns twice n.
// TODO: return 2 * n.
func Double(n int) int {
	return 0
}
```
README: Example naming, `// Output:` matching, docs + tests in one.

- [ ] **Step 6: `06_benchmarks`**

`fib_test.go`:
```go
package fib

import "testing"

func TestFib(t *testing.T) {
	if got, want := Fib(10), 55; got != want {
		t.Fatalf("Fib(10) = %d, want %d", got, want)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}

func TestBenchmarkRuns(t *testing.T) {
	if res := testing.Benchmark(BenchmarkFib); res.N <= 0 {
		t.Fatal("benchmark did not run")
	}
}
```
Stub `Fib` returns 0 (TestFib fails; benchmark runs). Stub (`fib.go`):
```go
// Package fib teaches benchmarks.
package fib

// Fib returns the n-th Fibonacci number (Fib(0) == 0, Fib(1) == 1).
// TODO: loop accumulating a, b (or recurse — slower, still correct).
func Fib(n int) int {
	return 0
}
```
README: `go test -bench . -run '^$'`, what benchmarks measure (describe, no code).

- [ ] **Step 7: `07_errassert`**

`divide_test.go`:
```go
package divide

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	if got, err := Divide(6, 3); err != nil || got != 2 {
		t.Fatalf("Divide(6, 3) = (%d, %v), want (2, nil)", got, err)
	}
	if _, err := Divide(1, 0); !errors.Is(err, ErrZeroDivisor) {
		t.Fatalf("Divide(1, 0) err = %v, want ErrZeroDivisor", err)
	}
}
```

Stub (`divide.go`):
```go
// Package divide teaches asserting errors.
package divide

import "errors"

// ErrZeroDivisor is returned for division by zero.
var ErrZeroDivisor = errors.New("division by zero")

// Divide returns a/b, or ErrZeroDivisor when b is 0.
// TODO: branch on b == 0.
func Divide(a, b int) (int, error) {
	return 0, nil
}
```
README: never compare errors with `==`; use `errors.Is`.

- [ ] **Step 8: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/08_testing/...`
Expected: FAIL (7 packages on assertions; `04_testdata` fails on content mismatch, never missing-file).

```bash
git add exercises/08_testing
git commit -m "feat: add 08_testing section (wave 2)"
```

---

### Task 2: New section `09_strings` (+6 → 63)

**Files:** New dirs `01_fields` (words.go, words_test.go), `02_trim` (trim.go, trim_test.go), `03_contains` (mail.go, mail_test.go), `04_replace` (redact.go, redact_test.go), `05_builder` (concat.go, concat_test.go), `06_strconv` (parse.go, parse_test.go), each + `README.md`. Packages: `words`, `trim`, `mail`, `redact`, `concat`, `parse`.

**Interfaces:**
- Consumes: nothing.
- Produces: `func Words(s string) []string` + `func JoinWords(w []string) string`, `func Clean(s string) string` + `func TrimExt(name string) string`, `func IsEmail(s string) bool`, `func Redact(s string) string`, `func ConcatN(s string, n int) string`, `func SumStrings(a, b string) (int, error)`.

- [ ] **Step 1: `01_fields`**

`words_test.go`:
```go
package words

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	want := []string{"go", "is", "fun"}
	if got := Words("  go  is fun "); !reflect.DeepEqual(got, want) {
		t.Fatalf("Words = %v, want %v", got, want)
	}
}

func TestJoinWords(t *testing.T) {
	if got, want := JoinWords([]string{"a", "b"}), "a b"; got != want {
		t.Fatalf("JoinWords = %q, want %q", got, want)
	}
}
```

Stubs (`words.go`):
```go
// Package words teaches splitting and joining.
package words

// Words splits s on whitespace runs.
// TODO: use strings.Fields (import strings).
func Words(s string) []string {
	return nil
}

// JoinWords joins with single spaces.
// TODO: use strings.Join(w, " ").
func JoinWords(w []string) string {
	return ""
}
```
README: `strings.Fields` vs `Split`, joining with separator.

- [ ] **Step 2: `02_trim`**

`trim_test.go`:
```go
package trim

import "testing"

func TestClean(t *testing.T) {
	if got, want := Clean("  hi\n"), "hi"; got != want {
		t.Fatalf("Clean = %q, want %q", got, want)
	}
}

func TestTrimExt(t *testing.T) {
	if got, want := TrimExt("notes.txt"), "notes"; got != want {
		t.Fatalf("TrimExt = %q, want %q", got, want)
	}
}
```

Stubs (`trim.go`):
```go
// Package trim teaches trimming.
package trim

// Clean strips surrounding whitespace.
// TODO: use strings.TrimSpace (import strings).
func Clean(s string) string {
	return s
}

// TrimExt strips a trailing ".txt".
// TODO: use strings.TrimSuffix.
func TrimExt(name string) string {
	return name
}
```
README: `TrimSpace`, `TrimSuffix`, cutsets (describe).

- [ ] **Step 3: `03_contains`**

`mail_test.go`:
```go
package mail

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"a@b.co", true},
		{"nope", false},
		{"@", true},
	}
	for _, tt := range tests {
		if got := IsEmail(tt.in); got != tt.want {
			t.Errorf("IsEmail(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
```

Stub (`mail.go`):
```go
// Package mail teaches substring checks.
package mail

// IsEmail is a deliberately naive check: an address contains "@".
// TODO: use strings.Contains (import strings).
func IsEmail(s string) bool {
	return false
}
```
README: `Contains`/`HasPrefix`/`HasSuffix` (name only).

- [ ] **Step 4: `04_replace`**

`redact_test.go`:
```go
package redact

import "testing"

func TestRedact(t *testing.T) {
	if got, want := Redact("my secret here"), "my [redacted] here"; got != want {
		t.Fatalf("Redact = %q, want %q", got, want)
	}
	if got, want := Redact("nothing to hide"), "nothing to hide"; got != want {
		t.Fatalf("Redact = %q, want %q", got, want)
	}
}
```

Stub (`redact.go`):
```go
// Package redact teaches replacement.
package redact

// Redact replaces every "secret" with "[redacted]".
// TODO: use strings.ReplaceAll (import strings).
func Redact(s string) string {
	return s
}
```
README: `ReplaceAll` vs `Replace` count arg.

- [ ] **Step 5: `05_builder`**

`concat_test.go`:
```go
package concat

import "testing"

func TestConcatN(t *testing.T) {
	if got, want := ConcatN("ab", 3), "ababab"; got != want {
		t.Fatalf("ConcatN = %q, want %q", got, want)
	}
	if got := ConcatN("x", 0); got != "" {
		t.Fatalf("ConcatN(x, 0) = %q, want empty", got)
	}
}
```

Stub (`concat.go`):
```go
// Package concat teaches strings.Builder.
package concat

// ConcatN repeats s n times efficiently.
// TODO: loop n times writing into a strings.Builder (import strings).
func ConcatN(s string, n int) string {
	return ""
}
```
README: why `Builder` beats `+=` in loops (allocations, describe).

- [ ] **Step 6: `06_strconv`**

Test:
```go
package parse

import "testing"

func TestSumStrings(t *testing.T) {
	if got, err := SumStrings("3", "4"); err != nil || got != 7 {
		t.Fatalf("SumStrings = (%d, %v), want (7, nil)", got, err)
	}
	if _, err := SumStrings("x", "4"); err == nil {
		t.Fatal("SumStrings(x) accepted, want error")
	}
}
```
Stub (`parse.go`):
```go
// Package parse teaches string/number conversion.
package parse

// SumStrings parses a and b and returns their sum.
// TODO: use strconv.Atoi on each (import strconv); return wrapped errors as-is.
func SumStrings(a, b string) (int, error) {
	return 0, nil
}
```
README: `strconv.Atoi`/`Itoa`, string/number boundary.

- [ ] **Step 7: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/09_strings/...`
Expected: FAIL (6 packages on assertions).

```bash
git add exercises/09_strings
git commit -m "feat: add 09_strings section (wave 2)"
```

---

### Task 3: New section `10_time` (+4 → 67)

**Files:** New dirs `01_format` (day.go, day_test.go), `02_parse` (day.go, day_test.go), `03_duration` (hours.go, hours_test.go), `04_compare` (expiry.go, expiry_test.go), each + `README.md`. Packages: `dayfmt`, `dayparse`, `hours`, `expiry`.

**Interfaces:**
- Consumes: nothing.
- Produces: `func FormatDay(t time.Time) string`, `func ParseDay(s string) (time.Time, error)`, `func HoursBetween(a, b time.Time) float64`, `var Now func() time.Time` + `func IsExpired(expiry time.Time) bool`.

- [ ] **Step 1: `01_format`**


`day_test.go`:
```go
package dayfmt

import (
	"testing"
	"time"
)

func TestFormatDay(t *testing.T) {
	in := time.Date(2026, 9, 21, 15, 4, 0, 0, time.UTC)
	if got, want := FormatDay(in), "2026-09-21"; got != want {
		t.Fatalf("FormatDay = %q, want %q", got, want)
	}
}
```

Stub (`day.go`):
```go
// Package dayfmt teaches time formatting.
package dayfmt

import "time"

// FormatDay renders t as YYYY-MM-DD.
// TODO: use t.Format with the 2006-01-02 layout.
func FormatDay(t time.Time) string {
	return ""
}
```
README: Go's reference-time layout (`Mon Jan 2 15:04:05 MST 2006`), `time.RFC3339`.

- [ ] **Step 2: `02_parse`**


`day_test.go`:
```go
package dayparse

import (
	"testing"
	"time"
)

func TestParseDay(t *testing.T) {
	want := time.Date(2026, 9, 21, 0, 0, 0, 0, time.UTC)
	got, err := ParseDay("2026-09-21")
	if err != nil || !got.Equal(want) {
		t.Fatalf("ParseDay = (%v, %v), want (%v, nil)", got, err, want)
	}
	if _, err := ParseDay("not-a-date"); err == nil {
		t.Fatal("ParseDay(not-a-date) accepted, want error")
	}
}
```

Stub (`day.go`):
```go
// Package dayparse teaches time parsing.
package dayparse

import "time"

// ParseDay parses YYYY-MM-DD.
// TODO: use time.Parse with the 2006-01-02 layout.
func ParseDay(s string) (time.Time, error) {
	return time.Time{}, nil
}
```
README: layout must match, errors on mismatch.

- [ ] **Step 3: `03_duration`**


`hours_test.go`:
```go
package hours

import (
	"testing"
	"time"
)

func TestHoursBetween(t *testing.T) {
	a := time.Date(2026, 9, 21, 10, 0, 0, 0, time.UTC)
	b := time.Date(2026, 9, 21, 11, 30, 0, 0, time.UTC)
	if got, want := HoursBetween(a, b), 1.5; got != want {
		t.Fatalf("HoursBetween = %v, want %v", got, want)
	}
	if got := HoursBetween(a, a); got != 0 {
		t.Fatalf("HoursBetween(same) = %v, want 0", got)
	}
}
```

Stub (`hours.go`):
```go
// Package hours teaches durations.
package hours

import "time"

// HoursBetween returns (b-a) in hours, possibly fractional.
// TODO: return b.Sub(a).Hours().
func HoursBetween(a, b time.Time) float64 {
	return 0
}
```
README: `Sub`, `Hours()`, `time.Hour` constants.

- [ ] **Step 4: `04_compare`**

`expiry_test.go`:
```go
package expiry

import (
	"testing"
	"time"
)

func TestIsExpired(t *testing.T) {
	old := Now
	Now = func() time.Time { return time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC) }
	defer func() { Now = old }()
	if !IsExpired(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("past expiry not expired")
	}
	if IsExpired(time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("future expiry reported expired")
	}
}
```
Stub (`expiry.go`):
```go
// Package expiry teaches testable time via a fake clock.
package expiry

import "time"

// Now is the clock. Tests replace it with a fixed time; production
// leaves it as time.Now.
// TODO: keep this variable; use Now() inside IsExpired.
var Now = time.Now

// IsExpired reports whether expiry is in the past per Now().
// TODO: return Now().After(expiry).
func IsExpired(expiry time.Time) bool {
	return false
}
```
README: fake-clock variable pattern for deterministic time tests (describe, don't solve).

- [ ] **Step 5: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/10_time/...`
Expected: FAIL (4 packages on assertions; no sleeps, fixed fixtures only).

```bash
git add exercises/10_time
git commit -m "feat: add 10_time section (wave 2)"
```

---

### Task 4: New section `11_json` (+5 → 72, Ticket-flavored)

**Files:** New dirs `01_marshal`, `02_unmarshal`, `03_tags`, `04_strict`, `05_custom`, each `ticket.go`, `ticket_test.go`, `README.md`, `package ticket`. Self-contained Ticket per dir (never import siblings). Base shape: `{ID int \`json:"id"\`, Title string \`json:"title"\`, Description string \`json:"description,omitempty"\`, Status string \`json:"status"\`}` (`05_custom` uses `Status Status` int-based type instead — see step).

**Interfaces:**
- Consumes: nothing.
- Produces: `func MarshalTicket(t Ticket) (string, error)`, `func UnmarshalTicket(data string) (Ticket, error)`, tag-correct `Ticket`, `func UnmarshalStrict(data string) (Ticket, error)`, `type Status int` + `func (s Status) MarshalJSON() ([]byte, error)`.

- [ ] **Step 1: `01_marshal`**

Test round-trips through decode-and-compare (never key order):
```go
in := Ticket{ID: 1, Title: "Fix bug", Description: "Crash on login", Status: "open"}
out, err := MarshalTicket(in)
if err != nil { t.Fatalf(...) }
var got Ticket
if err := json.Unmarshal([]byte(out), &got); err != nil { t.Fatalf(...) }
if !reflect.DeepEqual(got, in) { t.Fatalf(...) }
```
Stub (`ticket.go`, base Ticket shape):
```go
// Package ticket teaches JSON marshaling.
package ticket

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// MarshalTicket encodes t as JSON.
// TODO: use json.Marshal and return string(out) (import encoding/json).
func MarshalTicket(t Ticket) (string, error) {
	return "", nil
}
```
README: marshal, tags control names.

- [ ] **Step 2: `02_unmarshal`**


`ticket_test.go`:
```go
package ticket

import "testing"

func TestUnmarshalTicket(t *testing.T) {
	got, err := UnmarshalTicket(`{"id":1,"title":"Fix bug","status":"open"}`)
	if err != nil {
		t.Fatalf("valid JSON rejected: %v", err)
	}
	if got.Title != "Fix bug" || got.ID != 1 {
		t.Fatalf("decoded = %+v", got)
	}
	if _, err := UnmarshalTicket(`{"id":1,"status":"open"}`); err == nil {
		t.Fatal("missing title accepted, want validation error")
	}
}
```

Stub (`ticket.go`, same base Ticket shape as `01_marshal`):
```go
// UnmarshalTicket decodes JSON and validates: title must be non-empty.
// TODO: json.Unmarshal into Ticket, then check Title (import encoding/json).
func UnmarshalTicket(data string) (Ticket, error) {
	return Ticket{}, nil
}
```
README: unmarshal + validate pattern.

- [ ] **Step 3: `03_tags`**

`ticket.go` (base Ticket shape plus `InternalNote string \`json:"-"\``). Stub
marshal returns hardcoded JSON **with the secret leaked**:
```go
// MarshalTicket encodes t as JSON.
// TODO: use json.Marshal so the struct tags take effect. The hardcoded
// string below leaks InternalNote — real marshaling with json:"-" won't.
func MarshalTicket(t Ticket) (string, error) {
	return `{"id":1,"title":"t","internal_note":"leak"}`, nil
}
```
`ticket_test.go`:
```go
package ticket

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestTags(t *testing.T) {
	out, err := MarshalTicket(Ticket{ID: 1, Title: "t", InternalNote: "leak"})
	if err != nil {
		t.Fatalf("MarshalTicket errored: %v", err)
	}
	if strings.Contains(out, "leak") {
		t.Fatalf("secret leaked into JSON: %s", out)
	}
	var back Ticket
	if err := json.Unmarshal([]byte(out), &back); err != nil || back.Title != "t" {
		t.Fatalf("round-trip = (%+v, %v)", back, err)
	}
}
```
The stub `ticket.go` keeps the base Ticket shape plus
`InternalNote string \`json:"-"\`` and the leaking hardcoded marshal from the
brief. README: `json:"-"`, `omitempty` (describe).

- [ ] **Step 4: `04_strict`**

`ticket_test.go`:
```go
package ticket

import "testing"

func TestUnmarshalStrict(t *testing.T) {
	got, err := UnmarshalStrict(`{"id":1,"title":"t","status":"open"}`)
	if err != nil || got.Title != "t" {
		t.Fatalf("valid JSON = (%+v, %v)", got, err)
	}
	if _, err := UnmarshalStrict(`{"id":1,"title":"t","zzz":1}`); err == nil {
		t.Fatal("unknown field accepted, want strict error")
	}
}
```

Stub (`ticket.go`, base Ticket shape):
```go
import "encoding/json"

// UnmarshalStrict decodes, rejecting unknown fields.
// TODO: use json.Decoder with DisallowUnknownFields instead of Unmarshal.
func UnmarshalStrict(data string) (Ticket, error) {
	var t Ticket
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		return Ticket{}, err
	}
	return t, nil
}
```
README: `Decoder.DisallowUnknownFields`, forward-compat tradeoff (describe).

- [ ] **Step 5: `05_custom`**

`ticket.go`:
```go
type Status int

const (
	StatusOpen Status = iota
	StatusClosed
)

type Ticket struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}

// MarshalJSON renders the status as "open"/"closed".
// TODO: switch on s and return the quoted string (strconv.Quote helps).
func (s Status) MarshalJSON() ([]byte, error) {
	return nil, errors.New("TODO")
}
```
`ticket_test.go`:
```go
package ticket

import (
	"strings"
	"testing"
)

func TestCustomStatus(t *testing.T) {
	out, err := MarshalTicket(Ticket{ID: 1, Title: "t", Status: StatusOpen})
	if err != nil {
		t.Fatalf("MarshalTicket errored: %v", err)
	}
	if !strings.Contains(out, `"status":"open"`) {
		t.Fatalf("output = %s, want status as \"open\"", out)
	}
}
```
The stub `ticket.go` keeps the `Status` type, constants, `Ticket`, and the
`MarshalJSON` TODO from the brief, plus a plain delegating encoder:
```go
// MarshalTicket encodes t as JSON (status via MarshalJSON).
// TODO: use json.Marshal (import encoding/json).
func MarshalTicket(t Ticket) (string, error) {
	return "", nil
}
```
README: custom marshalers, implement `json.Marshaler` (describe).

- [ ] **Step 6: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/11_json/...`
Expected: FAIL (5 packages; assertion/decoder failures, never build).

```bash
git add exercises/11_json
git commit -m "feat: add 11_json section (wave 2)"
```

---

### Task 5: New section `12_io` (+6 → 78, Ticket-flavored)

**Files:** New dirs `01_reader` (readall.go, readall_test.go), `02_writer` (writelines.go, writelines_test.go), `03_scanner` (wcount.go, wcount_test.go), `04_file` (ticketfile.go, ticketfile_test.go), `05_copy` (copyn.go, copyn_test.go), `06_jsonlog` (ticketlog.go, ticketlog_test.go), each + `README.md`. Packages: `readall`, `writelines`, `wcount`, `ticketfile`, `copyn`, `ticketlog`. Ticket shape where needed: `{ID int, Title string}` (+ JSON tags in `06_jsonlog`).

**Interfaces:**
- Consumes: nothing.
- Produces: `func ReadAll(r io.Reader) (string, error)`, `func WriteLines(w io.Writer, lines []string) error`, `func WordCount(r io.Reader) (int, error)`, `func SaveTicket(path string, data []byte) error` + `func LoadTicket(path string) ([]byte, error)`, `func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error)`, `func AppendLog(path string, t Ticket) error` + `func ReadLog(path string) ([]Ticket, error)`.

- [ ] **Step 1: `01_reader`**

`readall_test.go`:
```go
package readall

import (
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	got, err := ReadAll(strings.NewReader("hello"))
	if err != nil || got != "hello" {
		t.Fatalf("ReadAll = (%q, %v), want (hello, nil)", got, err)
	}
}
```

Stub (`readall.go`):
```go
// Package readall teaches the io.Reader contract.
package readall

import "io"

// ReadAll drains r to a string.
// TODO: use io.ReadAll (it reads until EOF).
func ReadAll(r io.Reader) (string, error) {
	return "", nil
}
```
README: `io.Reader` contract (read until EOF), `strings.Reader`/`bytes.Buffer` as test doubles.

- [ ] **Step 2: `02_writer`**


`writelines_test.go`:
```go
package writelines

import (
	"bytes"
	"testing"
)

func TestWriteLines(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteLines(&buf, []string{"a", "b"}); err != nil {
		t.Fatalf("WriteLines errored: %v", err)
	}
	if buf.String() != "a\nb\n" {
		t.Fatalf("buffer = %q, want %q", buf.String(), "a\nb\n")
	}
}
```

Stub (`writelines.go`):
```go
// Package writelines teaches the io.Writer contract.
package writelines

import "io"

// WriteLines writes each line plus "\n".
// TODO: range over lines writing s + "\n" (Fprintf helps).
func WriteLines(w io.Writer, lines []string) error {
	return nil
}
```
README: `io.Writer`, buffering (describe).

- [ ] **Step 3: `03_scanner`**


`wcount_test.go`:
```go
package wcount

import (
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	if got, _ := WordCount(strings.NewReader("go is fun")); got != 3 {
		t.Fatalf("WordCount = %d, want 3", got)
	}
	if got, err := WordCount(strings.NewReader("")); got != 0 || err != nil {
		t.Fatalf("WordCount(empty) = (%d, %v), want (0, nil)", got, err)
	}
}
```

Stub (`wcount.go`):
```go
// Package wcount teaches bufio.Scanner.
package wcount

import "io"

// WordCount counts whitespace-separated words in r.
// TODO: bufio.Scanner with ScanWords; check scanner.Err().
func WordCount(r io.Reader) (int, error) {
	return 0, nil
}
```
README: `bufio.Scanner`, `ScanWords`, scanner.Err().

- [ ] **Step 4: `04_file`**

Test (uses `t.TempDir()`, nothing committed):
```go
dir := t.TempDir()
path := filepath.Join(dir, "ticket.json")
if err := SaveTicket(path, []byte(`{"id":1}`)); err != nil { t.Fatalf(...) }
got, err := LoadTicket(path)
if err != nil || string(got) != `{"id":1}` { t.Fatalf(...) }
```
Stubs:
```go
// SaveTicket writes data to path.
// TODO: os.WriteFile(path, data, 0o644).
func SaveTicket(path string, data []byte) error {
	return nil
}

// LoadTicket reads path back (already correct).
func LoadTicket(path string) ([]byte, error) {
	return os.ReadFile(path)
}
```
Load fails pre-fix because Save never wrote (FAIL via Load error). README: `os.WriteFile`/`ReadFile`, perm bits, `t.TempDir()`.

- [ ] **Step 5: `05_copy`**


`copyn_test.go`:
```go
package copyn

import (
	"bytes"
	"strings"
	"testing"
)

func TestCopyN(t *testing.T) {
	var buf bytes.Buffer
	n, err := CopyN(&buf, strings.NewReader("hello world"), 5)
	if err != nil || n != 5 || buf.String() != "hello" {
		t.Fatalf("CopyN = (%d, %q, %v), want (5, hello, nil)", n, buf.String(), err)
	}
}
```

Stub (`copyn.go`):
```go
// Package copyn teaches io.CopyN.
package copyn

import "io"

// CopyN streams exactly n bytes from src to dst.
// TODO: use io.CopyN (no need to load everything).
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
	return 0, nil
}
```
README: `io.Copy`/`CopyN`, streaming without loading all.

- [ ] **Step 6: `06_jsonlog`**

`ticketlog_test.go`:
```go
package ticketlog

import (
	"path/filepath"
	"testing"
)

func TestTicketLog(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tickets.jsonl")
	if err := AppendLog(path, Ticket{ID: 1, Title: "a"}); err != nil {
		t.Fatalf("AppendLog errored: %v", err)
	}
	if err := AppendLog(path, Ticket{ID: 2, Title: "b"}); err != nil {
		t.Fatalf("AppendLog errored: %v", err)
	}
	got, err := ReadLog(path)
	if err != nil {
		t.Fatalf("ReadLog errored: %v", err)
	}
	if len(got) != 2 || got[0].ID != 1 || got[1].ID != 2 {
		t.Fatalf("ReadLog = %+v, want 2 tickets in order", got)
	}
}
```

Full `ticketlog.go` (`Ticket{ID int \`json:"id"\`, Title string \`json:"title"\`}`):
```go
// Package ticketlog teaches JSON-lines logs.
package ticketlog

import (
	"bufio"
	"encoding/json"
	"os"
)

// Ticket is a logged support request.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// AppendLog appends one ticket as a JSON line (creating the file).
// TODO: open with O_APPEND|O_CREATE|O_WRONLY, encode with json.Encoder.
func AppendLog(path string, t Ticket) error {
	return nil
}

// ReadLog decodes every JSON line, in order.
func ReadLog(path string) ([]Ticket, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []Ticket
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var t Ticket
		if err := json.Unmarshal(sc.Bytes(), &t); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, sc.Err()
}
```
README: JSON-lines format, append mode (`O_APPEND`), line scanner (describe shape, elide AppendLog code).

- [ ] **Step 7: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/12_io/...`
Expected: FAIL (6 packages on assertions/IO errors; repo tree gains no new files — `git status` shows only the 18 intended files).

```bash
git add exercises/12_io
git commit -m "feat: add 12_io section (wave 2)"
```

---

### Task 6: Extend `solutions` branch and release-check Wave 2

**Files:** none new (branch operation + stub fixes only; never edit tests, READMEs, or `testdata/`).

**Interfaces:**
- Consumes: all Wave 2 stubs + tests from Tasks 1–5.
- Produces: `solutions` branch where `go test ./...` passes 78/78.

- [ ] **Step 1: Bring `solutions` forward and apply fixes**

```bash
git checkout solutions
git merge main -m "merge: bring wave 2 stubs onto solutions"
```
(One trivial conflict class possible — reworded lines from Wave 1 fixes; resolve by keeping the `solutions` fixed version. Anything beyond trivial → STOP, report NEEDS_CONTEXT.)
Apply minimal inverse-of-TODO fixes, one commit per new section:
```bash
git add exercises/08_testing && git commit -m "solution: 08_testing wave 2"
git add exercises/09_strings && git commit -m "solution: 09_strings wave 2"
git add exercises/10_time && git commit -m "solution: 10_time wave 2"
git add exercises/11_json && git commit -m "solution: 11_json wave 2"
git add exercises/12_io && git commit -m "solution: 12_io wave 2"
```

- [ ] **Step 2: Verify `solutions` is green**

Run (on `solutions`): `go vet ./... && go build ./... && go test -count=1 ./...`
Expected: PASS, 78/78 packages ok (12 sections: 5+15+8+7+5+6+4+7+6+4+5+6).

- [ ] **Step 3: Verify `main` and finish on `main`**

```bash
git checkout main
```
Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./...`
Expected: FAIL (28 new packages fail by design; 01_mutex still race-only per standing ruling).
Do NOT push.
