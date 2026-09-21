# Expand to 100 Exercises — Design

Date: 2026-09-21
Status: roadmap approved; this spec details Wave 1. Waves 2–3 get their own
spec → plan → implementation cycles later.
Prior art: `docs/superpowers/specs/2026-09-21-go-exercises-design.md` (v1, 24
exercises, shipped on `main` + `solutions`).
Source inspiration: https://github.com/mainmatter/100-exercises-to-learn-rust
(CC BY-NC 4.0 — all prose original with attribution; never copy Rust text).

## 1. Goal
Grow the course from 24 to 100 exercises in 3 waves, keeping the ticket-system
narrative and ending with a working ticket HTTP API server as capstone.

## 2. Target inventory (100 total)

| Section | Now | Target | Wave |
|---|---|---|---|
| 01_intro | 3 | 5 | 1 |
| 02_calculator | 11 | 15 | 1 |
| 03_ticket_v1 | 2 | 8 | 1 |
| 04_interfaces | 2 | 7 | 1 |
| 05_ticket_v2 | 2 | 5 | 1 |
| 06_ticket_store | 2 | 6 | 1 |
| 07_concurrency | 2 | 4 | 1 |
| 08_testing (new) | 0 | 7 | 2 |
| 09_strings (new) | 0 | 6 | 2 |
| 10_time (new) | 0 | 4 | 2 |
| 11_json (new) | 0 | 5 | 2 |
| 12_io (new) | 0 | 6 | 2 |
| 13_http (new) | 0 | 10 | 3 |
| 14_persistence (new) | 0 | 5 | 3 |
| 15_capstone (new) | 0 | 7 | 3 |
| Total | 24 | 100 | — |

Wave sizes: Wave 1 +26 → 50; Wave 2 +28 → 78; Wave 3 +22 → 100.

## 3. Wave 1 scope (+26 → 50; existing sections only)

- 01_intro +2: `03_packages` (imports, exported names across files in one
  dir), `04_tooling` (fix gofmt formatting + a behavior bug; README shows a
  `go vet` printf-mismatch example and the learner runs vet to confirm the
  exercise itself is clean — no live vet failure exists on `main`, since vet
  must stay green).
- 02_calculator +4: `12_floats` (truncation, `math` pkg), `13_runes`
  (byte vs rune, `len` vs rune count), `14_strings` (concat, `==`,
  immutability), `15_bitwise` (`<<`, `&`, `|`, `iota` flag sets).
- 03_ticket_v1 +6: `03_constructor` (validated constructor),
  `04_setters` (pointer-receiver setters), `05_receivers` (value vs pointer
  mutation), `06_zero_value` (usable zero Ticket), `07_string` (String
  method), `08_embed` (struct embedding).
- 04_interfaces +5: `03_any` (`any`, formatting), `04_type_switch`,
  `05_embed_iface` (interface embedding), `06_comparable` (map keys,
  comparable constraint), `07_min_generic` (generic `Min`, `cmp.Ordered`).
- 05_ticket_v2 +3: `03_wrap` (`%w` + `errors.As`), `04_join`
  (`errors.Join`), `05_panic_vs_error` (recover Must-pattern).
- 06_ticket_store +4: `03_slices_fn` (len/cap/copy/clone), `04_sort_filter`,
  `05_pagination` (offset/limit edges), `06_map_idioms` (comma-ok, delete,
  lazy `make`).
- 07_concurrency +2: `03_buffered` (channel semaphore), `04_select`
  (`select` + `time.After`; timing-tolerant test: generous 5s timeout,
  asserts outcome not elapsed time).

## 4. Exercise contract (v1 rules + hardened lessons)

- Triplet per exercise: `<name>.go` + `<name>_test.go` + `README.md`;
  stdlib `testing` only; no external deps.
- Stub MUST compile under `go vet`; test MUST fail on assertion, never on
  build (v1 Task 4 lesson).
- READMEs teach without printing full solutions: signatures + `// ...`
  placeholders only (v1 Task 5 lesson).
- Concurrency tests verified under `-race`; README mandates `-race` where
  plain runs can flake (v1 Task 8 lesson).
- Timing tests assert outcomes, never elapsed durations.
- `go vet ./...` + `go build ./...` green on `main`; `go test ./...` red
  on `main` by design, green on `solutions`.

## 5. Process (Wave 1)

One spec → plan → implement → review cycle, subagent-driven with per-task
review (same as v1). Ends with `solutions`-branch extension (one fix commit
per touched section) and release check (solutions 50/50 green incl. `-race`
for concurrency; main vet/build green + tests red). Waves 2–3 out of scope
except as roadmap above. An `origin` remote now exists; push only on explicit
confirmation.

## 6. Non-goals (Wave 1)

- No new top-level sections; no HTTP server code yet.
- No custom CLI runner; no mdBook/site (per-exercise READMEs stand).
- No changes to the 24 shipped exercises except fixes required by review.
