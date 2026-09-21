# The test workflow

Go tests are table-driven: a slice of inputs and wants, looped with `t.Errorf`
on mismatch. You run one exercise with `go test <path> -v`.

## Task

Fix `Add` in `add.go`.

## Check

```bash
go test ./exercises/01_intro/02_test_workflow/ -v
```
