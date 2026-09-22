---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Saving Tickets to Disk"
weight: 4
draft: false
---

# Saving Tickets to Disk

The `os` package reads and writes whole files in one call: `os.WriteFile`
takes a path, bytes, and permission bits (like `0o644`, owner-writable and
world-readable), while `os.ReadFile` loads a path back into memory. Tests
that touch the filesystem must never write into the repo — `t.TempDir()`
gives you a fresh scratch directory that Go cleans up afterwards. Here
`LoadTicket` already works, but `SaveTicket` writes nothing, so the test
fails when the load finds no file.

## Task

Fill in `SaveTicket` in `ticketfile.go`:

```go
func SaveTicket(path string, data []byte) error {
	// ...
}
```

Use `os.WriteFile` with `0o644` permissions to persist `data` at `path`.

## Check

```bash
go test ./exercises/12_io/04_file/ -v
```

---

*Source: `exercises/12_io/04_file/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
