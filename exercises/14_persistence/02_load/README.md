# Loading State at Boot

Servers restart. On every boot after the first, the state file waits on
disk, ready to reload. But the *first* boot has no file — nothing was ever
saved — and that absence is the normal case, not a disaster. Code that
treats it as an error crashes fresh installs on day one. This exercise
teaches the loader to welcome the void.

## Absence is not failure

```go
func Load(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, nil // first boot: empty state, no error
        }
        return nil, err
    }
    return data, nil
}
```

Three outcomes, three answers. File present: contents, nil error.
File missing: `(nil, nil)` — empty state, success. Anything else
(permissions, a directory in the way, I/O faults): the real error,
propagated untouched. Callers start fresh on `(nil, nil)` and abort on
anything else, and the distinction lives in exactly one place instead of
every call site re-deriving it.

## Recognizing "not exist" through wrapping

```go
os.IsNotExist(err) // true even when the error wears wrappers
```

Filesystem errors travel wrapped — path context attached at every layer,
the corrupt-file lesson's `%w` chains in the wild. Comparing with `==`
against a sentinel would miss them all. `os.IsNotExist` (equivalently,
`errors.Is(err, fs.ErrNotExist)`) descends the chain asking each layer
"are you, underneath, a missing file?" Use the predicate, never the
comparison: identity through wrapping is the entire `errors.Is` doctrine,
and boot loaders are where it pays rent most visibly.

But reach for it narrowly. Mapping *every* error to empty state would
swallow permission failures and disk faults into silent fresh starts —
data loss wearing a clean boot's face. Only "not exist" converts; the rest
propagate. Absence is normal. Everything else is news.

## The test's two doors

```go
Load(path)                            // existing file → ("hi", nil)
Load(filepath.Join(dir, "missing.json")) // absent file → (nil, nil)
```

Present and missing, asserted side by side — the two-branch habit from the
error-assertion lesson, now with the subtle branch: success carrying
*nothing*. `(nil, nil)` looks like a non-answer, but it's the most
informative return in the function: "nothing saved yet, proceed
accordingly." Callers that can't distinguish it from failure will crash
day-one installs; callers that can will boot anywhere.

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
