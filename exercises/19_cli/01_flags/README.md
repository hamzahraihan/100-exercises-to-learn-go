# Command-Line Flags

Every real Go binary needs a way to accept options — a port number, a file path, a debug toggle. The `flag` package is stdlib's answer: declare flags, call `Parse`, read values. No dependencies, no magic, and it generates `-help` output for free.

## Declaring and parsing

```go
port := flag.Int("port", 8080, "TCP port to listen on")
debug := flag.Bool("debug", false, "enable verbose logging")
flag.Parse()
fmt.Println(*port, *debug)
```

`flag.Int` registers a flag named `"port"` with default `8080`; it returns a `*int` that is filled after `Parse`. The third argument is the usage string shown by `-help`. `flag.Parse` reads `os.Args[1:]` and populates all registered flags.

## FlagSet for testability

The global `flag` package functions touch `os.Args` and call `os.Exit` on error — untestable. Use a `*flag.FlagSet` instead:

```go
fs := flag.NewFlagSet("myapp", flag.ContinueOnError)
port := fs.Int("port", 8080, "TCP port")
fs.Parse(args) // args is []string, not os.Args
fmt.Println(*port)
```

`flag.ContinueOnError` makes `Parse` return an error instead of exiting — essential for tests. Wrap this in a `ParseConfig(args []string) (Config, error)` function and it becomes trivially testable.

## Positional arguments

```go
fs.Parse(args)
rest := fs.Args() // non-flag arguments after the flags
```

`fs.Args()` returns everything after the last flag. Convention: flags first, positional arguments after `--` or after all flags are consumed.

## Usage customization

```go
fs.Usage = func() {
    fmt.Fprintf(fs.Output(), "usage: myapp [flags] <file>\n")
    fs.PrintDefaults()
}
```

Override `fs.Usage` to control the `-help` message; `fs.PrintDefaults()` prints each flag's default and usage string.

<details><summary>Hint</summary>

In `ParseConfig`, create a `flag.NewFlagSet("server", flag.ContinueOnError)`. Register `port` (int, default 8080), `debug` (bool, default false), and `data` (string, default `"data.json"`). Call `fs.Parse(args)` and return `Config{Port: *port, Debug: *debug, DataFile: *data}` or the parse error.

</details>

## Task

Complete `ParseConfig(args []string) (Config, error)` in `flags.go`. The `Config` struct is provided.

## Check

```bash
go test ./exercises/19_cli/01_flags/ -v
```
