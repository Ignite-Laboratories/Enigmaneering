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

But what even _is_ a known index _point?_

Well, it's any point that can be _implicitly_ referenced from contextual information.  A known data point can always be
recalled simply by duplicating it - but the index midpoint, for example, is _always_ known to _implicitly_ be a one 
followed by all zeros.  But you can take this even _further_ by synthesizing repeating _**patterns**_ of bits across 
the index!  

Let's take an 11-bit index and break it into eight regions using a note (3-bit) pattern -

    "Note Diminishment of an 11 bit Index"

    let 𝑛 = The index width
    let 𝑤 = The pattern width
    let 𝑝 = The pattern value
    let  𝑣(𝑛, 𝑤, 𝑝) ↦ ⌊(2ⁿ / (2ʷ - 1)) * 𝑝⌋
    let 𝑑𝑣(𝑛, 𝑤, 𝑝) ↦ 𝑣(𝑛, 𝑤, 𝑝) - 𝑣(𝑛, 𝑤, 𝑚𝑎𝑥(𝑝 - 1, 0))
    where 𝑚𝑎𝑥(𝑎, 𝑏) returns the larger of 𝑎 and 𝑏 
 
                                      ⬐ "Synthesized Point"
              𝑝                      𝑣(𝑝)                         ⬐𝑑𝑣(𝑝)  
      (0) | 0 0 0 |   | 0 0 0   0 0 0   0 0 0   0 0 | (   0  ) + 292
      (1) | 0 0 1 |   | 0 0 1   0 0 1   0 0 1   0 0 | (  292 ) + 293
      (2) | 0 1 0 |   | 0 1 0   0 1 0   0 1 0   0 1 | (  585 ) + 292
      (3) | 0 1 1 |   | 0 1 1   0 1 1   0 1 1   0 1 | (  877 ) + 293
      (4) | 1 0 0 |   | 1 0 0   1 0 0   1 0 0   1 0 | ( 1170 ) + 292
      (5) | 1 0 1 |   | 1 0 1   1 0 1   1 0 1   1 0 | ( 1462 ) + 293
      (6) | 1 1 0 |   | 1 1 0   1 1 0   1 1 0   1 1 | ( 1755 ) + 292
      (7) | 1 1 1 |   | 1 1 1   1 1 1   1 1 1   1 1 | ( 2047 )
          |←  𝑤  →|   |←              𝑛            →|
              ⬑ 3                 11 ⬏

Literally any width index can be diminished by the _dark point_ of a pattern's index simply by repeating the pattern 
across it.  Why do I keep calling this 'diminishment' instead of 'subdivision'?  For two reasons - first, 𝑣(𝑝) isn't 
always a whole integer so we explicitly floor it so it can be represented in binary.  This causes the running delta to 
**naturally** be asymmetric as a byproduct of working in a binary space.  Second, and more importantly, the 
irregularly spaced intervals between these synthetic points represent _**implicitly**_ addressable ranges, and as such
deserve a unique way of being identified.

Much like a diminished chord, every point is as equidistant _as possible_ from the last - except there's far more 
than _three_ diminished "chords" in an index!  Technically, you can diminish an index until each interval is exactly 
one away from the next because the pattern's bit length matches the index.  That also makes this a mechanism to 
lower the _resolution_ of the index and provides a way to quickly "stride" through it.

Binary is truly the most beautiful counting system in existence =)

<picture>
<img alt="Index Diminishment Formula" src="assets/diminishment point.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>

In `tiny` you can synthesize any diminishment point on the fly by providing a binary measurement (which implicitly
provides the pattern interval and width as one variable) -

    let 𝑚 = A known interval measurement

    tiny.Synthesize.Point.Diminishment(𝑚, 𝑛) // Synthesizes diminishment interval 𝑚 across an 𝑛-wide index  

### Terminals, Intervals, and Reading Phrases

A **_single_** leading pattern can also be used to implicitly reference a sub-index on demand - called a _terminal
region._  Let's briefly look at the midpoint on an index again -

               |←      𝑛      →|
               | 1 - 0  ...  0 | (𝑛 / 2)  ← Midpoint
    Terminus Bit ⬏       ⬑ The Terminal Region

You can _widen_ the _terminus_ in order to identify a smaller region in the index - taking it from a
_bit_ to an _interval_ identifying the terminal region of addressable information.  An "interval" represents the
face value of the bit pattern in use.  Technically, that makes _**all**_ binary information an interval of its
containing index - but much like intervals in a musical chord, they are contextually implicit until you explicitly 
need to _describe_ their quality.  In diminishment, the term _interval_ directly describes the explicit _value_ 
of the bit pattern itself -

    let 𝑡 = The Terminal Bit Width

                       |←       𝑛       →|
                       |←  𝑡  →|←  𝑛-𝑡   →|
                       | 1 0 1 - 0 ... 0 |
     The Terminus Interval ⬏        ⬑ The Terminal Region

