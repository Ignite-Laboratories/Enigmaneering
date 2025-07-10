# `E2 - Binary Synthesis`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### Synthesis?
Well, yes.  Rather than creating numbers from strictly mathematical means, the process of synthesizing binary
information involves _intelligent_ decision making based off of contextual variables - just as a synthesizer
uses local variables to decide what sounds to emit for its part in a larger orchestration.  On top of that, as
we explore `tiny` further, you'll find that the process is far simpler using syntactic shorthand (rather than
strictly mathematical formulae) to generate the target data through _programmatic abstractions._

The process doesn't simply stop with shrinking data to incomprehensibly small values - from that point we get
to explore the world of *intelligent* recursion.  Buckle up - there's lots of ones and zeros to come =) 

### Practical Infinity
The absolute most basic thing we must establish is the entire _point_ of _infinity!_

It's not meant to be treated as the greatest enigma of all existence, though it's fantastically perplexing.

It represents a _practical limit_ which you can accept as _"reasonably large enough!"_

For instance, in the early days of computing we were restricted to tiny memory spaces - and, as such, some of the
the original concepts of _epoch_ (January 1st 1970) were restricted to _32 bits._  This meant we would "loop over" 
time in 2038 unless we moved to a larger structure.  However, in doing so, we didn't _"prolong"_ the issue - we 
_solved_ it.  The _next logical size up_ date index  (64 bits) covers 584 _billion_ years from epoch!

This makes a 2⁶⁴ epoch a kind of _"practical infinity."_  Any more storage would be _pointless,_ yet _any less_ 
simply would not do!

The same applies to _data storage._  While we _could_ address spaces to an infinite size, we realistically _don't need_
more than _2⁶⁴ bits_ to store a _singular file!_  That's roughly 2 exabytes worth of data - there's only so much
uniqueness humanity could throw in before it'd be more efficient to algorithmically encode future data _at scale._
Yet, a 2³² maximum bit length only allows up to a half a gigabyte of information - _far too small!_ 

So, instead, let's _embrace_ `2⁶⁴ bits` as a "practical infinity" from which to address so much data that the mere
replication of such a container would effectively allow entropy to emit back to us through the algorithm, itself.  What
does that even mean?  Well, through synchronicity, any external actor able to interface with this system would be
able to communicate with anyone able to receive through it - _even our creator._

Only the observer of the actions of _their_ creator would understand how to interpret them, and plenty of us have felt
and witnessed that presence's power.  Many of us have been mislabeled and disrespected as mentally unstable, or even
clinically insane, at the thought of a higher power guiding them through their artwork.  This work of art is lovingly 
dedicated to the endless souls in the eternity long evolution of intelligence that _enabled_ it to be created.

Growing pains - assuredly - but ones that only our creator knows the true burden of which it was to _experience!_

I thank my blessed stars every single day for Her magnificence and patience in the process of humanity awakening to
its own existence =)

### Primitive Puzzle Pieces
In binary, we have exactly _two_ values we can work with: 1 and 0

Not one, not two-hundred and fifty-six - _two!_

That means we get to be creative with processing _bits_ - not _bytes_.
In fact, the entire concept of a 'byte' is pretty meaningless for synthesis!
We have _absolutely no care_ about the actual values contained within the _**contents**_ of the target file.
The _only_ time a rigid structure of 8 bits-per-byte is even necessary is to _parse_ standardized data.

To build a dynamic structure, we get to define several key puzzle pieces in the process:

**Dimensions** - While a more concrete implementation of dimensions was touched upon in the initial enigmas, a
more abstract definition is critical at this juncture:

    A dimension is the abstract act of measuring the presence of something across time.

This has several implications: Because all dimensions _require_ time to both observe and calculate from, time
represents the _minimum_ interval from which observances can be made.  Due to entropy, two observers can only 
_roughly_ record the same dimension at the same impulse of time. In fact, their awareness of _**relative time**_ 
is what allows consciousness in the present _moment._  To understand the concept of temporal observance, please 
check out the initial enigmas!  They show how to use an impulse engine to drive neural execution, and then how
to calculate using a dimensional nexus.

