# 100 Exercises to Learn Go Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Scaffold the Go course repo and ship v1: intro + calculator sections complete plus pattern exercises for ticket/interfaces/errors/store/concurrency.

**Architecture:** Single Go module `learngo`; each exercise is a directory with a broken-by-design stub (compiles, test fails), a stdlib `testing` test file, and a `README.md` lesson. Learner loop is `go test <exercise path>`.

**Tech Stack:** Go >= 1.23, stdlib `testing` only (no testify), GitHub Actions (`go vet`, `go build`, `go test`).

## Global Constraints

- Go floor: `go 1.23` in `go.mod`; verified with local toolchain go1.27.1.
- Module path: `learngo` (exact, lowercase, no dashes).
- Test deps: stdlib `testing` only; no external dependencies in v1.
- Exercise contract: stub `.go` file MUST compile; its test MUST fail until `// TODO` is fixed.
- `go vet ./...` + `go build ./...` MUST pass on `main`; `go test ./...` is expected to FAIL on `main` and pass on `solutions` branch.
- All prose is written fresh for Go with attribution link to the Rust source; do not copy Rust chapter text (CC BY-NC 4.0).
- Each exercise dir has exactly: `<name>.go`, `<name>_test.go`, `README.md`.

---

### Task 1: Scaffold module, README, gitignore, CI

**Files:**
- Create: `go.mod`
- Create: `.gitignore`
- Create: `README.md`
- Create: `.github/workflows/ci.yml`

**Interfaces:**
- Consumes: nothing (first task).
- Produces: module path `learngo` used by all later tasks; CI commands `go vet ./...`, `go build ./...`, `go test ./...`.

- [ ] **Step 1: Write `go.mod`**

```text
module learngo

go 1.23
```

- [ ] **Step 2: Write `.gitignore`**

```text
# Binaries
*.exe
*.out
/bin/

# Go build cache / test cache are outside the repo by default; keep this minimal
.DS_Store
```

- [ ] **Step 3: Write root `README.md`**

```markdown
# 100 Exercises to Learn Go

Self-paced Go course inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) (CC BY-NC 4.0, all prose here is original).

## Requirements

- Go 1.23+ (check with `go version`)

## How it works

Each exercise lives in `exercises/<section>/NN_name/` with three files:

- `README.md` — the concept lesson and your task
- `<name>.go` — a stub with a `// TODO:` for you to fix (it compiles, but the test fails)
- `<name>_test.go` — the test that must pass

Solve one exercise at a time:

```bash
go test ./exercises/01_intro/01_syntax/ -v
```

Check everything (expected to FAIL until you solve all exercises):

```bash
go test ./...
```

Always-safe checks (must pass even with unsolved exercises):

```bash
go vet ./...
go build ./...
```

## Solutions

Solutions live on the `solutions` branch, same paths as `main`.
```

- [ ] **Step 4: Write `.github/workflows/ci.yml`**

```yaml
name: ci
on:
  push:
    branches: [main, solutions]
  pull_request:
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.23"
      - run: go vet ./...
      - run: go build ./...
      - run: go test ./...
        if: github.ref == 'refs/heads/solutions'
```

- [ ] **Step 5: Verify scaffold builds**

Run: `go vet ./... && go build ./...`
Expected: PASS (no packages yet, exit 0).

- [ ] **Step 6: Commit**

```bash
git add go.mod .gitignore README.md .github/workflows/ci.yml
git commit -m "feat: scaffold go module, readme, and ci"
```

---

### Task 2: Section `01_intro` (3 exercises)

**Files:**
- Create: `exercises/01_intro/00_welcome/welcome.go`, `welcome_test.go`, `README.md`
- Create: `exercises/01_intro/01_syntax/syntax.go`, `syntax_test.go`, `README.md`
- Create: `exercises/01_intro/02_test_workflow/add.go`, `add_test.go`, `README.md`

**Interfaces:**
- Consumes: module `learngo` from Task 1.
- Produces: package paths `learngo/exercises/01_intro/00_welcome`, `.../01_syntax`, `.../02_test_workflow`; functions `welcome.Message() string`, `syntax.Compute(a, b int) int`, `workflow.Add(a, b int) int`.

- [ ] **Step 1: Write failing test `exercises/01_intro/00_welcome/welcome_test.go`**

```go
package welcome

