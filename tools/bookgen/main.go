// tools/bookgen/main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type exercise struct {
	Section string // e.g. 01_intro
	Name    string // e.g. 01_syntax
	Title   string // from first H1
	Body    string // raw README body
}

func parseWeight(dir string) (int, error) {
	idx := strings.Index(dir, "_")
	if idx < 0 {
		return 0, fmt.Errorf("no numeric prefix in %q", dir)
	}
	n, err := strconv.Atoi(dir[:idx])
	if err != nil {
		return 0, fmt.Errorf("bad numeric prefix in %q: %w", dir, err)
	}
	return n, nil
}

func extractTitle(body string) (string, error) {
	for _, line := range strings.Split(body, "\n") {
		if strings.HasPrefix(line, "# ") {
			title := strings.TrimSpace(strings.TrimPrefix(line, "# "))
			if title == "" {
				break
			}
			return title, nil
		}
	}
	return "", fmt.Errorf("no '# Title' heading found")
}

func walkExercises(root string) ([]exercise, error) {
	var out []exercise
	seen := map[string]map[int]string{} // section -> weight -> name
	sections, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, s := range sections {
		if !s.IsDir() {
			continue
		}
		entries, err := os.ReadDir(filepath.Join(root, s.Name()))
		if err != nil {
			return nil, err
		}
		for _, e := range entries {
			if !e.IsDir() {
				continue
			}
			readme := filepath.Join(root, s.Name(), e.Name(), "README.md")
			raw, err := os.ReadFile(readme)
			if err != nil {
				return nil, fmt.Errorf("missing %s: %w", readme, err)
			}
			body := string(raw)
			title, err := extractTitle(body)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", readme, err)
			}
			w, err := parseWeight(e.Name())
			if err != nil {
				return nil, fmt.Errorf("%s: %w", readme, err)
			}
			if seen[s.Name()] == nil {
				seen[s.Name()] = map[int]string{}
			}
			if prev, dup := seen[s.Name()][w]; dup {
				return nil, fmt.Errorf("duplicate weight %d in %s: %s vs %s", w, s.Name(), prev, e.Name())
			}
			seen[s.Name()][w] = e.Name()
			out = append(out, exercise{Section: s.Name(), Name: e.Name(), Title: title, Body: body})
		}
	}
	return out, nil
}

func main() {
	exRoot, outRoot := "exercises", filepath.Join("book", "content")
	if len(os.Args) > 1 {
		exRoot = os.Args[1]
	}
	if len(os.Args) > 2 {
		outRoot = os.Args[2]
	}
	if err := run(exRoot, outRoot); err != nil {
		fmt.Fprintln(os.Stderr, "bookgen:", err)
		os.Exit(1)
	}
}

func quoteYAML(s string) string {
	r := strings.ReplaceAll(s, `\`, `\\`)
	r = strings.ReplaceAll(r, `"`, `\"`)
	return `"` + r + `"`
}

func emit(ex exercise) string {
	body := strings.ReplaceAll(ex.Body, "\r\n", "\n")
	w, _ := parseWeight(ex.Name)
	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.\n")
	fmt.Fprintf(&b, "title: %s\n", quoteYAML(ex.Title))
	fmt.Fprintf(&b, "weight: %d\n", w)
	b.WriteString("draft: false\n---\n\n")
	b.WriteString(body)
	if !strings.HasSuffix(body, "\n") {
		b.WriteString("\n")
	}
	b.WriteString("\n---\n\n")
	fmt.Fprintf(&b, "*Source: `exercises/%s/%s/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*\n", ex.Section, ex.Name)
	return b.String()
}

func run(exRoot, outRoot string) error {
	list, err := walkExercises(exRoot)
	if err != nil {
		return err
	}
	for _, ex := range list {
		dir := filepath.Join(outRoot, ex.Section)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		dest := filepath.Join(dir, strings.TrimPrefix(ex.Name, digitsPrefix(ex.Name))+".md")
		content := emit(ex)
		old, err := os.ReadFile(dest)
		if err == nil && string(old) == content {
			continue
		}
		if err := os.WriteFile(dest, []byte(content), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func digitsPrefix(name string) string {
	i := 0
	for i < len(name) && name[i] >= '0' && name[i] <= '9' {
		i++
	}
	if i < len(name) && name[i] == '_' {
		return name[:i+1]
	}
	return ""
}
