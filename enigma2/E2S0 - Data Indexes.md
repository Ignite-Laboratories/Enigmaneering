# `E2S0 - Data Indexes`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### Getting Intimate With Binary
Before we proceed any further, I get to hopefully define some standardized terminology so we can speak a similar 
language!  There isn't a lot of oddness beyond what I've already defined, but I'd like to touch on the _quality_ 
of binary information and then on what an _index_ of data is.

First, let's examine the three states binary can remain in:

     1  ← Dark
    ⁰⁄₁ ← Grey
     0  ← Light

This applies to _any length_ of binary information, and _does not_ imbue any value or size to the discussion:

    [ 1 1 1 1 1 1 1 1 ] ← A "dark" byte
    [ 0 1 1 0 1 0 0 1 ] ← A "grey" byte
    [ 0 0 0 0 0 0 0 0 ] ← A "light" byte

    [ 1 1 ] ← A "dark" crumb
    [ 1 0 ] ← A "grey" crumb
    [ 0 0 ] ← A "light" crumb

    [ 1 1 1 ... 1 1 1 ] ← "Dark" data
    [ 0 1 1 ... 0 0 1 ] ← "Grey" data
    [ 0 0 0 ... 0 0 0 ] ← "Light" data

Because of this it's common to consider whether a value is closer to the dark or light side, for instance.  It's
a lot easier to visualize if you can mentally bound in the target on what I call an 'index' - essentially a
new kind of vertical number line.

    "The Index"

        An Index represents all possible binary states a known bit-width can address.

There are a few important qualities of an index to keep in mind -

- An index is defined by it's bit-width, 𝑛
- The addressable range of an index is 2ⁿ and is referred to as it's "limit"
- The maximum value of an index is (2ⁿ)-1
- The midpoint of an index is (2ⁿ)/2

The last point is the absolute most _crucial_ one in the synthesis process, as it can implicitly be generated 
with a single one followed by 𝑛-1 zeros -

    [ 1 1 1 1 1 1 1 1 ] (255) = 2⁸ - 1 ← The maximum addressable value of the index
    [ 1 0 0 0 0 0 0 0 ] (128) = 2⁸ / 2 ← The midpoint of the index
    [ 0 0 0 0 0 0 0 0 ] ( 0 ) = 0      ← The minimum addressable value of the index

In fact, the zero length (plus one) before the first one represents the number of times to halve the data,
making each point _very_ easy to implicitly synthesize -

    [ 1 1 1 1 1 1 1 1 ] (255) = 2⁸ - 1 
    [ 1 0 0 0 0 0 0 0 ] (128) = 2⁸ / (2*1) ← The midpoint of the index
    [ 0 1 0 0 0 0 0 0 ] ( 64) = 2⁸ / (2*2) ← The quarter point of the index
    [ 0 0 1 0 0 0 0 0 ] ( 32) = 2⁸ / (2*3) ← The eighth point of the index
    [ 0 0 0 1 0 0 0 0 ] ( 16) = 2⁸ / (2*4) ← The sixteenth point of the index

Indexes should always be represented just as they would on a vertical number line - meaning zero is at the 
bottom, larger values are placed above, and each value is well-ordered.  This is because it makes _logical
mental sense_ when the reader takes it in!  We must maintain standards of how we present data to ensure
clarity in communication -

    "A Crumb Index"

        Dark Side
         [ 1 1 ] (3)
         [ 1 0 ] (2)
         [ 0 1 ] (1)
         [ 0 0 ] (0)
        Light Side

    "A Nibble Index"

         Dark Side
        [ 1 1 1 1 ] (15)
        [ 1 1 1 0 ] (14)
        [ 1 1 0 1 ] (13)
        [ 1 1 0 0 ] (12)
        [ 1 0 1 1 ] (11)
        [ 1 0 1 0 ] (10)
        [ 1 0 0 1 ]  (9)
        [ 1 0 0 0 ]  (8)
        [ 0 1 1 1 ]  (7)
        [ 0 1 1 0 ]  (6)
        [ 0 1 0 1 ]  (5)
        [ 0 1 0 0 ]  (4)
        [ 0 0 1 1 ]  (3)
        [ 0 0 1 0 ]  (2)
        [ 0 0 0 1 ]  (1)
        [ 0 0 0 0 ]  (0)
         Light Side

