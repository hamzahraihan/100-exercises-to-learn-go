// tools/fetch/main.go
package main

import (
	"encoding/json"
	"fmt"
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
	var hits []entry
	for _, e := range list {
		full := strings.ToLower(e.Section + "/" + e.Name)
		name := strings.ToLower(e.Name)
		if full == q || name == q {
			return e, nil
		}
		if strings.Contains(full, q) || strings.Contains(name, q) {
			hits = append(hits, e)
		}
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
	content := "module fetch/" + safe + "\n\ngo 1.23\n"
	return os.WriteFile(filepath.Join(outDir, "go.mod"), []byte(content), 0o644)
}

func loadList() ([]entry, error) {
	if _, err := os.Stat("exercises"); err == nil {
		return scanExercises("exercises")
	}
	data, err := os.ReadFile(filepath.Join("tools", "fetch", "manifest.json"))
	if err != nil {
		return nil, fmt.Errorf("no ./exercises dir and no tools/fetch/manifest.json: %w", err)
	}
	var list []entry
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return list, nil
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

func runLocal(args []string) error {
	var query, out string
	var force, listOnly bool
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "--list":
			listOnly = true
		case a == "--force":
			force = true
		case a == "--local":
			// default; accepted for forward compat with remote mode
		case a == "--remote":
			return fmt.Errorf("remote mode not yet implemented (Task 3)")
		case a == "--out" && i+1 < len(args):
			i++
			out = args[i]
		case strings.HasPrefix(a, "--out="):
			out = strings.TrimPrefix(a, "--out=")
		case strings.HasPrefix(a, "--branch"):
			// accepted in local mode, ignored until Task 3
			if !strings.Contains(a, "=") && i+1 < len(args) {
				i++
			}
		case strings.HasPrefix(a, "--"):
			return fmt.Errorf("unknown flag %q", a)
		default:
			if query == "" {
				query = a
			} else {
				return fmt.Errorf("too many arguments")
			}
		}
	}
	list, err := loadList()
	if err != nil {
		return err
	}
	if listOnly {
		for _, e := range list {
			fmt.Println(e.Section + "/" + e.Name)
		}
		return nil
	}
	if query == "" {
		return fmt.Errorf("usage: fetch <query> [--out DIR] [--force] [--list]")
	}
	e, err := resolve(query, list)
	if err != nil {
		return err
	}
	if out == "" {
		out = "./" + e.Name
	}
	empty, err := isEmptyDir(out)
	if err != nil {
		return err
	}
	if !empty && !force {
		return fmt.Errorf("%s not empty (use --force)", out)
	}
	if _, err := os.Stat("exercises"); err != nil {
		return fmt.Errorf("no ./exercises dir (remote mode lands in Task 3)")
	}
	if err := copyExercise("exercises", out, e); err != nil {
		return err
	}
	return writeGoMod(out, e.Name)
}