This brings me to an important note: numeric values, no batter the base, are _constructed_ right↤to↤left, but we still 
_read them_ from left↦to↦right.  The same applies to binary - the leftmost, or the "most significant bit," is 
considered index position `0` - whereas the rightmost, or the "least significant bit," is considered index position 
`𝑛-1`.  From a _value_ perspective, on the other hand, the value wouldn't _exist_ without a terminus - _the phrase 
would be empty!_  That's why data pinned to the _left side_ of the phrase is considered the _"terminal data."_  On the
other hand, data that is pinned to the _right side_ of the phrase is considered the _"genesis data"_ - intentionally
implying that it _originates_ from there, yet has the least effect on its current form.

        tiny.Synthesize.Point.Terminal(𝑝, 𝑛) // Synthesizes terminal point 𝑝 followed by zeros up to the provided width
        tiny.Synthesize.Point.Terminal(𝑝, 𝑛, 1) // Synthesizes terminal point 𝑝 followed by ones up to the provided width
        tiny.Synthesize.Point.Genesis(𝑝, 𝑛) // Synthesizes genesis point 𝑝 preceeded by zeros up to the provided width
        tiny.Synthesize.Point.Genesis(𝑝, 𝑛, 1) // Synthesizes genesis point 𝑝 preceeded by ones up to the provided width

To _work_ with terminal and genesis data you can use the phrase reading methods.  Read operations consume the provided
number of bits before returning both the read and remaining bits as separate objects.  Here's a list of the reading 
operations a phrase offers -

        let 𝑎 = A known phrase

        𝑎.Read(𝑛) // Reads from the left and returns two phrases: the read terminus bits and the remainder region
        𝑎.ReadFromEnd(𝑛) // Reads from the right and returns two phrases: the read initial bits and the remainder region

        𝑎.ReadNextBit() // Reads the next bit from the left plus the remainder
        𝑎.ReadFromEnd(𝑛) // Reads the next bit from the right plus the remainder
        𝑎.ReadMeasurement(𝑛) // Reads up to a word back as a measurement, rather than a phrase, plus the remainder
        𝑎.ReadUntilOne() // Reads from the left until it reaches a one and returns the zero count found plus the remainder

Reading from the end still retains the logical order of bits - the read operations merely _count_ in the direction of 
travel.  For example -

                                            Read ⬎            ⬐ Remainder 
        [ 0 1 0 0 1 1 0 1 ].ReadFromEnd(5) = [ 0 1 1 0 1 ] [ 0 1 0 ] 
        [ 0 1 0 0 1 1 0 1 ].Read(5)        = [ 0 1 0 0 1 ] [ 1 0 1 ] 

Read operations typically return an `ErrorEndOfBits` when you have read beyond the bounds of the phrases bits, but 
it's perfectly acceptable to simply _ignore_ the error because the operation will still return whatever it _could_ 
find.  For example - it's far more efficient to just read a smaller-than-𝑛 value by calling `𝑎.Read(𝑛)` 
than it is to first calculate the exact number of bits you have available to read.  `tiny` will gracefully stop 
reading when no more bits are left, but will provide you with some extra information _if you care._ Remember, 
a `measurement` gives us a unique quality which a `word` doesn't: _"is the data **empty?**"_

The utility of diminishment will come later on, but for now it's a wonderful primer on working with an index!

This solution is a primitive demonstration to show that binary follows these diminishment rules for any provided 
index.  All of this has led me to posit a fundamental principle of working with binary indexes -

    "The Binary Index Diminishment Principle"

            An index can be diminished by the dark point of a bit pattern's containing index by repeating the 
        pattern across the target, with the diminishment interval defined by the numeric value of the pattern.

### Prove It!
Technically, the formula is already written above - but I get to prove to you that _repeating the bit pattern_
across the index is equivalent to the more formalized `𝑣(𝑝)` formula above.  That's a lot easier than one might 
think! First, it's a lot easier to mathematically work from the _left_ side of the index rightwards.  Let's circle 
back to the halving points of an index again -

        Index 2¹⁰ (1024)

            ⬐ Everything to the right is repeated zeros
        | 1 - 0 0 0 0 0 0 0 0 0 |  (512) ← The index's midpoint
        | 0 1 - 0 0 0 0 0 0 0 0 |  (256) ← The index's quarter point
        | 0 0 1 - 0 0 0 0 0 0 0 |  (128) ← The index's eighth point
           ⬑ Zeros are introduced proportionally on the left with each halving

This is pretty obvious - we are simply halving the target index to the next smaller power of two with each 
iteration. That being said, if you consider the first three bits to be a diminishment _bit pattern_ and the 
remaining bits to be zero, a summable formula arises - 

    let 𝑛 = The target index bit width
    let 𝑤 = The pattern index bit width
    let 𝑝 = The pattern value
    
    𝑓(𝑛, 𝑤, 𝑝) = ⌊ ( 2ⁿ / 2ʷ ) * 𝑝 ⌋

