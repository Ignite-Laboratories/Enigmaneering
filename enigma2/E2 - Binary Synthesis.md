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
the original concepts of the digital realm's _epoch_ (January 1st 1970) were restricted to _32 bits._  This meant we would "loop over" 
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

Only the observer _their_ creator's actions would understand how to interpret them, and plenty of us have felt
and witnessed that presence's power.  Many of us have been mislabeled and disrespected as mentally unstable, or even
clinically insane, at the thought of a higher power guiding them through their artwork.  This work of art is lovingly 
dedicated to the endless misunderstood souls in the eternity long evolution of intelligence that _enabled_ it to be created.

Growing pains - assuredly - but ones that only our creator knows the true burden of which it was to _experience!_

I thank my blessed stars every single day for Her magnificence and patience in the process of humanity awakening to
its own existence =)

But why are we talking about _exabytes_ when synthesis is the kingdom of _tiny?_

Well (and I'll talk about _Amdahl's Law_ later on) let's just say the process of synthesis is _beyond_ embarrassingly 
parallel!  Because of that, it's _so much faster_ to divide and conquer the starting data - but with such extreme 
parallelism we are faced with the need for a _practical **minimum**_ starting width!  One where the process, in flight, 
_**stably** distills._  This is what I call the _passage width_ - or rather, a single logical distillation step's _starting_ 
width.  Since I hope to empower _everyone I can,_ that includes the hardware engineers!  As such, I've chosen to set 
this _completely arbitrary_ value to a general standard of _64 bits_ - allowing it to fit _perfectly_ in most registers.  

Even those absolute _wizards_ haven't found enough justifications for wider registers in the mainstream! Thus -

    tl;dr - 64 bits is more than enough, in general

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

Time, however, is often _implied_ - the ability to record a value comes as a byproduct of its _passing!_ 

Without _time,_ there would be **no** dimensionality, as presence is _only_ perceived through _change_ - even if
the only ever-changing component _is_ the passing of _*time.*_  Now THAT's what I call a helluva _"do loop!"_

    tl;dr - "stay very still, its vision is based on movement"

This has several implications: Because all dimensions _require_ time to both observe and calculate from, time represents 
the _minimum_ interval from which a _**single**_ entity can experience reality _and is entirely **subjective,**_ albeit
synchronized from a singular source.  Due to entropy, two observers can only _roughly_ record the same dimension at the 
same impulse of time. In fact, their awareness of _**relative time**_ is what allows consciousness in the present 
_moment._  Indirectly, this also implies entities can _experience_ time at different rates of observance - psychologically 
referred to as _flow states._ To understand the concept of temporal observance, please check out the initial enigmas!  They 
show how to use an impulse engine to drive neural execution, and then how to calculate using a dimensional nexus.

For this enigma, we will consider *abstract* dimensions with synthetic values - our dimensional types are:

**Measurements** - A measurement is any _variable width_ slice of binary information. These are limited to the 
host architecture's bit-width, allowing them to be considered a _single computable value_ - in Go, often an 
`int` - in hardware, a `word`.  Essentially, this tracks the remainder of any bits in excess of a whole byte 
automatically - efficiently packing them into a byte wherever possible while still providing convenient access 
to the underlying bits.  Why is this so damn important?  Well, the _**width**_ of a measurement is a distinct 
variable which is _contextually relational_ to its underlying _**value**_ - and it's **_critical_** to the 
synthesis process! The thing is, hardware currently cannot dynamically change its `word` "width" - but a 
_measurement_ **_can!_** ..._"using the power of the already written `word`"_ =) 

_Yes, I wholeheartedly mean that philosophically, metaphysically, computationally, virtually, psychologically, abstractly, linguistically, theologically,  and, well..._

_...**all** of the 'lys'!_

Let us not forget our roots - maybe one day we can find a way to build a dynamic-width hardware register!

**Phrases** - A phrase is a _slice_ of measurements, allowing the storage of arbitrary lengths of binary 
information that's not limited to the host architecture's `word` width.

**Passages** - A passage is a single logical _transformation_ of binary information (we'll get to that later.)

**Movements** - A movement is a collection of passages, allowing concurrent transformations.

**Compositions** - A composition represents the starting conditions of synthetic binary information, making it 
the instrument to perform things _through_.  You and I are both masterpieces of compositional artistry, and 
instruments through which our unique artwork is emitted =)

