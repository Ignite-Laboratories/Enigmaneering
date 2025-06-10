# `E1S2 - Waveform Reflection`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### Binary Signal Analysis
Okay, if you haven't noticed yet the entire concept of binary synthesis comes from _signal analysis._  Right now,
we're working towards defining how to reach a target value using the smallest _bit width_ possible.  In order to
accomplish that, we have a lot of options - from phasing the value through its cycle or to my desired route through
recursion and entropy.

Let's start by showing how to alternate the sub-index loops between incrementing and decrementing to "rectify" our
signal into a more pleasant form =)

This example is the first point where we get to see a _delta_ value.

    delta - the distance to the nearest subdivision point

A delta value could be subtracted from the subdivision point, or it could be added to it - depending on where
the data's value actually resides.