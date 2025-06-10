# `E1S1 - Binary Waveforms`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### Getting Intimate With Binary
The next part to recognize is the difference between numerical and logical binary data.  For synthesis, we work
with the numerical representation of the data while logically managing the leading zeros.  That will make a lot
more sense as we proceed forward - for now, just know that numeric equivalence does _not_ imply logical equivalence.
A single missing zero could entirely destroy the logical structure of data - fun! =)

    [ 0 0 1 0 1 0 1 0 ] [42] <- Logical form
        [ 1 0 1 0 1 0 ] [42] <- Numeric form

However, we're going to _exploit_ that for our purposes.  This solution is quite simple, it prints the numerical binary
values of indices by subdividing the bit range.  This allows you to visualize how binary information follows a ramping
linear wave pattern as you step through indices which plateaus at the midpoint of each sub-index - we aim to smooth this 
progression into a _rectified sine wave_ to maximize compression.

This presents the next colloquial term: the _subdivision._  This represents a point that actively subdivides the address
space.  These are calculated as (2ⁿ/sub-indices) - meaning 0 subdivisions yields one sub index while 3 subdivisions yields
four sub-indices.  Think of a subdivision as the actual splitting point between two binary values, rather than a specific
numeric value - for example:

    A nibble index subdivided three times:

         Dark Side
        [ 1 1 1 1 ] 
        [ 1 1 1 0 ]
             ├-----
        [ 1 1 0 1 ]
        [ 1 1 0 0 ]
           ├------- <- The upper quarter-point
        [ 1 0 1 1 ]
        [ 1 0 1 0 ]
             ├----- <- Eighth-points
        [ 1 0 0 1 ]
        [ 1 0 0 0 ]
         ├--------- <- The mid-point
        [ 0 1 1 1 ]
        [ 0 1 1 0 ]
             ├-----
        [ 0 1 0 1 ]
        [ 0 1 0 0 ]
           ├------- <- The lower quarter-point
        [ 0 0 1 1 ]
        [ 0 0 1 0 ]
             ├----- <- Eighth-points
        [ 0 0 0 1 ]
        [ 0 0 0 0 ]
         Light Side

While the subdivision _point_ is known, how to mathematically consider its value depends on the context of your 
algorithm - for example, do you treat the mid point of 2⁸ as 127 or 128 as, technically, 8 bits can only hold (2⁸)-1?
In my experience, it's easiest to work with 2ⁿ and simply subtract one _whenever necessary._

Feel free to experiment with the input values and witness the output waveforms change - genuinely, understanding
how binary digits grow and shrink is _critical_ to understanding binary synthesis.  For instance, the first thing
you'll likely notice is that they grow in bit-width _exponentially._

The next thing you'll notice is that the waveform is a sawtooth which plateaus at the midpoint of each sub-index,
meaning that only half of each sub-index can even shrink it's bit length when counting from the bottom.

We'll tackle that issue in the next solution.

