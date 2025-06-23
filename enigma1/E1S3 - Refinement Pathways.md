# `E1S3 - Refinement Pathways`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### Programmatic Approximation
Up to this point we've had a _single_ operation performed, which yielded intermittently decent approximation
results.  In order to get this project across the line we need to _continue_ the action we just performed, while
leaving a breadcrumb trail behind to repeat the steps we've performed.  In this step, we will start to work with
the `Approximation` structure to further _refine_ a synthesized value.  This structure contains a `Signature` field
that holds the information necessary to recreate the approximation, and we will _emit_ measurements into this phrase
whenever we have found an ideal value.

The signature will hold a cyclic set of bits that details the recursive subdivision process we will perform:

    Signature Phrase Encoding

                       ⬐ Relativity Bit
    [ ⁰⁄₁ ⁰⁄₁ ⁰⁄₁ ] [ ⁰⁄₁ ] ↺
           ⬑ Pattern        ⬑ Repeat

The idea is simple, we start by approximating the target using binary pattern subdivision of the _entire index._
From there, we know the _boundaries_ within which the target resides - providing the next points to subdivide
and find the closest approximation.  The relativity bit indicates if the target exists closer to the top or the
bottom of the defined region, and then the process repeats.  We won't be able to utilize binary pattern subidivision
to refine the value, since that subdivides the entire index, but we can keep the subdivision bit-width the same
for sub-index subdivision.  ( "Sub-index subdivision" - _that's a tongue twister!_ )

This process of "drilling down" through subdivision until you reach the target value eventually yields a final
delta of 0, meaning the micro-loop eventually runs out of bits to process.  This, in turn, yields a signature that 
represents a programmatic pathway of refinement towards a target value which can be replayed if the initial bit-width 
is known (which the next solution tackles).

The last solution is 

This solution creates the refinement pathway for randomly generated input data, while allowing you to mess with the 
bit widths in order to see how it affects the generated pathway.