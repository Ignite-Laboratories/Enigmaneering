# `E2S2 - Splitting the Byte`
### `Alex Petz, Ignite Laboratories, January 2026`

---

A measurement is more than just a value - it's an _implicit decision tree._  In the right context,
each addressable value could indicate a different branch another system should follow.  That makes a measurement
analogous to a _signaling symbol._

When measurements are strung together, these symbols produce a logical _**phrase**._

    " The Phrase "
        A logical sequence of diverse measurements.

There are two kinds of phrases - _explicit_ and _implicit._  _Explicit_ phrases have already
been decoded into known structures, while _implicit_ ones must actively be decoded.  An example
of an _explicit_ phrase is a "phrase literal" - which directly denotes the boundaries of each
measurement.  An implicit phrase requires walking the data bit-by-bit to form each measurement.

The simplest form of this is to split binary data into a _**vector**._

A vector has two components - a _direction_ and a _magnitude._

Fortunately for us, a single _bit_ technically fits the description of a _degenerate 1D vector_ with
zero magnitude - there's just no reference of what it means to go in the direction of _0_ vs _1._  To 
add magnitude, one would read the face value of the remaining binary information.

    tl;dr - binary can be parsed in phrase form as a vector

      why - in an infinitely wide register, You'll need to precisely locate data 
    