import "testing"

func TestMessage(t *testing.T) {
	got := Message()
	if got == "" {
		t.Fatal("Message() returned empty string, expected a greeting")
	}
}
```

- [ ] **Step 2: Write broken stub `exercises/01_intro/00_welcome/welcome.go`**

```go
// Package welcome orients you in the course.
package welcome

// Message returns a greeting for the learner.
// TODO: return a non-empty greeting, e.g. "Welcome to 100 Exercises to Learn Go!".
func Message() string {
	return ""
}
```

- [ ] **Step 3: Write `exercises/01_intro/00_welcome/README.md`**

```markdown
# Welcome

Welcome to 100 Exercises to Learn Go, inspired by Mainmatter's Rust course.

## Task

Open `welcome.go` and make `Message` return a non-empty greeting.

## Check

```bash
go test ./exercises/01_intro/00_welcome/ -v
```
```

- [ ] **Step 4: Write failing test `exercises/01_intro/01_syntax/syntax_test.go`**

```go
package syntax

import "testing"

func TestCompute(t *testing.T) {
	if got, want := Compute(2, 3), 5; got != want {
		t.Fatalf("Compute(2, 3) = %d, want %d", got, want)
	}
}
```

- [ ] **Step 5: Write broken stub `exercises/01_intro/01_syntax/syntax.go`**

```go
// Package syntax teaches basic Go function syntax.
package syntax

// Compute adds a and b.
// TODO: return a+b instead of 0.
func Compute(a, b int) int {
	return 0
}
```

- [ ] **Step 6: Write `exercises/01_intro/01_syntax/README.md`**

```markdown
# Syntax

Go functions look like this:

```go
func Compute(a, b int) int {
    return a + b
}
```

`package syntax` declares the package. `func` declares a function,
`(a, b int)` are typed parameters, the trailing `int` is the return type.

## Task

Fix `Compute` in `syntax.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/01_syntax/ -v
```
```

- [ ] **Step 7: Write failing test `exercises/01_intro/02_test_workflow/add_test.go`**

```go
package workflow

import "testing"

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

- [ ] **Step 8: Write broken stub `exercises/01_intro/02_test_workflow/add.go`**

```go
// Package workflow teaches the go test loop and table-driven tests.
package workflow

// Add returns a+b.
// TODO: return a+b instead of 0.
func Add(a, b int) int {
	return 0
}
```

- [ ] **Step 9: Write `exercises/01_intro/02_test_workflow/README.md`**

```markdown
# The test workflow

Go tests are table-driven: a slice of inputs and wants, looped with `t.Errorf`
on mismatch. You run one exercise with `go test <path> -v`.

## Task

Fix `Add` in `add.go`.

## Check

```bash
go test ./exercises/01_intro/02_test_workflow/ -v
```
```

- [ ] **Step 10: Verify new tests fail and vet/build pass**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/01_intro/...`
Expected: FAIL (3 failing packages — proves stubs are broken-by-design).

- [ ] **Step 11: Commit**

```bash
git add exercises/01_intro
git commit -m "feat: add 01_intro exercises (welcome, syntax, test workflow)"
```

---

### Task 3: Section `02_calculator` (10 exercises)

**Files:** For each `NN` below, create `exercises/02_calculator/NN_name/<file>.go`, `<file>_test.go`, `README.md`. Package name in every dir: `package calc`.

**Interfaces:**
- Consumes: nothing from Task 2 (independent section).
- Produces: `func Compute(a, b uint32) uint32`, `func Double(x int) int`, `func Max(a, b int) int`, `func Divide(a, b int) int` (panics on zero), `func Factorial(n int) int`, `func SumTo(n int) int`, `func SumRange(vals []int) int`, `func AddUint8(a, b uint8) (uint8, bool)`, `func SaturatingAdd(a, b uint8) uint8`, `func ToInt64(x int32) int64`, `func DivMod(a, b int) (int, int)`.

- [ ] **Step 1: `01_integers` — type-mismatch port of the Rust exercise**

`calc_test.go`:

```go
package calc

