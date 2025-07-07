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

Let's say you _know_ the logical data exists in a 6-bit index, as it takes six bits to represent it.  From 
there you synthesize the _midpoint_ of that bit range and calculate the _distance_ to the target:

    "The Midpoint Operation"

        |←   6 Bits  →|
        | 1 0 0 0 0 0 | (32)  ← Midpoint
        | 0 1 0 1 1 0 | (22)  ← Target
        |             |
        |   - 1 0 1 0 | (-10) ← Distance

Immediately, we have gained a _single_ bit of reduction!  Of course, we'd need to _implicitly_ know
the original index width to recreate the target value, but we'll get to that shortly.  Since we don't
yet have a mechanism for managing the sign, we consider that it is a *transient* component and represent
it by a negative sign.  This particular operation will yield one bit of reduction for all indexes
_at minimum!_  (Assuming its wide enough to be midpointed)

Or, to put it more formally:

    "The Index Midpointing Priciple"

        A logical index point can be rewritten as the distance from the midpoint of its 
        containing index in at least one bit less than counting from 0.

### Why?
Well, because binary is literally repeating the same values twice before growing by one bit -

    |←     8 Bits    →|
    | 1 1 1 1 1 1 1 1 | (255) ← The dark side
    | 1 0 0 0 0 0 0 0 | (128) ← The midpoint
    |   1 1 1 1 1 1 1 | (127) ← A next smaller dark side
    | 0 0 0 0 0 0 0 0 |   (0) ← The light side

      0 + 127 = 127 ← The lower address range
    128 + 127 = 255 ← The upper address range

Now, when one takes the _numeric form_ of a binary value, the growth rate of bits follows an _exponential_
curvature. If you count from _zero,_ it plateaus after the midpoint - meaning any _**high** address_ information 
could never shrink in bit length **_unless_** you count from the midpoint!

      Traditional Counting          Midpoint Counting
        [ 1 1 1 1 ] (15)            [ 1 1 1 ] (15) = 8 + 7
        [ 1 1 1 0 ] (14)            [ 1 1 0 ] (14) = 8 + 6
        [ 1 1 0 1 ] (13)            [ 1 0 1 ] (13) = 8 + 5
        [ 1 1 0 0 ] (12)            [ 1 0 0 ] (12) = 8 + 4
        [ 1 0 1 1 ] (11)              [ 1 1 ] (11) = 8 + 3
        [ 1 0 1 0 ] (10)              [ 1 0 ] (10) = 8 + 2
        [ 1 0 0 1 ]  (9)                [ 1 ]  (9) = 8 + 1
        [ 1 0 0 0 ]  (8) ←  Midpoint  → [ 0 ]  (8) = 8 + 0
          [ 1 1 1 ]  (7)                [ 1 ]  (7) = 8 - 1
          [ 1 1 0 ]  (6)              [ 1 0 ]  (6) = 8 - 2
          [ 1 0 1 ]  (5)              [ 1 1 ]  (5) = 8 - 3
          [ 1 0 0 ]  (4)            [ 1 0 0 ]  (4) = 8 - 4
            [ 1 1 ]  (3)            [ 1 0 1 ]  (3) = 8 - 5
            [ 1 0 ]  (2)            [ 1 1 0 ]  (2) = 8 - 6
              [ 1 ]  (1)            [ 1 1 1 ]  (1) = 8 - 7
              [ 0 ]  (0)          [ 1 0 0 0 ]  (0) = 8 - 8

Essentially, we've _reflected_ the exponential curvature of potential bit drop across the midpoint.  As most
binary information is _extremely grey,_ meaning it's often in the middle of the index, this action dramatically
increases our chances of getting a bit drop from representing the value in its numeric form!  

Let's formalize this particular step, as it's the lynch pin that holds everything together -

    "The Index Midpoint Delta Operation"
    
    𝑡 = The target value to encode
    𝑚 = The index midpoint

    Δ = 𝑡 - 𝑚  

Let's say you wish to encode the logical point `0101` in a nibble index using delta encoding -

    [ 1 0 0 0 ]  (8) ← The midpoint
    [ 0 1 0 1 ]  (5) ← The target point

    Δ = 5 - 8 = -3

      [ - 1 1 ]  (3) ← The delta

While this is pretty much the gist of delta encoding, this is _not_ enough information to reconstruct the original
data!  Most importantly - what do you do with the sign!?  The next solution tackles that through implicit recursion =)

### The Midpoint Averager
This particular solution calculates the _average_ bit drop of many cycles of midpointing grey data.