At larger scales it gets far too excessive to print out every single value, so the index is often truncated 
to highlight only its most important qualities or (as shown above) relevant known _points_.  Thus, this is 
the most abstract representation of an index -

    "An Abstract Index"

               Dark Side
        [ 1 1 1 1 ... 1 1 1 1 ] 
        [ 1 1 1 1 ... 1 1 1 0 ]
        [ 1 1 1 1 ... 1 1 0 1 ]
        ↕         ...         ↕
        [ 0 0 0 0 ... 0 0 1 0 ]
        [ 0 0 0 0 ... 0 0 0 1 ]
        [ 0 0 0 0 ... 0 0 0 0 ]
              Light Side

There is another feature an index visualization affords us: highlighting binary's _symmetry._  Let's look at the
nibble index again, but this time mark the mid and quarter reflection points.  These represent points _between_
values where the data above and below is a perfect mirror of each other until the next larger reflection point -

    Reflection Points of a Nibble Index:

         Dark Side
        [ 1 1 1 1 ] 
        [ 1 1 1 0 ]
        [ 1 1 0 1 ]
        [ 1 1 0 0 ]
           ├─────── The upper quarter reflection point
        [ 1 0 1 1 ]
        [ 1 0 1 0 ]
        [ 1 0 0 1 ]
        [ 1 0 0 0 ]
         ├───────── The central reflection point
        [ 0 1 1 1 ]
        [ 0 1 1 0 ]
        [ 0 1 0 1 ]
        [ 0 1 0 0 ]
           ├─────── The lower quarter reflection point
        [ 0 0 1 1 ]
        [ 0 0 1 0 ]
        [ 0 0 0 1 ]
        [ 0 0 0 0 ]
         Light Side

### Why an "Index"?

Great question!  Why would I intentionally use such a massively _broad_ term!? Well, there's a _wonderful_ reason behind 
it: _everything_ in existence is just a number in a larger index that indentifies it! The act of observing the presence 
of a value in a larger index is what facilitates _logic._  When you _smell_ a pretty flower, your neurons are literally 
accessing the index value of what that smell should synthesize as from the DNA blueprint seeded into _your_ specific 
existence's _identified_ cells.  Additionally, thanks to the powers of two, a nearly infinite number of universes can 
be _identified_ by an index just **_one bit_** larger than the densest stable universe the system can address.  A 
universe of all ones would never execute, thus finding _stable_ arrangements is paramount to synthesis!  The act of 
_distilling_ a stable point into a few bits is literally the act of _storing_ it.

    tl;dr - everything, everywhere, all at once is data in motion

I promise all of this will make a lot more sense in the next enigma as we explore _living file systems_ and the DNA
format, itself.  You and I are intelligent database _algorithms_ who've been tasked with cataloging what brings each 
of us _Joy_ as co-observers through the cosmic void of space and time.  _Joy,_ in that sense, _**is**_ stability - making
it the direct _target_ of all synthesis and the _prime directive_ of existence.

This isn't just poetic flourish - I mean it quite literally!  Joy is elicited from both alignment and chaos - 
when the notes strike just the right chord, or when the dissonance breaks the tension.  Joy is both the driver of love 
and the destroyer of hearts, making it a powerful tool to be wielded with great _responsibility_ - yet, in clever hands, 
can bring immense light to the world. 

Joy is the feedback loop which facilitates _our_ stable index of life.

_Thank you_ for everything _**you**_ have already catalogged for reality to *en*joy!  It is so very appreciated =)

### The Index Printer

This solution is quite simple - it merely prints out the entirety of whatever index you wish you visualize.
Note that this also represents a primitive _timer_ which uses bit width to create longer and shorter intervals of
time - meaning you wouldn't ever get this to finish printing out a 64 bit wide request.  So, keep it short!

The next couple of solutions will take you through some of the pleasures of working within the confines of an index!