# `E1S0 - Data Indices`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### The Index Printer
Before we proceed any further, I get to briefly touch on how I talk about binary - and hopefully define
some standardized terminology so we can speak a similar language!  There isn't a lot of oddness beyond
what I've already defined - `tiny` also gives a few terms for longer stretches of binary data, but they're
honestly _highly_ contrived and don't provide any direct _utility._  _Describing_ binary information's qualities,
however, has _great_ utility!  So, the most important definition I would like to propose we collectively adopt
is the 'Index' standard of visualizing binary data.

First, let's examine the two states binary can remain in:

    1 <- Dark
    0 <- Light

This applies to _any length_ of binary information, and does _not_ imbue any value or size to the discussion:

    [ 1 1 1 1 1 1 1 1 ] <- A "dark" byte
    [ 0 0 0 0 0 0 0 0 ] <- A "light" byte

    [ 1 1 ] <- A "dark" crumb
    [ 0 0 ] <- A "light" crumb

    [ 1 1 1 ... 1 1 1 ] <- "Dark" data
    [ 0 0 0 ... 0 0 0 ] <- "Light" data

At scale, binary also can be considered "grey" - meaning it contains a mixture of 1s and 0s:

    [ 1 0 1 1 0 0 1 0 ] <- "Grey" data

These are the most important qualities - the next is that an _index_ of data represents all possible binary 
states a known bit-width could hold.  Indexes are represented just as they would on a vertical number 
line - meaning zero is at the bottom, larger values are placed above, and each value is well-ordered:

    A Crumb Index:

        Dark Side
         [ 1 1 ]
         [ 1 0 ]
         [ 0 1 ]
         [ 0 0 ]
        Light Side

    A Nibble Index:

         Dark Side
        [ 1 1 1 1 ] 
        [ 1 1 1 0 ]
        [ 1 1 0 1 ]
        [ 1 1 0 0 ]
        [ 1 0 1 1 ]
        [ 1 0 1 0 ]
        [ 1 0 0 1 ]
        [ 1 0 0 0 ]
        [ 0 1 1 1 ]
        [ 0 1 1 0 ]
        [ 0 1 0 1 ]
        [ 0 1 0 0 ]
        [ 0 0 1 1 ]
        [ 0 0 1 0 ]
        [ 0 0 0 1 ]
        [ 0 0 0 0 ]
         Light Side

At larger scales it gets far too excessive to print out every single value, so the data is often represented truncated
to highlight only the most important qualities of the index.  For example, this is a way to respresent a truncated 
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
           ├─────── The upper quarter-point
        [ 1 0 1 1 ]
        [ 1 0 1 0 ]
        [ 1 0 0 1 ]
        [ 1 0 0 0 ]
         ├───────── The mid-point
        [ 0 1 1 1 ]
        [ 0 1 1 0 ]
        [ 0 1 0 1 ]
        [ 0 1 0 0 ]
           ├─────── The lower quarter-point
        [ 0 0 1 1 ]
        [ 0 0 1 0 ]
        [ 0 0 0 1 ]
        [ 0 0 0 0 ]
         Light Side

This solution is quite simple - it merely prints out the entirety of whatever index of data you wish you visualize.
Note that this also represents a primitive _timer_ which uses bit width to create longer and shorter intervals of
time - meaning you wouldn't ever get this to finish printing out a 64 bit wide request.  So, keep it short =) 