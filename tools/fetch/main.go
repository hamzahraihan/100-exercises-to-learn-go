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

func main() {
	if len(os.Args) == 3 && os.Args[1] == "--dump-manifest" {
		if err := dumpManifest("exercises", os.Args[2]); err != nil {
			fmt.Fprintln(os.Stderr, "fetch:", err)
			os.Exit(1)
		}
	}
}
