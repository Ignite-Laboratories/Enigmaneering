# `E2S1 - Index Diminishment`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### Binary Subdivision
The next part to recognize is the difference between numeric and logical binary data.  For synthesis, we work
with the numeric representation of data while logically managing the leading zeros.  That will make a lot
more sense as we proceed forward - for now, just know that value equivalence does _not_ imply logical equivalence.
A single missing zero could entirely destroy the logical structure of data - fun! =)

    Logical Binary Data ⬎           ⬐ Base 10 Value
               [ 0 0 1 0 1 0 1 0 ] (42)  
                   [ 1 0 1 0 1 0 ] (42)  
                         ⬑ Numeric Binary Data

Here, the logical form is a _byte_ - but that's not a hard requirement, just a universal standard.

This is what we ultimately will be exploiting to facilitate binary synthesis!  Some values can be stored in
_far less bits_ than they're logically stored at, relative to known points in the index.  

But what even _is_ a known "point" in an index?

Well, it's any point that can be _implicitly_ referenced from contextual information.  The midpoint, for example, 
is _always_ known to _implicitly_ be a one followed by all zeros _up to the index's bit length_.  But you can take 
this even _further_ by synthesizing repeating _**patterns**_ of bits across the index!  

Let's take an 11-bit index and break it into eight regions using a note (3-bit) pattern -

    "Note Diminishment of an 11 bit Index"
 
        Pattern              Synthesized Point        Value      Δ   
     (0) { 0 0 0 }   [ 0 0 0   0 0 0   0 0 0   0 0 ] (   0  ) + 292
     (1) { 0 0 1 }   [ 0 0 1   0 0 1   0 0 1   0 0 ] (  292 ) + 293
     (2) { 0 1 0 }   [ 0 1 0   0 1 0   0 1 0   0 1 ] (  585 ) + 292
     (3) { 0 1 1 }   [ 0 1 1   0 1 1   0 1 1   0 1 ] (  877 ) + 293
     (4) { 1 0 0 }   [ 1 0 0   1 0 0   1 0 0   1 0 ] ( 1170 ) + 292
     (5) { 1 0 1 }   [ 1 0 1   1 0 1   1 0 1   1 0 ] ( 1462 ) + 293
     (6) { 1 1 0 }   [ 1 1 0   1 1 0   1 1 0   1 1 ] ( 1755 ) + 292
     (7) { 1 1 1 }   [ 1 1 1   1 1 1   1 1 1   1 1 ] ( 2047 )

Literally any width target index can be evenly diminished by the _limit_ of the pattern's index simply by
repeating the pattern across it.  Why do I keep calling this 'diminishment' instead of 'subdivision'?  Because
unlike mathematical subdivision, the intervals are only _close enough._  If the value the synthesized point 
_should_ represent is a floating point number, binary patterning truncates it to the closest whole number _naturally._  

Much like a diminished chord, every point is as equidistant as possible from the last - except there's far more 
than _three_ diminished "chords" in an index!  Technically, you can diminish an index until each point is exactly 
one away from the next because the pattern bit length matches the index.  That also means this is a mechanism to 
lower the _resolution_ of the index while also providing a way to quickly "stride" through it.

Binary is truly the most beautiful counting system in existence =)

Diminishment, by itself, isn't anything special - in fact, we don't use it at all in the synthesis process!  However,
it provides a way to visualize counting to a number from a _synthetic point._  We will be counting specifically from 
the _midpoint,_ rather than _zero._

The solution here is a primitive demonstration to show that binary follows these diminishment rules for a provided 
index.  All of this has led me to posit a fundamental law -

    "The Law of Binary Index Diminishment"

        An index can be evenly diminished by the limit of a bit pattern's containing index from
        repeating the bits across the target, with the diminishment interval defined by the 
        numeric value of the pattern.

### Prove It
That's a lot easier than one might think!  But you must work from the _left_ side of the binary information
rightwards.  First, let's take a look at the abstract halving points of a 10-bit index -

        2¹⁰ = 1024

                ⬐ Everything to the right is a single repeated bit
        [ 1 0 0   0 0 0 0 0 0 0 ]  (512) ← The index's midpoint
        [ 0 1 0   0 0 0 0 0 0 0 ]  (256) ← The index's quarter point
        [ 0 0 1   0 0 0 0 0 0 0 ]  (128) ← The index's eighth point
           ⬑ Zeros are introduced proportionally with each halving

