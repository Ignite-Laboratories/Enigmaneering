# `E1S2 - Delta Encoding`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### The Beginnings of "Middle-Out Compression"
Let's circle back to the _numeric representation_ of a binary value again:

      ⬐ Full bit width   ⬐ Equivalent numeric value
    [ 0 0 1 0 1 0 1 0 ] (42) 🡨 Logical form
        [ 1 0 1 0 1 0 ] (42) 🡨 Numeric form
              ⬑ Truncated bit width

Alone, this number is _absolutely no different_ from any other binary value!  But if you applied it to
a different value than _zero_ you'd get an entirely different result.

Let's say you _know_ the data exists in an 6-bit index, as it takes six bits to represent it.  From there
you synthesize the _midpoint_ of that bit range, then take the _delta_ between the target the midpoint.
Finally, you keep the _signed_ result:

    "The Midpoint Operation"

        [ 1 1 1 1 1 1 ] (63)  🡨 Dark Boundary
        [ 1 0 1 0 1 0 ] (42)  🡨 Target Value
        [ 1 0 0 0 0 0 ] (32)  🡨 Midpoint
    
          [ - 1 0 1 0 ] (-10) 🡨 Delta
    
        [ 0 0 0 0 0 0 ] (0)   🡨 Light Boundary

Immediately, we have gained a _single_ bit of reduction!  Of course, we'd need to _implicitly_ know
the width of the data to synthesize a midpoint value from, but we'll get to that shortly.  Since we don't
yet have a mechanism for managing the sign, we consider that it is a *transient* component and represent
it by a negative sign.  This particular operation will yield one bit of reduction for _all_ sizes of binary 
information _on average!_  Or, to put it more formally:

    "The Binary Midpointing Priciple"

        Binary information can be rewritten as a delta from the midpoint of its 
        containing index in an average of one bit less than counting from 0.

This solution provides a simple way to _test_ that the average bit drop is, indeed, one bit for the ranges
relevant to our purposes.  If you'd like to help me formulate a proper mathematical proof, I'd be thrilled!
For now, I'm happy to simply _demonstrate_ the techology for others.

This is _not_ enough information to reconstruct the data, however - how do you _implicitly_ know what sign 
the resulting delta has?  If only we could consider the sign as an entirely separate _artifact_ and store the
**_absolute value_** of the delta, then we'd _technically_ be achieving an average bit drop of _'2'..._

The next solution tackles just that =)