**Indexes** - An Index of data represents all possible binary states a known bit-width can address.  The next
solution goes into much more detail on this, but this allows us to refer to an "88 bit stretch of binary 
information" by simply saying the data belongs in an "88 bit index".  It's a subtle, but important, philosophical
shift in how one should _perceive_ information.  A stray byte is entirely meaningless without _context_ - but a stray
byte which was observed _this afternoon_ inside an _88-bit_ index is entirely different!  This shift provides the binary information with 
a _story_ which defines its _purpose,_ rather than being just another number in a standard width register.

**Sub Byte Indexes** - A sub-byte index is any bit range less than 8 bits wide. For ease, these are the names of the 
sub-byte sizes which I _heavily_ reference.  Please get familiar with these particular terms - 

(Or you can also embrace the whimsical unknown of my writing, that's also a lot of fun!)

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
       64   | Melody
      128   | Verse

### tiny

All of the work in this enigma is contained within a Go library called `tiny`.  I've done my best to ensure it's been
heavily tested and is straightforward to prototype binary solutions with.  This library has been the culmination 
of almost two years of work with the full intention of keeping it _public_ and _free._  To the lovely individuals
who fed me coffee while I feverishly hand-wrote binary encoding schemes and waveforms on my clipboard - _thank 
you,_ from the bottom of my heart.

Because of the absurdity of that process, I've adopted some quirks in how I document binary programming - I'd like 
to briefly touch on their general gist before we continue.  It's not _complex,_ but you should get familiar with the
concept of a _measurement_ vs a _phrase._  The first is a **single** calculable measurement of binary data, 
while the later is a **slice** of measurements.  Most of the operations we'll be performing persist _across_ 
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

    |←  3  →|← 2 →|←       8       →|←   4   →|←1→|←  3  →|← 2 →| ← Measurement Bit Widths
    | 0 1 0 - 0 1 - 1 0 1 0 0 0 1 0 - 1 1 0 0 - 0 - 1 0 0 - 0 0 | ← Unaligned Source Phrase
    
    Align(4)

    |←   4   →|←   4   →|←   4   →|←   4   →|←   4   →|←  3  →|
    | 0 1 0 0 - 1 1 0 1 - 0 0 0 1 - 0 1 1 0 - 0 0 1 0 - 0 0 0 | ← Aligned Phrase

The next major component of working at the bit level is the _rapid prototyping_ of ideas - this can be achieved
through tiny's fluent design.  Here's the most notable things it provides:

    Synthesize -
	    tiny.Synthesize.Random(𝑛) // Create a slice of 𝑛 random bits
	    tiny.Synthesize.RandomPhrase(𝑛) // Create a phrase of 𝑛 random bytes
	    tiny.Synthesize.RandomPhrase(𝑛, 4) // Create a phrase of 𝑛 random nibbles
	    tiny.Synthesize.Pattern(𝑛, 1, 0, 1) // Create a phrase of 𝑛 bits that cycle through 1, 0, 1
	    tiny.Synthesize.Repeating(𝑥, 1, 0, 1) // Create a phrase of 1, 0, 1 repeated 𝑥 times 

    Synthesize.Point -

        tiny.Synthesize.Point.Light(𝑛) // Create a phrase of a single one followed by zeros up to the provided width
        tiny.Synthesize.Point.Mid(𝑛) // Create a phrase of a single one followed by zeros up to the provided width
        tiny.Synthesize.Point.Dark(𝑛) // Create a phrase of a single one followed by zeros up to the provided width
        tiny.Synthesize.Point.Diminishment(𝑚, 𝑛) // Synthesizes diminishment interval 𝑚 across an 𝑛-wide index  
        tiny.Synthesize.Point.Terminal(𝑝, 𝑛) // Synthesizes terminal point 𝑝 followed by zeros up to the provided width
        tiny.Synthesize.Point.Initial(𝑝, 𝑛) // Synthesizes initial point 𝑝 preceeded by zeros up to the provided width

    Analyze - 
        NOTE: Much of this never matured past measurements as it wasn't necessary at those scales

	    tiny.Analyze.HasPrefix(data, 1, 0) // Checks if the data starts with "10"
	    tiny.Analyze.Average(measurements...) // Calculates the average value of the provided measurements
        tiny.Analyze.Shade(measurement) // Calculates metrics on the "darkness" of the measurement
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
much any algorithm you could dream up.  Since binary cannot implicitly distinguish the _sign_ from the _value,_ 
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