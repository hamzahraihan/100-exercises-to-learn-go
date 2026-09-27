---
title: "JSON"
weight: 11
draft: false
---

# JSON

The wire format. Structs become bytes and bytes become structs again, with
every policy decision the trip requires made explicit: tags that rename,
hide, and omit; validation after decoding because the wire is trusted by
no one; strictness as a per-consumer choice; and custom marshalers for
types whose Go shape isn't their JSON shape. The HTTP section assumes all
of it.