import "testing"

func TestCompute(t *testing.T) {
	if got, want := Compute(1, 2), uint32(9); got != want {
		t.Fatalf("Compute(1, 2) = %d, want %d", got, want)
	}
}
```

`calc.go` (broken: drops the multiplier term and ignores the uint8→uint32 conversion lesson):

```go
package calc

// Compute returns a + b*multiplier where multiplier is 4.
// TODO: multiply b by the uint8 multiplier (convert it to uint32 first).
func Compute(a, b uint32) uint32 {
	var multiplier uint8 = 4
	_ = multiplier
	return a + b
}
```

`README.md`:

```markdown
# Integers

Go has `int`, `int8/16/32/64`, `uint`, `uint8/16/32/64`, `uintptr`.
Unlike Rust there is no `u128`, and `int` is 32 or 64 bits depending on platform.
Go never implicitly converts between integer types: `uint32(m)` is required.

## Task

Use `multiplier` (a `uint8`) in the computation. Convert it to `uint32` first.

## Check

```bash
go test ./exercises/02_calculator/01_integers/ -v
```
```

Solution (for `solutions` branch only): `return a + b*uint32(multiplier)`.

- [ ] **Step 2: `02_variables` — short declaration**

Test:

```go
package calc

import "testing"

func TestDouble(t *testing.T) {
	if got, want := Double(21), 42; got != want {
		t.Fatalf("Double(21) = %d, want %d", got, want)
	}
}
```

Stub `calc.go`:

```go
package calc

// Double returns twice x.
// TODO: use := to declare doubled, then return it.
func Double(x int) int {
	return 0
}
```

README: one paragraph on `var x int = 1` vs `x := 1`.

- [ ] **Step 3: `03_if_else` — Max**

Test (`calc_test.go`):

```go
package calc

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{3, 7, 7},
		{7, 3, 7},
		{4, 4, 4},
		{-1, -5, -1},
	}
	for _, tt := range tests {
		if got := Max(tt.a, tt.b); got != tt.want {
			t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
```

Stub (`calc.go`):

```go
package calc

// Max returns the larger of a and b.
// TODO: write an if/else returning a when a >= b, else b.
func Max(a, b int) int {
	return 0
}
```

README: one paragraph on `if/else` (braces required, no parens around condition).

- [ ] **Step 4: `04_panics` — Divide panics on zero**

Test:

```go
package calc

import "testing"

func TestDivide(t *testing.T) {
	if got := Divide(6, 3); got != 2 {
		t.Fatalf("Divide(6, 3) = %d, want 2", got)
	}
}

func TestDivideByZeroPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Divide(1, 0) did not panic")
		}
	}()
	_ = Divide(1, 0)
}
```

Stub:

```go
package calc

// Divide returns a/b, panicking on division by zero.
// TODO: panic("division by zero") when b == 0, else return a/b.
func Divide(a, b int) int {
	return 0
}
```

- [ ] **Step 5: `05_factorial`**

Test (`calc_test.go`):

```go
package calc

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
	}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
```

Stub (`calc.go`):

```go
package calc

// Factorial returns n! with Factorial(0) == 1.
// TODO: loop with for, multiplying result by 1..n.
func Factorial(n int) int {
	return 0
}
```

- [ ] **Step 6: `06_sum_to` (merges Rust while+for: Go has only `for`)**

Test (`calc_test.go`):

```go
package calc

import "testing"

func TestSumTo(t *testing.T) {
	if got, want := SumTo(100), 5050; got != want {
		t.Fatalf("SumTo(100) = %d, want %d", got, want)
	}
	if got := SumTo(0); got != 0 {
		t.Fatalf("SumTo(0) = %d, want 0", got)
	}
}
```

Stub (`calc.go`):

```go
package calc

