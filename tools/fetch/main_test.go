package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScanFindsSyntax(t *testing.T) {
	list, err := scanExercises(filepath.Join("..", "..", "exercises"))
	if err != nil {
		t.Fatalf("scanExercises = %v", err)
	}
	found := false
	for _, e := range list {
		if e.Section == "01_intro" && e.Name == "01_syntax" {
			found = true
			if len(e.Files) != 3 {
				t.Fatalf("Files = %v, want 3 entries", e.Files)
			}
		}
	}
	if !found {
		t.Fatal("01_intro/01_syntax not in scan")
	}
}

func TestResolveExactAndAmbiguous(t *testing.T) {
	list := []entry{
		{Section: "01_intro", Name: "01_syntax"},
		{Section: "01_intro", Name: "02_test_workflow"},
	}
	got, err := resolve("01_syntax", list)
	if err != nil || got.Name != "01_syntax" {
		t.Fatalf("resolve exact = %v, %v", got, err)
	}
	got, err = resolve("01_intro/01_syntax", list)
	if err != nil || got.Section != "01_intro" {
		t.Fatalf("resolve qualified = %v, %v", got, err)
	}
	if _, err := resolve("01_", list); err == nil {
		t.Fatal("expected ambiguous error for '01_'")
	}
	if _, err := resolve("nope_xyz", list); err == nil {
		t.Fatal("expected not-found error")
	}
}

func TestCopyAndGoMod(t *testing.T) {
	src := t.TempDir()
	exDir := filepath.Join(src, "01_intro", "01_syntax")
	if err := os.MkdirAll(exDir, 0o755); err != nil {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(exDir, "README.md"), []byte("# T\n"), 0o644)
	os.WriteFile(filepath.Join(exDir, "syntax.go"), []byte("package syntax\n"), 0o644)
	os.WriteFile(filepath.Join(exDir, "syntax_test.go"), []byte("package syntax\n"), 0o644)
	e := entry{Section: "01_intro", Name: "01_syntax", Files: []string{"README.md", "syntax.go", "syntax_test.go"}}
	out := filepath.Join(t.TempDir(), "out")
	if err := copyExercise(src, out, e); err != nil {
		t.Fatal(err)
	}
	if err := writeGoMod(out, "01_syntax"); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(out, "go.mod")); err != nil {
		t.Fatal("go.mod missing")
	}
}

func TestFetchRemote(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("content:" + r.URL.Path))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	e := entry{Section: "01_intro", Name: "01_syntax", Files: []string{"README.md", "syntax.go", "syntax_test.go"}}
	out := filepath.Join(t.TempDir(), "out")
	os.MkdirAll(out, 0o755)
	if err := fetchRemote(srv.Client(), srv.URL, "main", e, out); err != nil {
		t.Fatal(err)
	}
	data, _ := os.ReadFile(filepath.Join(out, "syntax.go"))
	if !strings.Contains(string(data), "content:") {
		t.Fatalf("unexpected body %q", data)
	}
}

func TestResolveBareNameCollision(t *testing.T) {
	list := []entry{
		{Section: "01_intro", Name: "01_syntax"},
		{Section: "02_other", Name: "01_syntax"},
	}
	_, err := resolve("01_syntax", list)
	if err == nil || !strings.Contains(err.Error(), "ambiguous") {
		t.Fatalf("expected ambiguous error, got %v", err)
	}
	if !strings.Contains(err.Error(), "01_intro/01_syntax") || !strings.Contains(err.Error(), "02_other/01_syntax") {
		t.Fatalf("candidates missing in %v", err)
	}
}

func TestFetchRemote404Hint(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer srv.Close()
	e := entry{Section: "01_intro", Name: "01_syntax", Files: []string{"README.md"}}
	err := fetchRemote(srv.Client(), srv.URL, "main", e, filepath.Join(t.TempDir(), "out"))
	if err == nil {
		t.Fatal("expected 404 error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "main") || !strings.Contains(msg, "try --local") {
		t.Fatalf("404 error missing branch/hint: %v", err)
	}
}

func TestParseArgsErrors(t *testing.T) {
	if _, err := parseArgs([]string{"--nope"}); err == nil || !strings.Contains(err.Error(), "unknown flag") {
		t.Fatalf("want unknown flag, got %v", err)
	}
	if _, err := parseArgs([]string{"a", "b"}); err == nil || !strings.Contains(err.Error(), "too many") {
		t.Fatalf("want too many, got %v", err)
	}
	if err := runLocal([]string{"--nope"}); err == nil {
		t.Fatal("runLocal should surface flag error")
	}
}

func TestRunWithEnvLocal(t *testing.T) {
	src := t.TempDir()
	exDir := filepath.Join(src, "01_intro", "01_syntax")
	if err := os.MkdirAll(exDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{"README.md", "syntax.go", "syntax_test.go"} {
		os.WriteFile(filepath.Join(exDir, f), []byte("x"), 0o644)
	}
	out := filepath.Join(t.TempDir(), "out")
	f := fetchEnv{exRoot: src, manifestPath: filepath.Join(src, "missing.json"), client: http.DefaultClient, baseURL: "http://example.invalid"}
	o, err := parseArgs([]string{"01_syntax", "--out", out, "--local"})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWithEnv(o, f); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(out, "syntax.go")); err != nil {
		t.Fatalf("copied file missing: %v", err)
	}
}

func TestRunWithEnvRemoteFallback(t *testing.T) {
	want := []entry{{Section: "01_intro", Name: "01_syntax", Files: []string{"README.md", "syntax.go"}}}
	mux := http.NewServeMux()
	mux.HandleFunc("/main/tools/fetch/manifest.json", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(want)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("remote:" + r.URL.Path))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	out := filepath.Join(t.TempDir(), "out")
	f := fetchEnv{exRoot: filepath.Join(t.TempDir(), "nope"), manifestPath: filepath.Join(t.TempDir(), "nope.json"), client: srv.Client(), baseURL: srv.URL}
	o, err := parseArgs([]string{"01_syntax", "--out", out, "--remote"})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWithEnv(o, f); err != nil {
		t.Fatal(err)
	}
	data, _ := os.ReadFile(filepath.Join(out, "syntax.go"))
	if !strings.Contains(string(data), "remote:") {
		t.Fatalf("remote body missing: %q", data)
	}
}
