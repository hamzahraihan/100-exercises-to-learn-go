---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Config From the Environment"
weight: 6
draft: false
---

# Config From the Environment

Servers take settings from the environment so the same binary runs in dev and
prod: `PORT` (default `8080`) and `DATA_FILE` (default `tickets.json`). The
pattern is `os.Getenv`, fall back to the default when empty, and convert the
port with `strconv.Atoi` — falling back to the default when conversion fails,
since a bad `PORT` should not crash the server.

## Task

Fill in `ConfigFromEnv` in `config.go`:

```go
func ConfigFromEnv() Config {
	// ...
}
```

Read both variables with defaults. The stub returns zeros, so the test fails
on the field values.

## Check

```bash
go test ./exercises/15_capstone/06_config/ -v
```

---

*Source: `exercises/15_capstone/06_config/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
