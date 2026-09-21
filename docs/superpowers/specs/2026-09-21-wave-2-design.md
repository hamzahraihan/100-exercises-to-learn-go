# Wave 2 Design: Go-native topics (50 → 78)

Date: 2026-09-21
Status: approved; details Wave 2 only. Roadmap context:
`docs/superpowers/specs/2026-09-21-expand-to-100-design.md`.
Prior art: v1 (24) + Wave 1 (50) shipped on `main` + `solutions`.
All prose original; attribution to the Rust source lives in root README.

## 1. Goal
Add 28 exercises in 5 new sections (78 total), teaching Go-native topics the
Rust course has no equivalent for. JSON and IO exercises use Ticket types to
foreshadow the Wave 3 HTTP capstone; testing/strings/time use generic examples.

## 2. Inventory

| Section | Count | Exercises |
|---|---|---|
| 08_testing | 7 | table tests, subtests, helpers (`t.Helper`), testdata/golden files, `Example` output comments, benchmarks, error assertions |
| 09_strings | 6 | `Fields`/`Split`/`Join`, `TrimSpace`/`Trim`, `Contains`/`HasPrefix`, `Replace`/`ReplaceAll`, `Builder`, `strconv` |
| 10_time | 4 | `Format`/`time.RFC3339`, parsing, duration arithmetic, `Add`/`Before` on fixed timestamps |
| 11_json (Ticket) | 5 | marshal ticket, unmarshal+validate, struct tags (`json:"-"`, `omitempty`), strict unknown fields, custom `MarshalJSON` (status as string) |
| 12_io (Ticket) | 6 | `strings.Reader`/`bytes.Buffer`, `bufio.Scanner` word count, write+read ticket file (`t.TempDir()`), `io.Copy`, `io.EOF` handling, JSON-lines ticket log |

## 3. Mechanics

- Meta-testing harness: the benchmark exercise is verified by the harness
  calling `testing.Benchmark` and asserting it ran (never asserting ns/op);
  `Example` exercises rely on normal output-comment matching; the golden-file
  exercise ships a committed `testdata/` fixture. No fuzz, no sleeps in Wave 2.
- Determinism: fixed `time.Date(...)` fixtures plus a fake `Now func()
  time.Time` where "current time" is needed; IO uses `t.TempDir()` (nothing
  written into the repo); JSON assertions decode-and-compare structs, never
  key order.
- Wave 3 preparation: 11_json/12_io define Ticket as
  `{ID int, Title, Description, Status string}` with JSON tags
  (`id,title,description,status`, `omitempty` where taught) — the shape the
  Wave 3 server will bind. Sections never import each other (self-contained
  types per dir, as in v1/Wave 1).

## 4. Exercise contract (standing rules)

Triplet per exercise (`<name>.go` + `<name>_test.go` + `README.md`); stdlib
`testing` only; stub MUST compile under `go vet`, test MUST fail on assertion
never on build; READMEs show signatures + `// ...` only, never full solutions;
`*.go` LF enforced by `.gitattributes`; `go vet`+`go build` green on `main`,
`go test` red on `main` by design, green on `solutions`.

## 5. Process (Wave 2)

One spec → plan → implement → review cycle, subagent-driven with per-task
review. Ends with `solutions` extension (5 fix commits, one per new section)
and release check (solutions 78/78 green; main vet/build green + tests red).
No push without explicit confirmation. Wave 3 (HTTP capstone, +22 → 100) is
out of scope except as the consumer of the Ticket JSON shape above.

## 6. Non-goals (Wave 2)

No HTTP code, no persistence outside `t.TempDir()`, no fuzz tests, no custom
CLI runner, no changes to the 50 shipped exercises except review-driven fixes.
