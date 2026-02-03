# `E1S1 - Diminishment Proof`
### `Alex Petz, Ignite Laboratories, February 2026`

---

Before we continue on, I get to take a moment and mathematically _prove_ diminishment.

First, let's define how a diminishment interval synthesizes into a larger index.

<picture>
<img alt="Index Diminishment Formula" src="assets/diminishment interval.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>

It's not complex. But that doesn't necessarily _prove_ that tiling a pattern of bits across an index evenly subdivides index's 
potentials.  To do that, I'd like to hand walk through a summation of each _instance_ of the pattern in isolation to definitively
_prove_ what I'm saying.  This gives us an opportunity to start from a 1-bit diminishment, which also highlights the _halving points_
of an index and how they _**visually appear**_ in binary:

        " 1-bit Diminishment of a 4-bit Index "
    
        [ 0 0 0 0 ]  (0) ← Interval 0
        [ 1 1 1 1 ] (15) ← Interval 1
                      ⬑ Synthesized Potential

        " Isolated Instances of 'Interval 1' "

         ⬐ Instance Numbers          ⬐ Resulting halving points
        [0] | 1 0 0 0 | (8) ← The index's midpoint
        [1] | 0 1 0 0 | (4) ← The index's quarter point
        [2] | 0 0 1 0 | (2) ← The index's eighth point
        [3] | 0 0 0 1 | (1) ← The index's sixteenth point
       Instances ⬏       ⬑ Instance Values

This is pretty obvious - we are simply halving the target index to the next smaller power of two with each
iteration.  Next, let's expand the interval to a wider bit-width and derive each instance's _value._  I'll
start by identifying the value of the _**first**_ instance.  During summation, I'll shrink the target bit-width 
by the pattern's width on each step -

    " The First Patterned Instance Formula "

    let 𝑛 = The target index bit width
    let 𝑤 = The pattern index bit width
    let 𝑝 = The pattern value
    
    𝑓(𝑛, 𝑤, 𝑝) = ⌊ ( 2ⁿ / 2ʷ ) * 𝑝 ⌋

This means the first instance of `[ 0 1 0 ] (2)` against a 6-bit index should yield `[ 0 1 0 0 0 0 ] (16)` using the
above formula:

    𝑓(6, 3, 2) = ⌊ ( 2⁶ / 2³ ) * 2 ⌋
               = ⌊ ( 64 / 8 ) * 2 ⌋
               = ⌊ ( 8 ) * 2 ⌋
               = 16

Wonderful!  Now let's algorithmically sum this formula starting from the full index width, and then iteratively again
for each index one pattern width less wide.  I'll do so for every potential interval in parallel -

    " Hand Summation of a 3-bit Diminishment Over an 8-bit Index "

                            "Step 0"
    let 𝑛 = 8
    let 𝑤 = 3

     ⬐ 𝑝
    (0) | 0 0 0  -  0 0 0 0 0 |   (0) = ⌊(2⁸/2³) * 0⌋
    (1) | 0 0 1  -  0 0 0 0 0 |  (32) = ⌊(2⁸/2³) * 1⌋
    (2) | 0 1 0  -  0 0 0 0 0 |  (64) = ⌊(2⁸/2³) * 2⌋
    (3) | 0 1 1  -  0 0 0 0 0 |  (96) = ⌊(2⁸/2³) * 3⌋
    (4) | 1 0 0  -  0 0 0 0 0 | (128) = ⌊(2⁸/2³) * 4⌋
    (5) | 1 0 1  -  0 0 0 0 0 | (160) = ⌊(2⁸/2³) * 5⌋
    (6) | 1 1 0  -  0 0 0 0 0 | (192) = ⌊(2⁸/2³) * 6⌋
    (7) | 1 1 1  -  0 0 0 0 0 | (224) = ⌊(2⁸/2³) * 7⌋
        |←  3   →|            |   ⬑ 𝑓(8, 3, 𝑝)
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

    -----------------------------------------------------------

                   " Adding of Step Results "

                       Step 0 ⬎     Step 2 ⬎
    (0) | 0 0 0 0 0 0 0 0 |   (0) +  (0) + (0) = 0
    (1) | 0 0 1 0 0 1 0 0 |  (32) +  (4) + (0) = 36
    (2) | 0 1 0 0 1 0 0 1 |  (64) +  (8) + (1) = 73
    (3) | 0 1 1 0 1 1 0 1 |  (96) + (12) + (1) = 109
    (4) | 1 0 0 1 0 0 1 0 | (128) + (16) + (2) = 146
    (5) | 1 0 1 1 0 1 1 0 | (160) + (20) + (2) = 182
    (6) | 1 1 0 1 1 0 1 1 | (192) + (24) + (3) = 219
    (7) | 1 1 1 1 1 1 1 1 | (224) + (28) + (3) = 255
                              Step 1 ⬏

So, let's put that all together and validate that the 4ᵗʰ interval of a 3-bit diminishment over an 11-bit index
indeed matches our expected potential -

                     "The Desired Target"

    |←       11 Bits       →|
    | 0 1 0 0 1 0 0 1 0 0 1 |  (585) ← 0 1 0 repeated across the index

    -----------------------------------------------------------

               "The Starting Condition Phrase"

                                 |←     7 Bits    →|
                             (4) | 0 1 0 - 1 0 1 1 | (11)
    The diminishment interval ⬏                        ⬑ The target bit width

    -----------------------------------------------------------

                       "The Algorithm"

    |←       11 Bits       →|
    | 0 1 0 0 0 0 0 0 0 0 0 |  (512) ← ⌊(2¹¹/2³) * 2⌋
    |       0 1 0 0 0 0 0 0 | + (64) ← ⌊( 2⁸/2³) * 2⌋
    |             0 1 0 0 0 | +  (8) ← ⌊( 2⁵/2³) * 2⌋
    |                   0 1 | +  (1) ← ⌊( 2²/2³) * 2⌋
                              =  585

In essence, you are taking the 4ᵗʰ 8ᵗʰ (the interval value followed by the diminishment's resolution) of each subsequently 
smaller index and then summing the values together.  Ultimately, that yields the following formula for a _summation_ variant
of the more simple Diminishment Interval Formula.

<picture>
<img alt="Index Diminishment Summation Formula" src="assets/diminishment summation.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>

I'm not sure how much more proof one would need that you can simply repeat a _bit pattern_ across index and **always** get
as equidistant of points as possible - its a fundamental principle of indexes =)

_**Far more importantly,**_ however, we just inadvertently executed the essence of synthesis: using _starting
conditions_ and an _algorithm_ to recreate a larger _target._