// SumTo returns 1+2+...+n (0 for n <= 0). Go has only `for`:
// it covers while-style loops too.
// TODO: for i := 1; i <= n; i++ { total += i }.
func SumTo(n int) int {
	return 0
}
```

- [ ] **Step 7: `07_sum_range` (for-range)**

Test (`calc_test.go`):

```go
package calc

import "testing"

func TestSumRange(t *testing.T) {
	if got, want := SumRange([]int{1, 2, 3, 4}), 10; got != want {
		t.Fatalf("SumRange = %d, want %d", got, want)
	}
	if got := SumRange(nil); got != 0 {
		t.Fatalf("SumRange(nil) = %d, want 0", got)
	}
}
```

Stub (`calc.go`):

```go
package calc

// SumRange adds all slice elements.
// TODO: for _, v := range vals { total += v }.
func SumRange(vals []int) int {
	return 0
}
```

- [ ] **Step 8: `08_overflow` — detect uint8 overflow**

Test:

```go
package calc

import "testing"

func TestAddUint8(t *testing.T) {
	if got, overflow := AddUint8(100, 100); got != 200 || overflow {
		t.Fatalf("AddUint8(100,100) = (%d,%v), want (200,false)", got, overflow)
	}
	if _, overflow := AddUint8(200, 100); !overflow {
		t.Fatal("AddUint8(200,100) should report overflow")
	}
}
```

Stub:

```go
package calc

// AddUint8 adds a and b, reporting wraparound. Go wraps silently,
// so callers must check: overflow is true when the result wrapped.
// TODO: compute sum and set overflow = sum < a.
func AddUint8(a, b uint8) (sum uint8, overflow bool) {
	return a + b, false
}
```

README explains Go wraps (unlike Rust debug panic) so explicit checks are idiomatic.

- [ ] **Step 9: `09_saturating`, `10_conversions`, `11_divmod`**

`09_saturating` test (`calc_test.go`):

```go
package calc

import "testing"

func TestSaturatingAdd(t *testing.T) {
	if got := SaturatingAdd(10, 20); got != 30 {
		t.Fatalf("SaturatingAdd(10,20) = %d, want 30", got)
	}
	if got := SaturatingAdd(200, 100); got != 255 {
		t.Fatalf("SaturatingAdd(200,100) = %d, want 255 (saturated)", got)
	}
}
```

`09_saturating` stub (`calc.go`):

```go
package calc

// SaturatingAdd adds clamped at 255 instead of wrapping.
// TODO: reuse the overflow check: if sum < a, return 255.
func SaturatingAdd(a, b uint8) uint8 {
	return a + b
}
```

`10_conversions` test (`calc_test.go`):

```go
package calc

import "testing"

func TestToInt64(t *testing.T) {
	if got := ToInt64(-42); got != -42 {
		t.Fatalf("ToInt64(-42) = %d, want -42", got)
	}
}
```

`10_conversions` stub (`calc.go`):

```go
package calc

// ToInt64 converts x explicitly. Go has no implicit numeric
// conversions (unlike some languages) and no `as` casts:
// every conversion is written T(x).
// TODO: return int64(x).
func ToInt64(x int32) int64 {
	return 0
}
```

`11_divmod` test (`calc_test.go`):

```go
package calc

import "testing"

func TestDivMod(t *testing.T) {
	q, r := DivMod(7, 3)
	if q != 2 || r != 1 {
		t.Fatalf("DivMod(7,3) = (%d,%d), want (2,1)", q, r)
	}
}
```

`11_divmod` stub (`calc.go`):

```go
package calc

