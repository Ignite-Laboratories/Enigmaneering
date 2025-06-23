# `E1S1 - Binary Waveforms`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### Getting Intimate With Binary
The next part to recognize is the difference between numerical and logical binary data.  For synthesis, we work
with the numerical representation of data while logically managing the leading zeros.  That will make a lot
more sense as we proceed forward - for now, just know that numeric equivalence does _not_ imply logical equivalence.
A single missing zero could entirely destroy the logical structure of data - fun! =)

      ⬐ Full bit width   ⬐ Equivalent numeric value
    [ 0 0 1 0 1 0 1 0 ] (42) <- Logical form
        [ 1 0 1 0 1 0 ] (42) <- Numeric form
              ⬑ Truncated bit width

Here, the logical form is a _byte_ - but that's not a hard requirement, just the most commonly utilized logical size.

Now, to approximate a value we need to be able to _logarithmically_ subdivide any index of data reliably and with ease.
Luckily there's a binary trick we get to exploit to accomplish this:

**Pattern Subdivision** - Using a repeating binary pattern to subdivide an index of data.

It's a LOT easier to demonstrate this concept than it is to explain, so let's start by subdividing a flake (5-bit)
index of data using a note (3-bit) pattern.  Each pattern value is simply repeated across the five bits, while
stepping incrementally up all of the pattern's values.

    Note Subdivision of a Morsel Index:

    Value ⬎   Pattern    Synthesized   Value    ⬐ Delta   
          (0)[ 0 0 0 ] [ 0 0 0 - 0 0 ] (  0 ) + 4
          (1)[ 0 0 1 ] [ 0 0 1 - 0 0 ] (  4 ) + 5
          (2)[ 0 1 0 ] [ 0 1 0 - 0 1 ] (  9 ) + 4
          (3)[ 0 1 1 ] [ 0 1 1 - 0 1 ] ( 13 ) + 5
          (4)[ 1 0 0 ] [ 1 0 0 - 1 0 ] ( 18 ) + 4
          (5)[ 1 0 1 ] [ 1 0 1 - 1 0 ] ( 22 ) + 5
          (6)[ 1 1 0 ] [ 1 1 0 - 1 1 ] ( 27 ) + 4
          (7)[ 1 1 1 ] [ 1 1 1 - 1 1 ] ( 31 )

The delta value fluctuates because a logarithmic subdivision of this address space is a _non-integer value._
Since binary does not support floating point numbers without encoding schemes, the value _naturally_ fluctuates
around the true logarithmic step size.
I also chose the bit widths I did intentionally as the synthetic value's bit width does _not_ have to be an even
multiple of the pattern's bit width!
Because of the way binary values grow, this kind of logarithmic subdivision works for literally _any bit 
width_ - without fail!  

    Note Subdivision of a 22 bit Index:

        Pattern                         Synthesized                               Value       Delta   
    (0)[ 0 0 0 ] [ 0 0 0 - 0 0 0 - 0 0 0 - 0 0 0 - 0 0 0 - 0 0 0 - 0 0 0 - 0 ] (    0    ) + 599186
    (1)[ 0 0 1 ] [ 0 0 1 - 0 0 1 - 0 0 1 - 0 0 1 - 0 0 1 - 0 0 1 - 0 0 1 - 0 ] (  599186 ) + 599186
    (2)[ 0 1 0 ] [ 0 1 0 - 0 1 0 - 0 1 0 - 0 1 0 - 0 1 0 - 0 1 0 - 0 1 0 - 0 ] ( 1198372 ) + 599186
    (3)[ 0 1 1 ] [ 0 1 1 - 0 1 1 - 0 1 1 - 0 1 1 - 0 1 1 - 0 1 1 - 0 1 1 - 0 ] ( 1797558 ) + 599187
    (4)[ 1 0 0 ] [ 1 0 0 - 1 0 0 - 1 0 0 - 1 0 0 - 1 0 0 - 1 0 0 - 1 0 0 - 1 ] ( 2396745 ) + 599186
    (5)[ 1 0 1 ] [ 1 0 1 - 1 0 1 - 1 0 1 - 1 0 1 - 1 0 1 - 1 0 1 - 1 0 1 - 1 ] ( 2995931 ) + 599186
    (6)[ 1 1 0 ] [ 1 1 0 - 1 1 0 - 1 1 0 - 1 1 0 - 1 1 0 - 1 1 0 - 1 1 0 - 1 ] ( 3595117 ) + 599186
    (7)[ 1 1 1 ] [ 1 1 1 - 1 1 1 - 1 1 1 - 1 1 1 - 1 1 1 - 1 1 1 - 1 1 1 - 1 ] ( 4194303 )

You can also use _any pattern width,_ as long as its shorter than the target's bit width - which coincidentally
yields more subdivision values for the index, directly tied to the bit-width of the pattern:

    Nibble Subdivision of an 11 bit Index:

        Pattern               Synthesized          Value    Delta   
     (0)[ 0 0 0 0 ] [ 0 0 0 0 - 0 0 0 0 - 0 0 0 ] (   0  ) + 136
     (1)[ 0 0 0 1 ] [ 0 0 0 1 - 0 0 0 1 - 0 0 0 ] (  136 ) + 137
     (2)[ 0 0 1 0 ] [ 0 0 1 0 - 0 0 1 0 - 0 0 1 ] (  273 ) + 136
     (3)[ 0 0 1 1 ] [ 0 0 1 1 - 0 0 1 1 - 0 0 1 ] (  409 ) + 137
     (4)[ 0 1 0 0 ] [ 0 1 0 0 - 0 1 0 0 - 0 1 0 ] (  546 ) + 136
     (5)[ 0 1 0 1 ] [ 0 1 0 1 - 0 1 0 1 - 0 1 0 ] (  682 ) + 137
     (6)[ 0 1 1 0 ] [ 0 1 1 0 - 0 1 1 0 - 0 1 1 ] (  819 ) + 136
     (7)[ 0 1 1 1 ] [ 0 1 1 1 - 0 1 1 1 - 0 1 1 ] (  955 ) + 137
     (8)[ 1 0 0 0 ] [ 1 0 0 0 - 1 0 0 0 - 1 0 0 ] ( 1092 ) + 136
     (9)[ 1 0 0 1 ] [ 1 0 0 1 - 1 0 0 1 - 1 0 0 ] ( 1228 ) + 137
    (10)[ 1 0 1 0 ] [ 1 0 1 0 - 1 0 1 0 - 1 0 1 ] ( 1365 ) + 136
    (11)[ 1 0 1 1 ] [ 1 0 1 1 - 1 0 1 1 - 1 0 1 ] ( 1501 ) + 137
    (12)[ 1 1 0 0 ] [ 1 1 0 0 - 1 1 0 0 - 1 1 0 ] ( 1638 ) + 136
    (13)[ 1 1 0 1 ] [ 1 1 0 1 - 1 1 0 1 - 1 1 0 ] ( 1774 ) + 137
    (14)[ 1 1 1 0 ] [ 1 1 1 0 - 1 1 1 0 - 1 1 1 ] ( 1911 ) + 136
    (15)[ 1 1 1 1 ] [ 1 1 1 1 - 1 1 1 1 - 1 1 1 ] ( 2047 )

This solution is an index subdivision printer - it will synthesize the subdivision values for any bit width and
output the deltas as I've shown above =)