This is pretty obvious - we are simply halving the target index's limit by one power of two less with each 
iteration.  That being said, if you consider the first three bits to be a _pattern_ and the remaining bits 
to be zero, a formula arises - 

    𝑛 = The target bit width
    ℓ = The pattern index's limit
    𝑖 = The pattern interval value

    𝑥 = ⌊(2ⁿ/ℓ) * 𝑖⌋

    let 𝑛 = 7

        ⬐ The pattern
    [ 0 0 0   0 0 0 0 ]   (0) = ⌊(2⁷/8) * 0⌋
    [ 0 0 1   0 0 0 0 ]  (16) = ⌊(2⁷/8) * 1⌋
    [ 0 1 0   0 0 0 0 ]  (32) = ⌊(2⁷/8) * 2⌋
    [ 0 1 1   0 0 0 0 ]  (48) = ⌊(2⁷/8) * 3⌋
    [ 1 0 0   0 0 0 0 ]  (64) = ⌊(2⁷/8) * 4⌋
    [ 1 0 1   0 0 0 0 ]  (80) = ⌊(2⁷/8) * 5⌋
    [ 1 1 0   0 0 0 0 ]  (96) = ⌊(2⁷/8) * 6⌋
    [ 1 1 1   0 0 0 0 ] (112) = ⌊(2⁷/8) * 7⌋
                  ⬑ The trailing zeros

Now you can recursively apply this operation against an index one pattern bit width smaller in size until you
reach zero bits -

    𝑛 = 𝑛 - 3

          [ 0 0 0   0 ]   (0) = ⌊(2⁴/8) * 0⌋
          [ 0 0 1   0 ]   (2) = ⌊(2⁴/8) * 1⌋
          [ 0 1 0   0 ]   (4) = ⌊(2⁴/8) * 2⌋
          [ 0 1 1   0 ]   (6) = ⌊(2⁴/8) * 3⌋
          [ 1 0 0   0 ]   (8) = ⌊(2⁴/8) * 4⌋
          [ 1 0 1   0 ]  (10) = ⌊(2⁴/8) * 5⌋
          [ 1 1 0   0 ]  (12) = ⌊(2⁴/8) * 6⌋
          [ 1 1 1   0 ]  (14) = ⌊(2⁴/8) * 7⌋

    𝑛 = 𝑛 - 3

                  [ 0 ]   (0) = ⌊(2¹/8) * 0⌋
                  [ 0 ]   (0) = ⌊(2¹/8) * 1⌋
                  [ 0 ]   (0) = ⌊(2¹/8) * 2⌋
                  [ 0 ]   (0) = ⌊(2¹/8) * 3⌋
                  [ 1 ]   (1) = ⌊(2¹/8) * 4⌋
                  [ 1 ]   (1) = ⌊(2¹/8) * 5⌋
                  [ 1 ]   (1) = ⌊(2¹/8) * 6⌋
                  [ 1 ]   (1) = ⌊(2¹/8) * 7⌋
                           ⬑ NOTE: This is floored

So, let's put that all together and find the 4ᵗʰ interval of a 3 bit diminishment of an 11 bit index -

    The starting conditions -

                        ⬐ The pattern's bit width
                    [ 0 1 0 ]  (4) ← The pattern interval
                  [ 1 0 1 1 ] (11) ← The index bit width

    The target -

    [ 0 1 0 0 1 0 0 1 0 0 1 ]  (585) ← 0 1 0 repeated across the index

    The algorithm -

    [ 0 1 0 0 0 0 0 0 0 0 0 ]  (512)
          [ 0 1 0 0 0 0 0 0 ]   (64) +
                [ 0 1 0 0 0 ]    (8) +
                      [ 0 1 ]    (1) +
                              =  585

In essence, you are taking the 4ᵗʰ 8ᵗʰ of each subsequently smaller index and then summing the values together,
simply as a _byproduct_ of using an index to reference the data.  Ultimately, that can be wrapped up into a very
simple little formula to calculate a desired diminishment point -

<picture>
<img alt="Index Diminishment Formula" src="assets/diminishmentPoint.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>

I'm not sure how much more proof one would need - this appears to be a fundamental _law_ of binary indexes =)

But, and _far_ more importantly, we just executed the standard process of synthesis: using _starting conditions_
and an _algorithm_ to create a _target!_