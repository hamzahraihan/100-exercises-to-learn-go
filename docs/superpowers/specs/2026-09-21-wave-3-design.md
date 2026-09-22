# Wave 3 Design: HTTP API capstone (78 → 100)

Date: 2026-09-21
Status: approved; details Wave 3 only. Roadmap context:
`docs/superpowers/specs/2026-09-21-expand-to-100-design.md`.
Prior art: v1 (24) + Wave 1 (50) + Wave 2 (78) shipped on `main` + `solutions`.
All prose original; attribution to the Rust source lives in root README.

## 1. Goal
Add the final 22 exercises in 3 new sections (100 total), teaching `net/http`
servers the Rust course has no equivalent for and capping the course with a
working, production-pattern ticket HTTP API server.

## 2. Inventory

| Section | Count | Exercises |
|---|---|---|
| 13_http | 10 | handler basics, status codes, query params, ServeMux routing (Go 1.22+ patterns + PathValue), JSON request/response, httptest unit testing, middleware, request context, HTTP client, structured error responses |
| 14_persistence | 5 | atomic save (temp+rename), load on boot (missing file = empty), corrupt-file errors, JSON export, validated import |
| 15_capstone | 7 | store+handler wiring (create/list), get-one with 404, update with validation, delete (204/404), graceful shutdown, env config with defaults, full e2e (create→get→update→delete) |

## 3. Mechanics

- Test strategy: HTTP handler exercises tested in-memory via `httptest`
  (`NewRecorder`/`NewRequest`); only the capstone e2e binds a real socket via
  `httptest.NewServer` on an ephemeral port. Persistence exercises use
  `t.TempDir()`. Shutdown/config tests use short timeouts with generous
  margins, asserting outcomes never durations.
- Shared Ticket shape across all Wave 3 sections:
  `{ID int \`json:"id"\`, Title string \`json:"title"\`, Description string
  \`json:"description,omitempty"\`, Status string \`json:"status"\`}` — the
  same shape as Wave 2's JSON/IO sections, so knowledge transfers directly.
  Sections never import each other (self-contained types per dir).
- No third-party routers: stdlib `http.ServeMux` with Go 1.22+ method+path
  patterns only.

## 4. Exercise contract (standing rules)

Triplet per exercise (`<name>.go` + `<name>_test.go` + `README.md`); stdlib
`testing` only; stub MUST compile under `go vet`, test MUST fail on assertion
never on build; READMEs show signatures + `// ...` only, never full solutions;
`*.go` + fixtures LF via `.gitattributes`; `go vet`+`go build` green on `main`,
`go test` red on `main` by design, green on `solutions`.

## 5. Process (Wave 3)

One spec → plan → implement → review cycle, subagent-driven with per-task
review. Ends with `solutions` extension (3 fix commits, one per new section)
and release check (solutions 100/100 green; main vet/build green + tests red).
No push without explicit confirmation.

## 6. Non-goals (Wave 3)

No databases (file persistence only), no auth, no TLS, no frontend, no custom
CLI runner, no changes to the 78 shipped exercises except review-driven fixes.
