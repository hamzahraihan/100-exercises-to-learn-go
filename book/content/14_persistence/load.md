---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Loading State at Boot"
weight: 2
draft: false
---

# Loading State at Boot

Servers usually reload their state from disk when they start, but on the
very first run there is no file yet — and that is normal, not a crash. The
convention is to treat a missing file as empty state: return `(nil, nil)`
and let the caller start fresh. Any other error (permissions, a directory
in the way) is still a real error. Go reports a missing file through
`os.IsNotExist`, which recognizes the "not exist" error even when it is
wrapped.

## Task

Fill in `Load` in `load.go`:

```go
func Load(path string) ([]byte, error) {
	// ...
}
```

Read the file with `os.ReadFile`, but map "file does not exist" to
`(nil, nil)` using `os.IsNotExist`. The stub returns the raw error, so the
test fails on the missing-file case.

## Check

```bash
go test ./exercises/14_persistence/02_load/ -v
```

---

*Source: `exercises/14_persistence/02_load/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
