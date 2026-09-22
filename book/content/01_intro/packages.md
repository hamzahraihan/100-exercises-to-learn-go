---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Packages and imports"
weight: 3
draft: false
---

# Packages and imports

Go code lives in packages. `import "strings"` makes the standard library's
string helpers available, and only capitalized names (`ToUpper`) are visible
outside their package.

## Task

Fix `Shout` in `greet.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/03_packages/ -v
```

---

*Source: `exercises/01_intro/03_packages/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
