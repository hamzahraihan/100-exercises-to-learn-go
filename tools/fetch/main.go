// tools/fetch/main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type entry struct {
	Section string   `json:"section"`
	Name    string   `json:"name"`
	Files   []string `json:"files"`
}

func scanExercises(root string) ([]entry, error) {
	var out []entry
	sections, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, s := range sections {
		if !s.IsDir() {
			continue
		}
		exDirs, err := os.ReadDir(filepath.Join(root, s.Name()))
		if err != nil {
			return nil, err
		}
		for _, e := range exDirs {
			if !e.IsDir() {
				continue
			}
			dir := filepath.Join(root, s.Name(), e.Name())
			ents, err := os.ReadDir(dir)
			if err != nil {
				return nil, err
			}
			var files []string
			for _, f := range ents {
				n := f.Name()
				if n == "README.md" || strings.HasSuffix(n, ".go") {
					files = append(files, n)
				}
			}
			if len(files) < 3 {
				return nil, fmt.Errorf("missing files in %s", dir)
			}
			out = append(out, entry{Section: s.Name(), Name: e.Name(), Files: files})
		}
	}
	return out, nil
}

func dumpManifest(root, dest string) error {
	list, err := scanExercises(root)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(dest, data, 0o644)
}

func resolve(query string, list []entry) (entry, error) {
	q := strings.ToLower(strings.TrimSpace(query))
	var exact []entry
	var hits []entry
	for _, e := range list {
		full := strings.ToLower(e.Section + "/" + e.Name)
		name := strings.ToLower(e.Name)
		if full == q || name == q {
			exact = append(exact, e)
			continue
		}
		if strings.Contains(full, q) || strings.Contains(name, q) {
			hits = append(hits, e)
		}
	}
	if len(exact) == 1 {
		return exact[0], nil
	}
	if len(exact) > 1 {
		var names []string
		for _, h := range exact {
			names = append(names, h.Section+"/"+h.Name)
		}
		return entry{}, fmt.Errorf("ambiguous %q: %s", query, strings.Join(names, ", "))
	}
	if len(hits) == 1 {
		return hits[0], nil
	}
	if len(hits) > 1 {
		var names []string
		for _, h := range hits {
			names = append(names, h.Section+"/"+h.Name)
		}
		return entry{}, fmt.Errorf("ambiguous %q: %s", query, strings.Join(names, ", "))
	}
	return entry{}, fmt.Errorf("unknown exercise %q (try --list)", query)
}

func copyExercise(srcDir, outDir string, e entry) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	for _, f := range e.Files {
		data, err := os.ReadFile(filepath.Join(srcDir, e.Section, e.Name, f))
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(outDir, f), data, 0o644); err != nil {
			return err
		}
	}
	return nil
}

func writeGoMod(outDir, name string) error {
	safe := strings.ToLower(strings.ReplaceAll(name, "-", "_"))
	content := "module fetch/" + safe + "\n\ngo " + goVersion + "\n"
	return os.WriteFile(filepath.Join(outDir, "go.mod"), []byte(content), 0o644)
}

const goVersion = "1.23" // keep in sync with root go.mod

const defaultBase = "https://raw.githubusercontent.com/hamzahraihan/100-exercises-to-learn-go"

func fetchRemote(client *http.Client, baseURL, branch string, e entry, outDir string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	for _, f := range e.Files {
		url := strings.TrimSuffix(baseURL, "/") + "/" + branch + "/exercises/" + e.Section + "/" + e.Name + "/" + f
		resp, err := client.Get(url)
		if err != nil {
			return fmt.Errorf("GET %s: %w (try --local)", url, err)
		}
		body, err := io.ReadAll(resp.Body)
		status := resp.StatusCode
		closeErr := resp.Body.Close()
		if err != nil {
			return fmt.Errorf("GET %s: %w", url, err)
		}
		if closeErr != nil {
			return fmt.Errorf("GET %s: %w", url, closeErr)
		}
		if status != 200 {
			return fmt.Errorf("GET %s: status %d (check --branch, try --local)", url, status)
		}
		if err := os.WriteFile(filepath.Join(outDir, f), body, 0o644); err != nil {
			return err
		}
	}
	return writeGoMod(outDir, e.Name)
}

func loadList() ([]entry, error) {
	return defaultEnv().loadList("main", false)
}

// fetchEnv injects filesystem/network seams so dispatch is testable
// without depending on process CWD.
type fetchEnv struct {
	exRoot       string
	manifestPath string
	client       *http.Client
	baseURL      string
}

