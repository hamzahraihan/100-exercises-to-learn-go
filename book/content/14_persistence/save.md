---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Crash-Safe Saves"
weight: 1
draft: false
---

# Crash-Safe Saves

Writing straight to the final path is risky: if the process crashes
mid-write, the file is left half-written and the old data is gone. The fix
is to write to a temporary file first and then rename it over the target.
On most filesystems a rename within the same directory is atomic — readers
see either the complete old file or the complete new file, never a mix.
The temp file must live in the same directory as the target, because only
same-directory renames are guaranteed atomic (a cross-directory rename may
silently become a slow, non-atomic copy).

## Task

Fill in `Save` in `atomic.go`:

```go
func Save(path string, data []byte) error {
	// ...
}
```

Create a temp file in the target's directory with `os.CreateTemp`, write
`data` into it, `Close` it, then `os.Rename` it onto `path`. The stub
writes nothing, so the test fails when the file is missing.

## Check

```bash
go test ./exercises/14_persistence/01_save/ -v
```

---

*Source: `exercises/14_persistence/01_save/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
