---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Sort and Filter"
weight: 4
draft: false
---

# Sort and Filter

Ordering and selecting are the two everyday slice transforms. For ordering,
sort a copy rather than the caller's slice when the contract says the input
stays untouched — the standard library offers `slices.SortFunc` and
`sort.Slice` by name for custom orderings such as ascending `ID`. For
selecting, the idiom is filter-by-append: walk the input, append each match
into a fresh result slice, and return that result. Neither transform should
mutate or alias the input's backing array in surprising ways. This is the Go
equivalent of the ordering-and-filtering step in the ticket-store section of
the Rust course this section is adapted from.

## Task

Complete `SortedByID` and `OpenOnly` in `store.go` so one returns tickets
ordered by `ID` ascending and the other keeps only tickets that are not
closed, and the tests pass:

```go
func SortedByID(ts []Ticket) []Ticket {
	// ...
}

func OpenOnly(ts []Ticket) []Ticket {
	// ...
}
```

## Check

```bash
go test ./exercises/06_ticket_store/04_sort_filter/ -v
```

---

*Source: `exercises/06_ticket_store/04_sort_filter/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