// DivMod returns quotient and remainder. Integer division
// truncates toward zero: -5/2 == -2.
// TODO: return a / b, a % b.
func DivMod(a, b int) (quotient, remainder int) {
	return 0, 0
}
```

Each of the three gets a 10–20 line README with a concept paragraph, a Task section, and a Check section running `go test ./exercises/02_calculator/NN_name/ -v`.

- [ ] **Step 10: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/02_calculator/...`
Expected: FAIL (all packages fail — broken-by-design).
Run a solution spot-check mentally: `Compute` fix `a + b*uint32(multiplier)` compiles.

```bash
git add exercises/02_calculator
git commit -m "feat: add 02_calculator exercises"
```

---

### Task 4: Section `03_ticket_v1` pattern (2 exercises)

**Files:**
- Create: `exercises/03_ticket_v1/01_struct/ticket.go`, `ticket_test.go`, `README.md`
- Create: `exercises/03_ticket_v1/02_validation/ticket.go`, `ticket_test.go`, `README.md`

**Interfaces:**
- Consumes: nothing new.
- Produces: `type Ticket struct { Title, Description string }`, `func NewTicket(title, description string) Ticket`, `func (t Ticket) Validate() error`.

- [ ] **Step 1: `01_struct` test + stub**

Test:

```go
package ticket

import "testing"

func TestNewTicket(t *testing.T) {
	tk := NewTicket("Fix bug", "Crash on login")
	if tk.Title != "Fix bug" || tk.Description != "Crash on login" {
		t.Fatalf("got %+v", tk)
	}
}
```

Stub (`ticket.go` — human ruling 2026-09-21: fields present, constructor broken,
so vet stays green and the test fails on assertion rather than on build):

```go
// Package ticket models a support ticket.
package ticket

// Ticket is a support request with a title and description.
// TODO: populate and return the struct in NewTicket.
type Ticket struct {
	Title       string
	Description string
}

// NewTicket builds a Ticket.
// TODO: return Ticket{Title: title, Description: description} instead of Ticket{}.
func NewTicket(title, description string) Ticket {
	return Ticket{}
}
```

README: structs, struct literals, exported fields (capitalized = visible outside package — replaces Rust visibility/encapsulation lesson); Task wording: populate `NewTicket`.

- [ ] **Step 2: `02_validation` test + stub**

Test:

```go
package ticket

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	if err := NewTicket("ok title", "long enough description").Validate(); err != nil {
		t.Fatalf("valid ticket rejected: %v", err)
	}
	if err := NewTicket("", "long enough description").Validate(); err == nil {
		t.Fatal("empty title accepted")
	}
	if err := NewTicket("ok", "x").Validate(); err == nil || !strings.Contains(err.Error(), "description") {
		t.Fatalf("short description should error mentioning description, got %v", err)
	}
}
```

Stub:

```go
// Package ticket models a support ticket with validation.
package ticket

import "errors"

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// NewTicket builds a Ticket.
func NewTicket(title, description string) Ticket {
	return Ticket{Title: title, Description: description}
}

// Validate checks the ticket.
// TODO: return an error when Title is empty or Description is shorter than 10 chars.
func (t Ticket) Validate() error {
	return errors.New("TODO")
}
```

Note: stub returns non-nil error so the valid-ticket case fails (broken-by-design) while compiling.

- [ ] **Step 3: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/03_ticket_v1/...`
Expected: FAIL.

```bash
git add exercises/03_ticket_v1
git commit -m "feat: add 03_ticket_v1 struct and validation exercises"
```

---

### Task 5: Section `04_interfaces` pattern (2 exercises)

**Files:**
- Create: `exercises/04_interfaces/01_stringer/order.go`, `order_test.go`, `README.md`
- Create: `exercises/04_interfaces/02_generic_sum/sum.go`, `sum_test.go`, `README.md`

**Interfaces:**
- Consumes: nothing new.
- Produces: `func (o Order) String() string` (implements `fmt.Stringer` implicitly), `func Sum[T int64 | float64](vals []T) T` — generic sum replacing Rust trait-bounds lesson.

- [ ] **Step 1: `01_stringer` test + stub**

Test:

```go
package order

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringer(t *testing.T) {
	o := Order{ID: 7, Item: "book"}
	if got := fmt.Sprint(o); !strings.Contains(got, "7") || !strings.Contains(got, "book") {
		t.Fatalf("String() = %q, want id and item inside", got)
	}
}
```

Stub:

```go
// Package order teaches implicit interfaces via fmt.Stringer.
package order

