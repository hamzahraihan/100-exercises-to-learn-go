---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "The test workflow"
weight: 2
draft: false
---

# The test workflow

Go tests are table-driven: a slice of inputs and wants, looped with `t.Errorf`
on mismatch. You run one exercise with `go test <path> -v`.

## Task

Fix `Add` in `add.go`.

## Check

```bash
go test ./exercises/01_intro/02_test_workflow/ -v
```

---

*Source: `exercises/01_intro/02_test_workflow/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