func defaultEnv() fetchEnv {
	return fetchEnv{
		exRoot:       "exercises",
		manifestPath: filepath.Join("tools", "fetch", "manifest.json"),
		client:       http.DefaultClient,
		baseURL:      defaultBase,
	}
}

func (f fetchEnv) loadList(branch string, useRemote bool) ([]entry, error) {
	if _, err := os.Stat(f.exRoot); err == nil {
		return scanExercises(f.exRoot)
	}
	if data, err := os.ReadFile(f.manifestPath); err == nil {
		var list []entry
		if err := json.Unmarshal(data, &list); err != nil {
			return nil, err
		}
		return list, nil
	}
	if !useRemote {
		return nil, fmt.Errorf("no ./%s dir and no %s", f.exRoot, f.manifestPath)
	}
	return fetchManifestRemote(f.client, f.baseURL, branch)
}

func (f fetchEnv) materialize(e entry, outDir, branch string, useRemote bool) error {
	if useRemote {
		return fetchRemote(f.client, f.baseURL, branch, e, outDir)
	}
	if _, err := os.Stat(f.exRoot); err != nil {
		return fmt.Errorf("no ./%s dir (try --remote)", f.exRoot)
	}
	if err := copyExercise(f.exRoot, outDir, e); err != nil {
		return err
	}
	return writeGoMod(outDir, e.Name)
}

func isEmptyDir(dir string) (bool, error) {
	ents, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return len(ents) == 0, nil
}

func main() {
	if len(os.Args) == 3 && os.Args[1] == "--dump-manifest" {
		if err := dumpManifest("exercises", os.Args[2]); err != nil {
			fmt.Fprintln(os.Stderr, "fetch:", err)
			os.Exit(1)
		}
		return
	}
	if err := runLocal(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "fetch:", err)
		os.Exit(1)
	}
}

type fetchOpts struct {
	query, out, branch    string
	force, listOnly       bool
	wantRemote, wantLocal bool
}

func parseArgs(args []string) (fetchOpts, error) {
	var o fetchOpts
	o.branch = "main"
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "--list":
			o.listOnly = true
		case a == "--force":
			o.force = true
		case a == "--local":
			o.wantLocal = true
		case a == "--remote":
			o.wantRemote = true
		case a == "--out" && i+1 < len(args):
			i++
			o.out = args[i]
		case strings.HasPrefix(a, "--out="):
			o.out = strings.TrimPrefix(a, "--out=")
		case a == "--branch" && i+1 < len(args):
			i++
			o.branch = args[i]
		case strings.HasPrefix(a, "--branch="):
			o.branch = strings.TrimPrefix(a, "--branch=")
		case strings.HasPrefix(a, "--"):
			return o, fmt.Errorf("unknown flag %q", a)
		default:
			if o.query == "" {
				o.query = a
			} else {
				return o, fmt.Errorf("too many arguments")
			}
		}
	}
	return o, nil
}

func runLocal(args []string) error {
	o, err := parseArgs(args)
	if err != nil {
		return err
	}
	return runWithEnv(o, defaultEnv())
}

func runWithEnv(o fetchOpts, f fetchEnv) error {
	useRemote := o.wantRemote || (!o.wantLocal && isMissingDir(f.exRoot))
	list, err := f.loadList(o.branch, useRemote)
	if err != nil {
		return err
	}
	if o.listOnly {
		for _, e := range list {
			fmt.Println(e.Section + "/" + e.Name)
		}
		return nil
	}
	if o.query == "" {
		return fmt.Errorf("usage: fetch <query> [--out DIR] [--force] [--list]")
	}
	e, err := resolve(o.query, list)
	if err != nil {
		return err
	}
	out := o.out
	if out == "" {
		out = "./" + e.Name
	}
	empty, err := isEmptyDir(out)
	if err != nil {
		return err
	}
	if !empty && !o.force {
		return fmt.Errorf("%s not empty (use --force)", out)
	}
	return f.materialize(e, out, o.branch, useRemote)
}

func isMissingDir(p string) bool {
	fi, err := os.Stat(p)
	return err != nil || !fi.IsDir()
}

func fetchManifestRemote(client *http.Client, baseURL, branch string) ([]entry, error) {
	url := strings.TrimSuffix(baseURL, "/") + "/" + branch + "/tools/fetch/manifest.json"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GET %s: status %d (check --branch)", url, resp.StatusCode)
	}
	var list []entry
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}
	return list, nil
}
