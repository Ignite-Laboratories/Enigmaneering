# `E2S3 - Breadcrumbs`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### "Here's Your Sign"
At this point we get to accept that _sometimes the data will get bigger,_ but on average it will _reduce
in length._  One bit of reduction is _not_ anything to write home about, but the value itself becomes
a new binary value which can _also_ be midpointed.  In fact, every time you apply the midpoint operation
you simply are storing how much the value is away from an _implicitly_ reproducible point one bit 
shorter than the last.

But there's an issue - the _sign!_

**_Where do we put it!?_**

Well, let's first consider that we have the _entire width_ of the initial index to describe a **path** to the
target value - and the faster we do so, the smaller the data gets with each cycle.  That means we can leave
behind notes to ourselves of what we did to affect the value, and in this case we can leave behind a single
bit representing the sign at each step!  This means the delta is _always_ stored as a positive integer which 
_fluctuates_ in length while the pathway's length grows _linearly_ with each stored sign.

For example, let's encode the number 42 in an 8 bit index:

    Step 0
    [ 1 0 0 0 0 0 0 0 ] (128) 🡨 Midpoint
    [ 0 0 1 0 1 0 1 0 ]  (42) 🡨 Target
      [ 1 0 1 0 1 1 0 ]  (86) 🡨 Delta
                  [ 1 ] ⟵────── Pathway

    Step 1
      [ 1 0 0 0 0 0 0 ]  (64) 🡨 Midpoint
      [ 1 0 1 0 1 1 0 ]  (86) 🡨 Target
          [ 1 0 1 1 0 ]  (22) 🡨 Delta
                [ 1 0 ] 🡨────── Pathway

    Step 2
        [ 1 0 0 0 0 0 ]  (32) 🡨 Midpoint
          [ 1 0 1 1 0 ]  (22) 🡨 Target
            [ 1 0 1 0 ]  (10) 🡨 Delta
              [ 1 0 1 ] 🡨────── Pathway

    Step 3
          [ 1 0 0 0 0 ]  (16) 🡨 Midpoint
            [ 1 0 1 0 ]  (10) 🡨 Target
              [ 1 1 0 ]   (6) 🡨 Delta
            [ 1 0 1 1 ] 🡨────── Pathway

    Step 4
            [ 1 0 0 0 ]   (8) 🡨 Midpoint
              [ 1 1 0 ]   (6) 🡨 Target
                [ 1 0 ]   (2) 🡨 Delta
          [ 1 0 1 1 1 ] 🡨────── Pathway

    Step 5
              [ 1 0 0 ]   (4) 🡨 Midpoint
                [ 1 0 ]   (2) 🡨 Target
                [ 1 0 ]   (2) 🡨 Delta
        [ 1 0 1 1 1 1 ] 🡨────── Pathway

    Step 6
                [ 1 0 ]   (2) 🡨 Midpoint
                [ 1 0 ]   (2) 🡨 Target
                  [ 0 ]   (0) 🡨 Delta
      [ 1 0 1 1 1 1 0 ] 🡨────── Pathway

Now, this is what we are left with:

        ⬐ The Pathway   ⬐ The Remainder
    [ 1 0 1 1 1 1 0 ] [ 0 ]

This value will _always_ be the same bit length as the original information if you take it _all the way across the
index._ In reality, many values reach a '0' delta value long before the end!  For now please take notice of the 
behavior of the delta value's _bit length_.  Not only does the value _dramatically_ drop in a couple of steps, the 
bit length often yields an overall reduction when you include the pathway's bits.

    Delta Value | Bit Length | Overall Length
         86     |     7      |       8
         22     |     5      |       7
         10     |     4      |       7
          6     |     3      |       7
          2     |     2      |       7
          2     |     2      |       8
          0     |     1      |       8

More importantly, it reaches a minimal bit-length delta _several midpoint operations from the end!_  From my
experience I posit that this condition arises *on average* three steps from the end of recursively midpointing 
**any** size index.

    "The Recursive Binary Midpointing Conjecture"

        Recursively midpointing a binary value will reach a 1-bit delta
        on average three steps from the end of its containing index.

To demonstrate this I've provided two examples:

 **The Breadcrumb Printer** - This simply prints out the breadcrumb path of a randomly generated binary value.
But, and more importantly, it highlights the wave-like chaos of the delta's variable bit length =)

 **The 1-Bit Averager** - This averages the number of positions from the end of the data that it took to reach a 
 1-bit value.  
 
Neither of these exhaustively test the above conjecture, but they provide evidence for _why_ this algorithm works.
The reason for testing the average depth where we reach a 1-bit value is because it affords us a consistent
point at which to _stop_ midpointing, since the delta would (on average) be a 1-bit value at that point.  If we can
consistently bailout of the recursion _before_ the end of the index, we can actually start to _shrink_ the binary 
information!

In the next solution we will begin exploring the concept of accomplishing this by using a _**passage**._