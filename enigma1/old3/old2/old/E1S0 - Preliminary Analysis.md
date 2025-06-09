# `E1S0 - Preliminary Analysis`
### `Alex Petz, Ignite Laboratories, May 2025`

---

### What's the general amount the data can divide by two?
The first thing we need to know before we can start writing an encoding scheme is how many permutations
of action we even need to encode. While the _3n+1_ only has a single permutation, the division by two
side has an infinite number of times it could be called before reaching 1.  So, we need some heuristics
of what the average reduction range is.

I'm also going to officially coin a term here: _binary reduction_

This is, abstractly, "division by two of an arbitrary amount," and at the same time implies "any shrinkage
of the number of bits necessary to represent data" as division by two will _always_ reduce the bit-length
by a `n` bits per `2ⁿ` reduction.

The solution provided here performs such analysis.

### Tiny
This solution is the first introduction into the `tiny` module.  Tiny was written for general binary
bit analysis and operates off of two important structures - the `Measurement` and the `Phrase`:

**Measurement** - This is a variable slice of bits from 0 to 32 bits wide, making it always convertable to a
`int`.  Since storing bits individually is really inefficient, this stores them as bytes plus a remainder of
bits.  This allows binary values to be treated as variable "windows" that can have intelligent decisions made
from local context - most of which happens through the _Phrase_.

**Phrase** - A slice of measurements, simply put!  However, this structure empowers the measurement in many
useful ways - such as entirely _re-aligning_ where the measurement boundaries are, or by consuming arbitrary
amounts of bits at a time to make decisions of how to proceed with analysis.

While this library has lots of unit tests, it's entirely a labor of love by one guy who quite literally 
_dreams_ in 1s and 0s.  Some of the ideas might feel a bit strange, such as _trifurcation_ of binary data, but
I promise there is a method to my madness!

### Results
The results for my concept of using 3N+1 to encode are _quite_ damning - here's the results of how many times
each power of two was utilized by 2k of random binary data:

    Power | Times
        0 | 51913 
        1 | 38401 
        2 | 12039
        3 | 5424
        4 | 3976
        5 | 1409
        6 | 738
        7 | 353
        8 | 206
        9 | 93
       10 | 38
       11 | 16
       12 | 14
       13 | 8
       14 | 3

If the number of times power 1 was hit was low, 3N+1 _directly_ could be feasible.  However, this doesn't
stop me - I have ideas I've gleaned from what I've learned here today.

Running further, my conjecture is _immediately_ disproven - but that's okay!  It still holds statistically
likely that it would only ever have duplicates of a small bit-width, the height of which is not known.

I still have a better solution for approximating binary information - but here's the results of 7
iterations of 3N+1's power of 2 reductions:

    Power | Average Times
        0 | 210501
        1 | 156345
        2 | 48204
        3 | 22001
        4 | 16083
        5 | 5595
        6 | 3069
        7 | 1425
        8 | 778
        9 | 370
        1 | 182
        1 | 98
        1 | 39
        1 | 19
        1 | 10
        1 | 5
        1 | 4
        1 | 2
        1 | 1
        1 | 2

    Found Duplicates from 2kb source data:
        11
        101
        111
        1011
        1101
        10001
        10011
        11101
        100101
        110001
        100101
        110001
