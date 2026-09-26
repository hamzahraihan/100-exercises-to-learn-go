# Config From the Environment

One binary, many worlds: laptop, CI, staging, production — different
ports, different data files, same compiled artifact. Hardcoding those
values forks the binary per environment; flags help but must be passed
every launch. The environment is the middle path: set once per machine,
inherited by every process, invisible to the code's logic.

## Read, default, convert leniently

```go
type Config struct {
    Port     int
    DataFile string
}

func ConfigFromEnv() Config {
    port := 8080
    if v := os.Getenv("PORT"); v != "" {
        if n, err := strconv.Atoi(v); err == nil {
            port = n
        }
    }
    dataFile := os.Getenv("DATA_FILE")
    if dataFile == "" {
        dataFile = "tickets.json"
    }
    return Config{Port: port, DataFile: dataFile}
}
```

Three habits in one small function. **Defaults for everything**: `8080`
and `tickets.json` make the bare binary runnable with zero setup — a
server that demands configuration before its first run is a server nobody
tries. **Empty means unset**: `os.Getenv` returns `""` for missing
variables, so empty-string checks double as presence checks (the query
lesson's `""`-means-absent policy, now reading the process environment).

And the port conversion **fails soft**: a garbage `PORT` falls back to
`8080` instead of crashing. Startup code faces a harsh asymmetry — crashing
on bad config is sometimes right (fail fast on secrets), but a port is
operational detail, and refusing to serve over a typo'd number trades a
mystery for an outage. Lenient where safe, strict where it matters: the
strict-decoding lesson's policy thinking, applied to the process boundary.

## Tests set their own weather

```go
t.Setenv("PORT", "9090")
t.Setenv("DATA_FILE", "/tmp/t.json")
```

`t.Setenv` sets a variable *and registers its restoration* — automatic
cleanup when the test ends, no defer needed, parallel-safe where manual
`os.Setenv` would leak across tests. The assertion then reads the struct
whole (`%+v` on mismatch), pinning both fields through the env-to-struct
path. Environment-dependent code stays testable because the testing
package controls the environment — the fake-clock lesson's seam, wearing
operating-system clothes.

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
