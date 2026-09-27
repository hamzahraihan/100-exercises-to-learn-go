---
title: "Concurrency"
weight: 7
draft: false
---

# Concurrency

Do things at once, safely. The section walks Go's famous progression in
order: shared state guarded by mutexes, then the inversion — sharing by
communicating through channels — then buffered decoupling, then `select`
waiting on whichever future answers first. Four lessons, each with the
race detector as co-pilot, ending where real servers live.