// Order is a purchase.
type Order struct {
	ID   int
	Item string
}

// String implements fmt.Stringer. Go interfaces are implicit:
// no "implements" keyword (unlike Rust trait impls).
// TODO: return a string containing the id and item, e.g. fmt.Sprintf("Order %d: %s", o.ID, o.Item).
func (o Order) String() string {
	return ""
}
```

- [ ] **Step 2: `02_generic_sum` test + stub**

Test:

```go
package gsum

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int64{1, 2, 3}); got != 6 {
		t.Fatalf("Sum ints = %d, want 6", got)
	}
	if got := Sum([]float64{1.5, 2.5}); got != 4.0 {
		t.Fatalf("Sum floats = %v, want 4.0", got)
	}
}
```

Stub:

```go
// Package gsum teaches generics (the Go answer to Rust trait bounds).
package gsum

// Sum adds all values. The constraint permits int64 and float64.
// TODO: range over vals and accumulate.
func Sum[T int64 | float64](vals []T) T {
	var total T
	return total
}
```

README (human ruling 2026-09-21: elide the loop — show signature + `// ...`
placeholder only, never the full accumulation loop, so the lesson doesn't
hand out the answer).

- [ ] **Step 3: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/04_interfaces/...`
Expected: FAIL.

```bash
git add exercises/04_interfaces
git commit -m "feat: add 04_interfaces stringer and generics exercises"
```

---

### Task 6: Section `05_ticket_v2` pattern (2 exercises)

**Files:**
- Create: `exercises/05_ticket_v2/01_status/status.go`, `status_test.go`, `README.md`
- Create: `exercises/05_ticket_v2/02_errors/ticket.go`, `ticket_test.go`, `README.md`

**Interfaces:**
- Consumes: `Ticket` shape from Task 4 (copied, not imported — exercises stay independent).
- Produces: `type Status int` with `StatusOpen/StatusInProgress/StatusClosed` + `func (s Status) Valid() bool`; `var ErrUnknownStatus`, `func NewTicketWithStatus(...) (Ticket, error)` using `fmt.Errorf("...: %w", ...)`.

- [ ] **Step 1: `01_status` test + stub**

Test:

```go
package status

import "testing"

func TestStatusValid(t *testing.T) {
	for _, s := range []Status{StatusOpen, StatusInProgress, StatusClosed} {
		if !s.Valid() {
			t.Fatalf("Status(%d).Valid() = false, want true", int(s))
		}
	}
	if Status(99).Valid() {
		t.Fatal("Status(99).Valid() = true, want false")
	}
}
```

Stub:

```go
// Package status teaches iota enums (Go's answer to simple Rust enums).
package status

// Status is a ticket state.
// TODO: define Status as int-based type with iota constants
// StatusOpen, StatusInProgress, StatusClosed, and a Valid() method.
type Status int

const (
	StatusOpen Status = iota
	StatusInProgress
	StatusClosed
)

// Valid reports whether s is a known status.
func (s Status) Valid() bool {
	return false // TODO: switch on s
}
```

- [ ] **Step 2: `02_errors` test + stub**

Test:

```go
package ticket

import (
	"errors"
	"testing"
)

func TestNewTicketWithStatusRejectsUnknown(t *testing.T) {
	_, err := NewTicketWithStatus("t", "long enough description", 99)
	if !errors.Is(err, ErrUnknownStatus) {
		t.Fatalf("err = %v, want errors.Is(err, ErrUnknownStatus)", err)
	}
}

