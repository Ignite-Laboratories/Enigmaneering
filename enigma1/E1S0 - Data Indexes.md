# `E1S0 - Data Indexes`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### The Index Printer
Before we proceed any further, I get to briefly touch on how I talk about binary - and hopefully define
some standardized terminology so we can speak a similar language!  There isn't a lot of oddness beyond
what I've already defined, but the most important definition I would like to propose we collectively adopt
is the 'Index' standard of visualizing binary data.

First, let's examine the three states binary can remain in:

     1  🡨 Dark
    ⁰⁄₁ 🡨 Grey
     0  🡨 Light

This applies to _any length_ of binary information, and _does not_ imbue any value or size to the discussion:

    [ 1 1 1 1 1 1 1 1 ] 🡨 A "dark" byte
    [ 0 1 1 0 1 0 0 1 ] 🡨 A "grey" byte
    [ 0 0 0 0 0 0 0 0 ] 🡨 A "light" byte

    [ 1 1 ] 🡨 A "dark" crumb
    [ 1 0 ] 🡨 A "grey" crumb
    [ 0 0 ] 🡨 A "light" crumb

    [ 1 1 1 ... 1 1 1 ] 🡨 "Dark" data
    [ 0 1 1 ... 0 0 1 ] 🡨 "Grey" data
    [ 0 0 0 ... 0 0 0 ] 🡨 "Light" data

Because of this, it's common to consider that a value is 'close to the dark side', for instance - which
is very easy to mentally conceptualize if you can bound in the target on a kind of number line called
an 'index.'  

    "The Index"

        An Index represents all possible binary states a known bit-width could hold.

Since it makes the math _infinitely_ eaiser, the upper _**limit**_ of an index is considered to be 2ⁿ (where 𝑛 
represents the bit length of the index) while the upper maximum **_value_** of an index is (2ⁿ)-1.  This means
that we consider an 8-bit index to be a "256" index, even though it can never address the value "256".  This has 
a _very specific purpose_ as it makes the midpoint of an index equivalent to 1 followed by all zeros and a single 
division of 2 from the limit.

    [ 1 1 1 1 1 1 1 1 ] (255) | (2⁸)-1 🡨 The maximum value of the index
    [ 1 0 0 0 0 0 0 0 ] (128) | (2⁸)/2 🡨 The midpoint of the index
    [ 0 0 0 0 0 0 0 0 ] (0)   | 0      🡨 The minimum value of the index

Indexes are represented just as they would on a vertical number line - meaning zero is at the bottom, larger 
values are placed above, and each value is well-ordered:

    A Crumb Index:

        Dark Side
         [ 1 1 ] (3)
         [ 1 0 ] (2)
         [ 0 1 ] (1)
         [ 0 0 ] (0)
        Light Side

    A Nibble Index:

         Dark Side
        [ 1 1 1 1 ] (15)
        [ 1 1 1 0 ] (14)
        [ 1 1 0 1 ] (13)
        [ 1 1 0 0 ] (12)
        [ 1 0 1 1 ] (11)
        [ 1 0 1 0 ] (10)
        [ 1 0 0 1 ] (9)
        [ 1 0 0 0 ] (8)
        [ 0 1 1 1 ] (7)
        [ 0 1 1 0 ] (6)
        [ 0 1 0 1 ] (5)
        [ 0 1 0 0 ] (4)
        [ 0 0 1 1 ] (3)
        [ 0 0 1 0 ] (2)
        [ 0 0 0 1 ] (1)
        [ 0 0 0 0 ] (0)
         Light Side

At larger scales it gets far too excessive to print out every single value, so the index is often represented 
truncated to highlight only its most important qualities.  For example, this is a way to respresent a truncated 
index of any bit width:

    An Abstract Index:

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
nibble index again, but this time mark the mid and quarter reflection points.  These represent points between
values where the data above and below is a perfect mirror of each other, until the next major reflection point.

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

This solution is quite simple - it merely prints out the entirety of whatever index of data you wish you visualize.
Note that this also represents a primitive _timer_ which uses bit width to create longer and shorter intervals of
time - meaning you wouldn't ever get this to finish printing out a 64 bit wide request.  So, keep it short =) 