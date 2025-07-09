# `E2S3 - Passages`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### Binary Structures
So far, all the work we've accomplished has been entirely focused around the `Phrase` structure - which is
just a high-level way of dynamically managing an _unbounded_ run of bits - but here in _reality_ we don't
typically like to perform serial operations on infinite stretches of information!  

_We're smarter than that_ - thanks, in no small part, to Dr. Gene Amdahl =)

The process of binary synthesis is what the industry calls an _embarrassingly parallel_ operation - meaning
that _subdividing_ the process into what I call _passages_ and then logically managing them has _incredible_ 
performance implications!  Why?  Well, let's consider the abstract steps we are performing so far -

    0 - Read in the data

    1 - Synthesize the midpoint
    2 - Subtract the midpoint from the target

    3 - Repeat

The first point is a one-shot