func TestNewTicketWithStatusOK(t *testing.T) {
	tk, err := NewTicketWithStatus("t", "long enough description", StatusOpen)
	if err != nil || tk.Status != StatusOpen {
		t.Fatalf("got %+v, %v", tk, err)
	}
}
```

Stub (`ticket.go` with `Status int`, `ErrUnknownStatus = errors.New("unknown status")`, `NewTicketWithStatus` returning `fmt.Errorf("...: %w", ErrUnknownStatus)` as TODO — stub returns `errors.New("TODO")` so `errors.Is` fails):

```go
// Package ticket teaches sentinel + wrapped errors (errors.Is/As).
package ticket

import "errors"

// Status is a ticket state.
type Status int

const (
	StatusOpen Status = iota
	StatusInProgress
	StatusClosed
)

// ErrUnknownStatus is returned for bad status values.
var ErrUnknownStatus = errors.New("unknown status")

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
	Status      Status
}

// NewTicketWithStatus builds a Ticket or fails.
// TODO: validate title/description like 03_ticket_v1 and wrap
// ErrUnknownStatus with %w when status is unknown.
func NewTicketWithStatus(title, description string, status Status) (Ticket, error) {
	return Ticket{}, errors.New("TODO")
}
```

- [ ] **Step 3: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/05_ticket_v2/...`
Expected: FAIL.

```bash
git add exercises/05_ticket_v2
git commit -m "feat: add 05_ticket_v2 status and errors exercises"
```

---

### Task 7: Section `06_ticket_store` pattern (2 exercises)

**Files:**
- Create: `exercises/06_ticket_store/01_slice_store/store.go`, `store_test.go`, `README.md`
- Create: `exercises/06_ticket_store/02_map_store/store.go`, `store_test.go`, `README.md`

**Interfaces:**
- Consumes: nothing (self-contained `Ticket{ID int, Title string}`).
- Produces: `type Store struct{...}`, `func (s *Store) Add(title string) int`, `func (s *Store) Get(id int) (Ticket, bool)`.

- [ ] **Step 1: `01_slice_store` test + stub**

Test:

```go
package store

import "testing"

func TestSliceStore(t *testing.T) {
	var s Store
	id := s.Add("first")
	tk, ok := s.Get(id)
	if !ok || tk.Title != "first" {
		t.Fatalf("Get(%d) = %+v, %v", id, tk, ok)
	}
	if _, ok := s.Get(999); ok {
		t.Fatal("Get(999) found ticket, want missing")
	}
}
```

Stub:

```go
// Package store teaches slices as growable storage.
package store

// Ticket is a stored item.
type Ticket struct {
	ID    int
	Title string
}

// Store keeps tickets in a slice.
type Store struct {
	tickets []Ticket
	nextID  int
}

// Add appends a ticket and returns its id.
// TODO: use append, assign s.nextID then increment it.
func (s *Store) Add(title string) int {
	return 0
}

// Get finds a ticket by id.
// TODO: range over s.tickets.
func (s *Store) Get(id int) (Ticket, bool) {
	return Ticket{}, false
}
```

- [ ] **Step 2: `02_map_store` test + stub**

Test (`store_test.go`):

```go
package store

import "testing"

func TestMapStore(t *testing.T) {
	var s Store
	id := s.Add("first")
	tk, ok := s.Get(id)
	if !ok || tk.Title != "first" {
		t.Fatalf("Get(%d) = %+v, %v", id, tk, ok)
	}
	if _, ok := s.Get(999); ok {
		t.Fatal("Get(999) found ticket, want missing")
	}
}
```

Stub (`store.go`):

```go
// Package store teaches maps as indexed storage.
package store

// Ticket is a stored item.
type Ticket struct {
	ID    int
	Title string
}

// Store keeps tickets in a map.
type Store struct {
	tickets map[int]Ticket
	nextID  int
}

// Add inserts a ticket and returns its id.
// TODO: lazily make(s.tickets) when nil, then store and bump nextID.
func (s *Store) Add(title string) int {
	return 0
}

// Get finds a ticket by id using the comma-ok idiom.
// TODO: look up s.tickets[id].
func (s *Store) Get(id int) (Ticket, bool) {
	return Ticket{}, false
}
```