For this enigma, we will consider *abstract* dimensions with synthetic values - our dimensional types are:

**Measurements** - A measurement is any _variable width_ slice of binary information. These are limited to the 
host architecture's bit-width, allowing them to be considered a _single computable value_ - in Go, often an 
`int` - in hardware, a `word`.  Essentially, this tracks the remainder of any bits in excess of a whole byte 
automatically - efficiently packing them into a byte wherever possible while still providing convenient access 
to the underlying bits.  Why is this so damn important?  Well, the _**width**_ of a measurement is a distinct 
variable which is _contextually relational_ to its underlying _**value**_ - and it's **_critical_** to the 
synthesis process! The thing is, hardware currently cannot dynamically change its `word` "width" - but a 
_measurement_ **_can!_** ..._"using the power of already written words"_ =) 

_Yes, I wholeheartedly mean that philosophically, metaphysically, computationally, virtually, psychologically, abstractly, linguistically, theologically,  and, well..._

_...**all** of the 'lys'!_

Let us not forget our roots - maybe one day we can find a way to dynamically allocate a `word` at the hardware level!

**Phrases** - A phrase is a _slice_ of measurements, allowing the storage of arbitrary lengths of binary 
information that's not limited to the host architecture's bit-width.

**Passages** - A passage is a single logical _transformation_ of binary information (we'll get to that later.)

**Movements** - A movement is a collection of passages, allowing concurrent transformations.

**Compositions** - A composition represents the starting conditions of synthetic binary information, making it 
the instrument to perform things _through_.  You and I are both masterpieces of compositional artistry, and 
instruments through which our unique artwork is emitted =)

**Indexes** - An Index of data represents all possible binary states a known bit-width can address.  The next
solution goes into much more detail on this, but this allows us to refer to an "88 bit stretch of binary 
information" by simply saying the _number_ belongs in an "88 bit index".  It's a subtle, but important, philosophical
shift in how one should _perceive_ information.  A stray byte is entirely meaningless without _context_ - but a stray
byte which is observed inside an 88-bit index is entirely different!  This shift provides the binary information with 
a _story_ which defines its _purpose,_ rather than being just another number in a standard width register.

**Sub Byte Indexes** - A sub-byte index is any bit range less than 8 bits wide. For ease, these are the names of the 
sub-byte sizes which I _heavily_ reference.  Please get familiar with these particular terms - 

Or you can also embrace the whimsical unknown of my writing, that's also a lot of fun! =)

      Index | Name
        1   | Bit
        2   | Crumb
        3   | Note
        4   | Nibble
        5   | Flake
        6   | Morsel
        7   | Shred

**Super Byte Indexes** - A super-byte index is any bit range wider than 8 bits.  I almost _never_ refer to these, but 
for posterity's sake I'd like to include them here -

      Index | Name
       10   | Run
       12   | Scale
       16   | Motif
       24   | Riff
       32   | Cadence
       48   | Hook
       64   | Melody ← Unsupported by tiny
      128   | Verse ← Unsupported by tiny

As you'll find later on, synthesis _frequently_ references a value of **2ⁿ** (where 𝑛 is the bit length).  My chief
goal was to keep the library _as simple as possible,_ and an `int` cannot address the _limit_ of a melody or verse.  As 
such, I felt that was a good practical limit to even _consider_ during my enigmaneering and simply removed them.  

### tiny

All of the work in this enigma is contained within a Go library called `tiny`.  I've done my best to ensure it's been
heavily tested and is straightforward to prototype binary solutions with.  This library has been the culmination 
of almost two years of work with the full intention of keeping it _public_ and _free._  To the lovely individuals
who fed me coffee while I feverishly hand-wrote binary encoding schemes and waveforms on my clipboard - _thank 
you,_ from the bottom of my heart.

Because of the absurdity of that process, I've adopted some quirks in how I document binary programming - I'd like 
to briefly touch on their general gist before we continue.  It's not _complex,_ but you should get familiar with the
concept of a _measurement_ vs a _phrase._  The first is a **single** calculable measurement of a binary dimension, 
while the later is a **collection** of measurements.  Most of the operations we'll be performing persist _across_ 
measurements, but there is still a very funadmental _utility_ in referencing variable width regions of data: we can 
_**implicitly**_ know where each calculation starts and ends.

