# `E1S1 - Cyclic Refinement`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### Symmetric Reduction
At this point we get to accept that _sometimes the data will get bigger,_ but on average it will _reduce
in length._  One bit of reduction is _not_ anything to write home about, but the value itself becomes
a new binary value which can _also_ be midpointed.  In fact, every time you apply the midpoint operation
you simply are storing off how much the value is away from a _known synthetic point_ that can reliably
be reproduced.  