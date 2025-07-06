# `E2S2 - Delta Encoding`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### "Middle-Out Compression"
Let's circle back to the _numeric representation_ of a binary value again:

      ⬐ Full bit width   ⬐ Equivalent numeric value
    [ 0 0 1 0 1 0 1 0 ] (42) ← Logical form
        [ 1 0 1 0 1 0 ] (42) ← Numeric form
              ⬑ Truncated bit width

Alone, this number is _absolutely no different_ from any other binary value!  But if you implicitly
applied it to any other value than _zero_ you'd get an entirely different result.

Let's say you _know_ the data exists in a 6-bit index, as it takes six bits to represent it.  From there
you synthesize the _midpoint_ of that bit range and calculate the _distance_ to the target:

    "The Midpoint Operation"

        |←   6 Bits  →|
        [ 1 0 0 0 0 0 ] (32)  ← Midpoint
          [ 1 0 1 1 0 ] (22)  ← Target Value
    
          [ - 1 0 1 0 ] (-10) ← Distance

Immediately, we have gained a _single_ bit of reduction!  Of course, we'd need to _implicitly_ know
the original index width to recreate the target value, but we'll get to that shortly.  Since we don't
yet have a mechanism for managing the sign, we consider that it is a *transient* component and represent
it by a negative sign.  This particular operation will yield one bit of reduction for all indexes
_at minimum!_  (Assuming its wide enough to be midpointed)

Or, to put it more formally:

    "The Binary Midpointing Priciple"

        A binary value can be rewritten as the distance from the midpoint of its 
        containing index in an average of at least one bit less than counting from 0.

### Why?
Well, because binary is literally repeating the same values twice before growing by one bit -

    |←     8 Bits    →|
    [ 1 1 1 1 1 1 1 1 ] (255) ← The dark side
    [ 1 0 0 0 0 0 0 0 ] (128) ← The midpoint
      [ 1 1 1 1 1 1 1 ] (127) ← A dark 7-bit point
    [ 0 0 0 0 0 0 0 0 ]   (0) ← The light side

      0 + 127 = 128 ← The lower address range
    128 + 127 = 255 ← The upper address range

Now, when one takes the _numeric form_ of a binary value, the growth rate of bits follows an _exponential_
curvature. If you count from _zero,_ it plateaus after the midpoint - meaning any _**high** address_ information 
could never shrink in bit length **_unless_** you count from the midpoint!

      Traditional Counting          Midpoint Counting
        [ 1 1 1 1 ] (15)            [ 1 1 1 ] (15) +7
        [ 1 1 1 0 ] (14)            [ 1 1 0 ] (14) +6
        [ 1 1 0 1 ] (13)            [ 1 0 1 ] (13) +5
        [ 1 1 0 0 ] (12)            [ 1 0 0 ] (12) +4
        [ 1 0 1 1 ] (11)              [ 1 1 ] (11) +3
        [ 1 0 1 0 ] (10)              [ 1 0 ] (10) +2
        [ 1 0 0 1 ]  (9)                [ 1 ]  (9) +1
        [ 1 0 0 0 ]  (8) ←  Midpoint  → [ 0 ]  (8) +0
          [ 1 1 1 ]  (7)                [ 1 ]  (7) -1
          [ 1 1 0 ]  (6)              [ 1 0 ]  (6) -2
          [ 1 0 1 ]  (5)              [ 1 1 ]  (5) -3
          [ 1 0 0 ]  (4)            [ 1 0 0 ]  (4) -4
            [ 1 1 ]  (3)            [ 1 0 1 ]  (3) -5
            [ 1 0 ]  (2)            [ 1 1 0 ]  (2) -6
              [ 1 ]  (1)            [ 1 1 1 ]  (1) -7
              [ 0 ]  (0)          [ 1 0 0 0 ]  (0) -8

Essentially, we've _reflected_ the exponential curvature of potential bit drop across the midpoint.  As most
binary information is _extremely grey,_ this action increases our chances of getting a potential bit drop 
from representing the value in it's numeric form _dramatically!_  The _official_ midpoint formula is very simple, but
intentionally yields the _inverse_ of the above midpoint counting scheme.  I promise that will make _absolute_
sense in the next solution -

    "The Midpoint Delta Operation"
    
    𝑚 = The index midpoint
    𝑡 = The target value to encode

    Δ = 𝑚 - 𝑡 

Let's say you wish to encode the value `5` in a nibble index using delta encoding -

    [ 1 0 0 0 ] (8) ← The midpoint
    [ 0 1 0 1 ] (5) ← The value to encode

    Δ = 8 - 5 = 3

        [ 1 1 ] (3) ← The absolute delta

While this is pretty much the gist of _delta encoding,_ this is _not_ enough information to reconstruct the original
data!  How do you _implicitly_ know what sign the resulting delta has? If only we could consider the sign as an 
entirely separate _artifact_ and just store the **_absolute value_** of the delta...

We'll tackle that in the next solution =)