- [ ] **Step 3: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/06_ticket_store/...`
Expected: FAIL.

```bash
git add exercises/06_ticket_store
git commit -m "feat: add 06_ticket_store slice and map exercises"
```

---

### Task 8: Section `07_concurrency` pattern (2 exercises)

**Files:**
- Create: `exercises/07_concurrency/01_mutex/counter.go`, `counter_test.go`, `README.md`
- Create: `exercises/07_concurrency/02_channels/pipeline.go`, `pipeline_test.go`, `README.md`

**Interfaces:**
- Consumes: nothing.
- Produces: `type Counter struct{...}` with `func (c *Counter) Add(n int)`, `func (c *Counter) Value() int` (Mutex-guarded); `func DoubleAll(nums []int) []int` (goroutine + channel pipeline, order-preserving).

- [ ] **Step 1: `01_mutex` test + stub**

Test:

```go
package counter

import (
	"sync"
	"testing"
)

func TestConcurrentAdds(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Add(1) }()
	}
	wg.Wait()
	if got := c.Value(); got != 100 {
		t.Fatalf("Value() = %d, want 100 (hint: sync.Mutex)", got)
	}
}
```

Stub:

```go
// Package counter teaches sync.Mutex for shared state.
package counter

import "sync"

// Counter is a goroutine-safe int.
// TODO: add a sync.Mutex field and guard n.
type Counter struct {
	mu sync.Mutex
	n  int
}

// Add increments by n.
func (c *Counter) Add(n int) {
	c.n += n // TODO: Lock/Unlock
}

// Value returns the current count.
func (c *Counter) Value() int {
	return c.n // TODO: Lock/Unlock
}
```

README must warn: test with `-race`: `go test -race ./exercises/07_concurrency/01_mutex/ -v`.

- [ ] **Step 2: `02_channels` test + stub**

Test:

```go
package pipeline

import (
	"reflect"
	"testing"
)

func TestDoubleAll(t *testing.T) {
	got := DoubleAll([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(got, []int{2, 4, 6, 8}) {
		t.Fatalf("DoubleAll = %v, want [2 4 6 8]", got)
	}
}
```

Stub:

```go
// Package pipeline teaches goroutines + channels.
package pipeline

// DoubleAll doubles each number using a goroutine and channel,
// preserving input order.
// TODO: launch a goroutine per index (or one worker), send results
// over a channel, collect in order.
func DoubleAll(nums []int) []int {
	out := make([]int, len(nums))
	return out
}
```

- [ ] **Step 3: Verify (including race detector) and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test -race ./exercises/07_concurrency/...`
Expected: FAIL (counter wrong under concurrency; pipeline returns zeros).

```bash
git add exercises/07_concurrency
git commit -m "feat: add 07_concurrency mutex and channels exercises"
```

---

### Task 9: Solutions branch and release check

**Files:** none new (branch operation + doc touch-up if needed).

**Interfaces:**
- Consumes: all exercise stubs + tests from Tasks 2–8.
- Produces: `solutions` branch where `go test ./...` passes.

- [ ] **Step 1: Create solutions branch**

```bash
git branch solutions
```

- [ ] **Step 2: On `solutions`, apply the minimal fixes** (one commit per section, e.g. `01_integers`: `return a + b*uint32(multiplier)`; `Double`: `doubled := x * 2; return doubled`; etc. — each fix is the removal of the `// TODO` noted in the task).

- [ ] **Step 3: Verify solutions branch is green**

Run (on `solutions`): `go vet ./... && go build ./... && go test ./...`
Expected: PASS, all packages ok.

- [ ] **Step 4: Verify main is red-green-correct**

Run (on `main`): `go vet ./... && go build ./...`
Expected: PASS.
Run (on `main`): `go test ./...`
Expected: FAIL (proves exercises are unsolved-by-design).

- [ ] **Step 5: Push/announce** — report branch state; do not push unless a remote exists.
