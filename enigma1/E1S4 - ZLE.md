# `E1S2 - Zero Length Encoding`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### The Gold Sixty-Four Standard
At the beginning of this enigma I said that we can treat 2⁶⁴ as a kind of 'practical infinity' - this is where
that realistically applies the most!  The concept of ZLE is simple: store smaller values in as few bits as possible,
while adding as few bits as possible to larger values.  It's a tug-of-war between efficiency and resolution, and
I've literally spent hundreds of hours with pen and paper meticulously curating exactly the most efficient way to
accomplish this goal.  I probably looked like a lunatic writing copious amounts of ones and zeros on a clipboard
by hand at my local diner, but to the ladies there I say 'thank you very much for the coffee.'

Now, the idea is simple - but the specific _ranges_ to encode are entirely contextual!  Some situations might call
for a zero to four bit range, others up to sixty four bits, or even just 'read the rest of the data'.  As a result,
rather than specifying a _hard rule_ for ZLE, I'd rather give an encoding _scheme_ and allow you to articulate what
the map looks like for your given context.  Again, we'll talk about defining _growth schemes_ later on - but for now,
let's look at the anatomy of a ZLE phrase:

          ⬐ Terminus
    [ ... 1 ] [ ⁰⁄₁ ... ]
       ⬑ Key     ⬑ Projection Bits

The zeros, plus the one, are considered the 'key' - while the bit range identified by the key is considered 
the 'projection'.  The meaning of the key, again, is up to your context!  So, let's talk about the most important
key range I've found thus far - the 'gold standard' - `Fuzzy.SixtyFour`

Fuzzy is a facet of tiny that provides access to fuzzy operations - like reading variable widths of data based on
a key value.  There's a lot of things Fuzzy provides, but ZLE is chief of it's list.

#### `The Fuzzy.SixtyFour Map`

        Key | Projection Bit Range
          1 | 4
        0 1 | 8
      0 0 1 | 16
    0 0 0 1 | 32
    0 0 0 0 | 64

This has a few notable qualities - first, it's _bounded,_ which means that you read up to a fixed width of bits (in
this case, 4) before interpreting a projection value.  Second, it limits out at 2⁶⁴ - making it fit within our 
'practical infinity' bit range nicely.  Lastly, it does _not_ offer a zero bit length condition - which is desirable
in our case because we typically don't need to encode that _zero_ '0's zipped off of our data!

### Alternatives
There are a couple of notable alternatives as well to consider, depending on your specific application:

#### `The Fuzzy.Five Map`

This map gives up to 32 unique addressable values, but they can be interpreted cumulatively to give a range
_nearly_ identical to a six bit measurement.

        Key | Projection | Range  | Cumulative Interpretation
          1 |     1      | 0 - 1  |  0 - 1
        0 1 |     2      | 0 - 3  |  2 - 5
      0 0 1 |     3      | 0 - 7  |  6 - 13
    0 0 0 1 |     4      | 0 - 15 | 14 - 29
    0 0 0 0 |     5      | 0 - 31 | 30 - 61

#### `The Fuzzy.Power Map`

This map can be interpreted against a power of two to give a futher projection range to read into the data,
allowing more granular bit lengths to be addressed than 16, then 32, then 64.

         Key | Projection | Value Range | Power Interpretation
           1 |      2     |   1 - 4     |      2ⁿ - 1
         0 1 |      3     |   1 - 8     |      2ⁿ - 1
       0 0 1 |      4     |   1 - 16    |      2ⁿ - 1
     0 0 0 1 |      5     |   1 - 32    |      2ⁿ - 1
     0 0 0 0 |      6     |   1 - 64    |      2ⁿ - 1

#### `The Unbounded Fuzzy.ZLE Map`

This is the standard _unbounded_ ZLE map.  It walks the bit range _as a power of two_ defined by the zero length.
It's utility is, well, minimal at this point - but it still deserves a formal definition.

    NOTE: This will overflow if you let it read too far =)
    
            Key | Projection
              1 | 1 [2⁰]
            0 1 | 2 [2¹]
          0 0 1 | 4 [2²]
        0 0 0 1 | 8 [2³]
               ...
          𝑛   1 | 2ⁿ
