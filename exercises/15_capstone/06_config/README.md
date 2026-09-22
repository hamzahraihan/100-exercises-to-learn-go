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