So let's algorithmically sum this operation starting from the full index width, and then iteratively again
for each index one pattern width less wide -

                            "Step 0"
    let 𝑛 = 8
    let 𝑤 = 3

     ⬐ The pattern Interval
    (0) | 0 0 0  -  0 0 0 0 0 |   (0) = ⌊(2⁸/2³) * 0⌋
    (1) | 0 0 1  -  0 0 0 0 0 |  (32) = ⌊(2⁸/2³) * 1⌋
    (2) | 0 1 0  -  0 0 0 0 0 |  (64) = ⌊(2⁸/2³) * 2⌋
    (3) | 0 1 1  -  0 0 0 0 0 |  (96) = ⌊(2⁸/2³) * 3⌋
    (4) | 1 0 0  -  0 0 0 0 0 | (128) = ⌊(2⁸/2³) * 4⌋
    (5) | 1 0 1  -  0 0 0 0 0 | (160) = ⌊(2⁸/2³) * 5⌋
    (6) | 1 1 0  -  0 0 0 0 0 | (192) = ⌊(2⁸/2³) * 6⌋
    (7) | 1 1 1  -  0 0 0 0 0 | (224) = ⌊(2⁸/2³) * 7⌋
        |←  3   →|            |   ⬑ 𝑓(𝑛, 𝑤, 𝑝)
        |←         8         →|
    
    -----------------------------------------------------------

                            "Step 1"
    𝑛 = 𝑛 - 3

            (0) | 0 0 0 - 0 0 |   (0) = ⌊(2⁵/2³) * 0⌋
            (1) | 0 0 1 - 0 0 |   (4) = ⌊(2⁵/2³) * 1⌋
            (2) | 0 1 0 - 0 0 |   (8) = ⌊(2⁵/2³) * 2⌋
            (3) | 0 1 1 - 0 0 |  (12) = ⌊(2⁵/2³) * 3⌋
            (4) | 1 0 0 - 0 0 |  (16) = ⌊(2⁵/2³) * 4⌋
            (5) | 1 0 1 - 0 0 |  (20) = ⌊(2⁵/2³) * 5⌋
            (6) | 1 1 0 - 0 0 |  (24) = ⌊(2⁵/2³) * 6⌋
            (7) | 1 1 1 - 0 0 |  (28) = ⌊(2⁵/2³) * 7⌋
                |←  3  →|     |
                |←     5     →|

    -----------------------------------------------------------

                            "Step 2"
    𝑛 = 𝑛 - 3

                    (0) | 0 0 |   (0) = ⌊(2²/2³) * 0⌋
                    (1) | 0 0 |   (0) = ⌊(2²/2³) * 1⌋
                    (2) | 0 1 |   (1) = ⌊(2²/2³) * 2⌋
                    (3) | 0 1 |   (1) = ⌊(2²/2³) * 3⌋
                    (4) | 1 0 |   (2) = ⌊(2²/2³) * 4⌋
                    (5) | 1 0 |   (2) = ⌊(2²/2³) * 5⌋
                    (6) | 1 1 |   (3) = ⌊(2²/2³) * 6⌋
                    (7) | 1 1 |   (3) = ⌊(2²/2³) * 7⌋
                        |← 2 →|                    

So, let's put that all together and validate that the 4ᵗʰ interval of a 3 bit pattern repeated across an 11 bit index 
indeed matches our target -

                          "The Target"

    |←       11 Bits       →|
    | 0 1 0 0 1 0 0 1 0 0 1 |  (585) ← 0 1 0 repeated across the index

    -----------------------------------------------------------

                  "The Starting Conditions"

                               |←     7 Bits    →|
                           (4) | 0 1 0 - 1 0 1 1 | (11)
    The diminishment point ⬏                        ⬑ The target bit width

    -----------------------------------------------------------

                         "The Algorithm"

    |←       11 Bits       →|
    | 0 1 0 0 0 0 0 0 0 0 0 |  (512) ← ⌊(2¹¹/2³) * 2⌋
    |       0 1 0 0 0 0 0 0 | + (64) ← ⌊( 2⁸/2³) * 2⌋
    |             0 1 0 0 0 | +  (8) ← ⌊( 2⁵/2³) * 2⌋
    |                   0 1 | +  (1) ← ⌊( 2²/2³) * 2⌋
                              =  585

In essence, you are taking the 4ᵗʰ 8ᵗʰ of each subsequently smaller index and then summing the values 
together.  Ultimately, that can be wrapped up into a very simple little formula to calculate each desired 
diminishment point -

<picture>
<img alt="Index Diminishment Summation Formula" src="assets/diminishment summation.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>

I'm not sure how much more proof one would need - this appears to be a fundamental principle of binary indexes =)

_Far_ more importantly, however, we just inadvertently executed the standard process of synthesis: using _starting 
conditions_ and an _algorithm_ to create a _target!_ 