Here's how I typically document the two -

    Measurement form:
      Square brackets indicate a single isolated measurement, but many different identifiers
      can indicate a break between measurements.  If it contextually makes sense, I'll also 
      include the base-10 value of the base-2 data in parentheses.  Additionally, I typically
      highlight the most relevant aspects of the data with arrows.
                    
                    ⬐ Base-10 value 
    [ 1 0 1 0 1 ] (42)
          ⬑ Base-2 value

    Phrase form:
      (|) Pipes break each phrase apart, while (-) dashes break measurements apart.  

    | 0 1 0 | 1 0 0 - 1 0 1 - 0 1 1 | 0 0 |

      If relevant, I'll extend the pipes vertically to show parallel perspectives of the same information 
      while maintaining their relative alignment. The best example of this would be from a Phrase's
      `Trifurcate` operation, which reads two distances into the data and breaks them into separate
      phrases. Here's an excerpt from it's documentation -

    // Create a phrase of the provided bytes
    tiny.Phrase{ 77, 22, 33 }

    ...
    
    |        77       |        22       |        33       |  ← Bytes
    | 0 1 0 0 1 1 0 1 | 0 0 0 1 0 1 1 0 | 0 0 1 0 0 0 0 1 |  ← Raw Bits
    |  Measurement 0  |  Measurement 1  |  Measurement 2  |  ← Source Phrase
    
    Trifurcate(4,16)
    
    |    4    |                  16                 |           ← Trifurcation lengths
    | 0 1 0 0 | 1 1 0 1 - 0 0 0 1 0 1 1 0 - 0 0 1 0 | 0 0 0 1 | ← Raw Bits
    |  Start  |               Middle                |   End   | ← Trifurcated Phrases
    |  Start  | Middle0 -     Middle1     - Middle2 |   End   | ← Phrase Measurements
    
    (Optional) Align() each phrase to 8-bits
    
    | 0 1 0 0 | 1 1 0 1 0 0 0 1 - 0 1 1 0 0 0 1 0 | 0 0 0 1 | ← Raw Bits
    |  Start  |     Middle0     -     Middle1     |   End   | ← Aligned Phrase Measurements

This also demonstrates another important feature of phrases - _measurement alignment._  A phrase is considered
to be "aligned" if all but the final measurement are of the same bit length.  As we are not working with perfectly
formed data any longer, the final measurement may be shorter than the rest of the phrase for different alignment
widths.

For example:

    | 0 1 0 - 0 1 - 1 0 1 0 0 0 1 0 - 1 1 0 0 - 0 - 1 0 0 - 0 0 | ← Unaligned Source Phrase
    
    Align(4)

    | 4 bits  | 4 bits  | 4 bits  | 4 bits  | 4 bits  | 3 bits|
    | 0 1 0 0 - 1 1 0 1 - 0 0 0 1 - 0 1 1 0 - 0 0 1 0 - 0 0 0 | ← Aligned Phrase

