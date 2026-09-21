# 100 Exercises to Learn Go — Design

Date: 2026-09-21
Source: https://github.com/mainmatter/100-exercises-to-learn-rust
Target: Go-idiomatic port, single module, `go test` verification.

## 1. Goal
Self-paced Go course modelled on Mainmatter's Rust course: small failing-test
exercises ordered into a narrative (calculator → ticket → store → concurrency).
Learner fixes `// TODO` stubs until `go test ./...` is green.

## 2. Non-goals (v1)
- No custom CLI runner (`golearn list/run`). Learner uses `go test <path>`.
- No mdBook site. Prose lives in per-exercise `README.md` (renders on GitHub).
- No 1:1 mirror of Rust-only topics: ownership/borrowck, lifetimes,
  `Deref`/`Sized`/`Drop`/`Copy`/`Clone`, orphan rule, `match` exhaustiveness
  as a language feature, `thiserror`, async runtimes. Each has a Go replacement
  or is dropped with a note in the section outro.
- License: Rust source is CC BY-NC 4.0. All prose is rewritten fresh for Go;
  retain attribution link to the original. Do not copy Rust chapter text.

## 3. Repo layout
```
go.mod            # module 100-exercises-to-learn-go, go >= 1.23 (range-over-func, slices/maps stdlib)
README.md         # setup, toolchain, how to solve, how tests work, FAQ
exercises/
  01_intro/
    00_welcome/       # doc.go + README.md (no test, orientation)
    01_syntax/        # syntax.go + syntax_test.go + README.md
    02_test_workflow/ # learn go test / table tests
  02_calculator/     # ~10 exercises: ints, vars, if/else, panic, for, funcs, overflow, conversions
    NN_name/         # calc.go + calc_test.go + README.md per exercise
  03_ticket_v1/      # ~9: structs, NewTicket, validation, packages/exported, methods, pointer receivers
  04_interfaces/     # ~10: interfaces, implicit impl, Stringer/error, assertions, embedding, generics
  05_ticket_v2/      # ~10: iota enums + sealed interfaces, (T,error), sentinel/wrapped errors, errors.Is/As
  06_ticket_store/   # ~12: arrays, slices, maps, range, slices/maps pkgs, sorting, store CRUD
  07_concurrency/    # ~11: goroutines, WaitGroup, channels, select, Mutex/RWMutex, context, worker pool
.github/workflows/ci.yml  # go vet ./... + go test ./...
.gitignore
```
Exercise contract:
- Stub `.go` file compiles but its test fails until `// TODO` is fixed.
- `*_test.go` uses stdlib `testing` only (table-driven where fitting), no testify
  dependency in v1 to keep `go test` zero-setup.
- `README.md` per exercise: concept (short), task, hints, further reading
  (pkg.go.dev / Tour of Go). Keep each < 250 lines.

## 4. Curriculum mapping (Rust → Go)
| Rust section | Go section | Notes |
|---|---|---|
| 01_intro (welcome, syntax) | 01_intro | `cargo test` → `go test`, `fn` → `func`, `let` → `:=`/`var` |
| 02_basic_calculator (ints, vars, if/else, panic, factorial, while, for, overflow, saturating, as) | 02_calculator | Go has one loop (`for`); overflow wraps silently — teach explicit checks; conversions are explicit `T(x)` |
| 03_ticket_v1 (struct, validation, modules, visibility, encapsulation, ownership, setters, stack/heap, refs, Drop) | 03_ticket_v1 | Exported vs unexported replaces mod/visibility; `NewTicket` constructor; pointer vs value receivers replaces ownership/stack/heap; drop `Drop`/destructors |
| 04_traits | 04_interfaces | Implicit interfaces replace traits; generics (`[T any]`, constraints) replace bounds/`From`/`Clone`; drop orphan/`Deref`/`Sized` |
| 05_ticket_v2 (enums, match, if-let, Option/Result, unwrap, error enums, Error, packages, deps, thiserror, TryFrom, source) | 05_ticket_v2 | `iota` + sealed interfaces replace enums; `(T, error)` + sentinel/wrapped errors + `errors.Is/As` replace `Result`/`thiserror` |
| 06_ticket_management (arrays, vec, iterators, lifetimes, slices, Index, HashMap/BTreeMap) | 06_ticket_store | Slices/maps/range replace vec/iter/lifetimes; `slices`/`maps` stdlib; store CRUD replaces Index traits; drop lifetimes |
| 07_threads + 08_futures | 07_concurrency | Goroutines + channels + `select` + `sync` + `context` replace threads/futures/runtime; worker-pool ticket server as capstone |

Total: ~65 exercises (subset, expandable to 100 later).

## 5. Learner loop
1. Read `exercises/<section>/NN_name/README.md`.
2. Fix `// TODO` in the `.go` stub.
3. Run `go test ./exercises/<section>/NN_name/ -v`.
4. Whole course green: `go test ./...`.

## 6. Solutions strategy
Default: `solutions` branch mirroring `main` path layout (same as Rust repo),
so `main` stays broken-by-design. Alternative deferred: per-folder
`solution.go.txt` if branch workflow proves awkward. v1 implements `main` only;
solutions branch is populated at release.

## 7. Testing / CI
- `go vet ./...` + `go build ./...` must pass on `main` (stubs compile).
  `go test ./...` is expected to FAIL on `main` by design and must pass on the
  `solutions` branch head. CI on `main` runs vet+build only; CI on `solutions`
  runs vet+build+test.
- Prefer table-driven tests; test names `TestX/case`.
- No external test deps in v1.

## 8. Risks
- Over-copying Rust prose → license issue. Mitigation: fresh writing + attribution.
- Scope creep to 100 on first pass. Mitigation: ship 01–02 + one exercise per later
  section as template, then fill in batches.
