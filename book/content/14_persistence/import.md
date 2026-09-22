---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Importing Backups Safely"
weight: 5
draft: false
---

# Importing Backups Safely

Import is where bad data tries to get back in: a backup may be truncated,
hand-edited, or from an older version. So every import validates before it
trusts — here, each ticket must have a non-empty title, and the whole
import fails if any ticket does not. When validation passes, this course
uses replace semantics: the imported tickets become the entire state,
rather than merging with whatever is already loaded. Replace is simpler to
reason about (no duplicate or ordering questions); merge is a later,
deliberate feature.

## Task

Fill in `Import` in `importtk.go`:

```go
func Import(path string) ([]Ticket, error) {
	// ...
}
```

Read the file with `os.ReadFile`, decode it with `json.Unmarshal`, and
reject any ticket with an empty title. The stub returns `(nil, nil)`, so
the test fails on the good file.

## Check

```bash
go test ./exercises/14_persistence/05_import/ -v
```

---

*Source: `exercises/14_persistence/05_import/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