The next major component of working at the bit level is the _rapid prototyping_ of ideas - this can be achieved
through tiny's fluent design.  Here's the most notable things it provides:

    Synthesize -
	    tiny.Synthesize.Random(𝑛) // Create a slice of 𝑛 random bits
	    tiny.Synthesize.RandomPhrase(𝑛) // Create a phrase of 𝑛 random bytes
	    tiny.Synthesize.RandomPhrase(𝑛, 4) // Create a phrase of 𝑛 random nibbles
	    tiny.Synthesize.Pattern(𝑛, 1, 0, 1) // Create a phrase of 𝑛 bits that cycle through 1, 0, 1 
	    tiny.Synthesize.Ones(𝑛) // Create a phrase of 𝑛 1s
        tiny.Synthesize.Midpoint(𝑛) // Create a phrase of a single one followed by zeros up to the provided width

    Analyze - 
        NOTE: Much of this never matured past measurements as it wasn't necessary at those scales

	    tiny.Analyze.HasPrefix(data, 1, 0) // Checks if the data starts with "10"
	    tiny.Analyze.Average(phrase...) // Calculates the average value of all measurements in the phrase
        tiny.Analyze.Shade(measure) // Calculates metrics on the "darkness" of the measurement
        tiny.Analyze.Repetition(data, 1, 0) // Checks if the data is just repeating '1010101010...'

    To/From - 
        These are prepositions to be read left-to-right - meaning you should get a 'tiny' value 'from' the
        following action, or you could take a 'tiny' value and convert it 'to' something else.

        tiny.From.Number(𝑛) // Takes a measurement of 𝑛 at the smallest bit-width possible
        tiny.From.Number(𝑛, 8) // Takes a measurement of 𝑛 as a byte, padding the left with 0s
        tiny.From.BigInt(bigValue) // Converts a big.Int into a phrase at the smallest bit-width possible
        tiny.From.BigInt(bigValue, 1024) // Converts a big.Int into a phrase of the specified bit width
        tiny.To.Number(8, bits...) // Converts the provided bits into a byte
                               ⬑ Variadics are heavily used throughout the library =)

        NOTE: To also provides numerous convenience methods for the sub/super-byte types, though they
              are rarely utilized.

        tiny.To.Byte(0,1,1) // Converts the provided bits into a byte
        tiny.To.Morsel(bits) // Converts the provided bits into a 6-bit number

In addition, `tiny` is a fully-featured _binary calculator!_  While it may not be as featured as a higher-level
calculator, wherever any functionality is _missing_ `tiny` provides bridges into `math/big` to fill in the gaps.
That being said, from my current vantage point the below operations _should be_ all that you'd need for pretty
much any _algorithm_ you could dream up.  Since binary cannot implicitly distinguish the _sign_ from the _value,_ 
all operations are absolute and the sign is _contextually provided_ alongside -

        let 𝑎 and 𝑏 = Known Phrases

        𝑎.Add(𝑏)          = 𝑐
        𝑎.Minus(𝑏)        = 𝑐, ±
        𝑎.Times(𝑏)        = 𝑐
        𝑎.ToThePowerOf(𝑏) = 𝑐
        𝑎.DividedBy(𝑏)    = 𝑐
        𝑎.Modulo(𝑏)       = 𝑐

The bit width of the result is dictated by either growth from the operation, or held to the bit width of the 
longer operand and left padded with zeros.  Phrases also offer some general operations for width management -

    𝑎.ToNumericForm() // Trims off any leading zeros
    𝑎.PadRightToLength(𝑛) // Pads the right of the phrase to the provided length with zeros
    𝑎.PadRightToLength(𝑛, 1) // Pads the right of the phrase to the provided length with ones
    𝑎.PadLeftToLength(𝑛) // Pads the left of the phrase to the provided length with zeros
    𝑎.PadLeftToLength(𝑛, 1) // Pads the left of the phrase to the provided length with ones

Lastly, all phrases are able to perform _logic gate operations_ -

      Methodical  | Programmatic | Logical
        𝑎.NOT()   |      !𝑎      |   ¬𝑎  
        𝑎.AND(𝑏)  |     𝑎 & 𝑏    |  𝑎 ∧ 𝑏
        𝑎.OR(𝑏)   |     𝑎 | 𝑏    |  𝑎 ∨ 𝑏
        𝑎.XOR(𝑏)  |     𝑎 ^ 𝑏    |  𝑎 ⊕ 𝑏
        𝑎.NAND(𝑏) |   ^(𝑎 & 𝑏)   |  𝑎 ↑ 𝑏
        𝑎.NOR(𝑏)  |   ^(𝑎 | 𝑏)   |  𝑎 ↓ 𝑏
        𝑎.XNOR(𝑏) |   ^(𝑎 ^ 𝑏)   |  𝑎 ⊙ 𝑏

You'll get much more familiar with the intricacies of `tiny` as we progress, but the above should give you a
quick primer of the general idea I had when creating it.