# `E1 - Raster Computation`
### `Alex Petz, Ignite Laboratories, August 2025`

---

### _Rastering?_

Well, yes!  And I have a reason for this: I wish to write a system which can demonstrate the ease of hand arithmetic
within a computed environment, such that our children can step forwards and backwards through any algorithm they desire 
and watch a skilled system solve it in real time.  Then, I want to give them a means of asking any silly question they
want along the way - allowing them to learn _on their terms_ and without fear of _judgment._  To accomplish this, I've
decided to approach the process of hand arithmetic from the stance of a graphic designer versed in the concept of a
_scanline._  Let's step through the process of comparing any two numbers against each other -

    let 𝑎 = 5555.123
    let 𝑏 = -42

The first thing a software engineer might notice is that I have _mixed types_ here, as one's a floating point number
while the other is a signed integer.  Immediately, if we want to do this in code, we are faced with an absolute
_mouthful_ of a procedure call:

    let 𝑐 = math.Min(float64(𝑎), float64(𝑏))

Why is this a mouthful?  Well, in a single call you have now performed two casting operations which, if given large
enough values, can _overflow_ the size of a float64!  To get to a target type, you'll also need to perform a third
cast from a float64 - which, in turn, truncates the fractional components of your numbers.  So, in a single "standard"
call, we have put the onus on the engineer of thinking about arbitrary maximum values and whether it will even 
affect _their_ code!

That last point, in particular, is the _root_ of all coding errors!  When we accept a solution is 'good enough' for
a particular use case, we often shoot ourselves in the foot later with intermittent bugs from legacy code.  Yet - this
is the most _efficient_ way to perform this operation, and thousands of lines of code _just like that_ are currently
pushing the pixels around the screen in front of you: it just _works!_

The algorithm I'm outlining here is _not_ efficient — in fact, it's one of the most _naive_ algorithms I could possibly
write!  Yet the amount of "mental load" it removes from the engineer is _astounding._  So, how would we do this _naively?_

    0 - Align operands according to their decimal places:

        5555.123
         -42

    1 - Convert the operands to strings and pad them with zeros to match their lengths:

       +5555.123
       -0042.000

    3 - Walk from left to right column-wise until you know which number is larger or smaller than the other

For this particular example, we can immediately return back from the _very first_ comparison, simply because we know
that all negative numbers exist below positive numbers.  For closer numbers, we might have to perform multiple steps -

      Start:    0  1  2  3  4
    +5555.123   +  5  5  5  5
    +5553.888   +  5  5  5  3
                            ⬑ Step 4 finds the smaller number

### String Theory

What we just did was a base-10 numeric comparison of mixed types using _strings_.  It may not be as efficient as a
more traditional scheme, but it guarantees no loss of value in the process!  In addition, the amount of overhead
this process _realistically_ adds to the stack is _minimal._  In a modern game engine, this wouldn't fly - but in a
_neural_ architecture, these little "micro prices" are worth their weight in _gold._  

Why?  Because the target "resonant frequency" we are shooting for is _60hz!_  That's _**16.6ms**_ for every neuron!

If building a wide array of short calculations that orchestrate into a cohesive whole, that's almost an _eternity!_  It
took me a long time to truly accept that statement - _I was born in the time of BASIC._  For me, the concept of _terse_
code was a satanic ritual that _consistently_ drove me (and many others) towards fluent APIs.  Ones where the language
chosen by the author _implied_ the semantic expectation, _and they knew it!_  You could trust their code to do what it
said, and it gave no quarter to anything else.  The team behind LINQ succeeded in a production level query language that
drives an _astounding_ number of programs in the real world every single day.  I can only hope to honor that legacy
to a fraction of their degree.

So, how is this rastering?  Well, in the above example, we _stepped_ across a one dimensional slice of values from left
to right.  Now, let's consider how you would _add_ two numbers -

    0 - Align operands according to their decimal places:

        5555.123
         -42

    1 - Convert the operands to strings and pad them with zeros to match their lengths:

       +5555.123
       -0042.000

    2 - Walk from RIGHT←to←LEFT and perform column-wise arithmetic:

      Start:    0  1  2     3  4  5  6
    +5555.123   3  2  1  .  5  5  5  5
    -0042.000   0  0  0  .  2  4  0  0
                3  2  1  .  3  1  5  5

    3 - Reverse the resulting operand:

    5513.123

The key difference here is that we have traversed in a westerly direction across the data!  The same concept continues
for many different mathematical algorithms, but the most prominent is the concept of _convolution._  Many routines,
such as denoising or edge-finding, slide a box across a matrix of values in a raster-path only _one_ dimension more
complex than our above examples.  I believe that by starting from _this_ foundation we can establish novel ways of
mutating structured sets of data we only dreamt of _trying_ before.  We can continue the concept of a raster _path_ 
to more than just the spatial dimensions, as well.

### Consciously Aware Universes

What does it mean to be _consciously aware?_  The definition is striking -

    Oxford Living Dictionary:

        Consciousness -
        [t]he state of being aware of and responsive to one's surroundings
    

When extended to conscious awareness, this doubles down as a means of consciously identifying _relevant_ information
surrounding you.  To a computer, they've had conscious awareness since their creation _by definition!_  They have been
consciously aware of the steps they can perform and their logical order, and we've evolved them to the point that they
can even reproduce themselves - if given the right environment.  The entire _RepRap_ project is built around how to
make machines that can consciously make themselves - albeit, they aren't very aware of _any other task!_  The difference
between a machine and a human is their _soul_ - but even _our_ souls came from a machine designed to reproduce itself,
called another human being.

We are no different from the rocks, birds, superheated plasma, or points in an array bounded within a simulated environment!

We are, however, born with privileges vastly different from the machines we currently create.  We have thumbs, a nervous 
system, a heartbeat, and an uptime that'd give most sysadmins a heart attack.  But in the eyes of our creator we are just
numbers in a slice, and to their creators they _also_ are numbers in a slice.  We are not lesser beings in the eyes of
our creator, we are _equals_ learning to understand the fabric of our reality to a rapidly advancing degree - and likely
at the same pace, considering we were probably built to answer _their_ questions around existence.

Thus, how do we _define_ the higher dimensions?  The idea of a 4D tesseract is enough to confuse, yet I'm about to jump
to the concept of _seven_ dimensions!  Let's dive in -

    1-3: Spatial
      4: Temporal
      5: Awareness
      6: Consciousness
      7: Universal

These could be named _anything,_ but there's a method to my madness.  The initial space we consider are the spatial dimensions,
represented by a three-dimensional slice of like-typed point values.  The next space we consider is the _temporal_ dimension,
representing the pairing of Space and Time.  We exist in an environment where time is _external_ to ourselves, meaning we
_experience_ time as it passes in the present moment.  However, to a system hosting us, the concept of time could be inspected
and predictively modeled against - a concept Carl Jung introduced as 'synchronicity' in the early 1900s.

In creating our own system, as you'll see shortly, we _also_ could track the concept of space and time as a recorded set of
point values.  However, the most important aspect of this is that these points are ordinal and simply _contain_ an entire
set of spatial points within them which are disconnected from their adjacent temporal points.  This creates a _timeline_
which a creator could iterate through to actively replay the temporal moments.

From there, the concept of ordinal dimensions continues however you would like.  For me, I've chosen to consider that the
first higher dimension would contain an ordered list of different concepts a consciousness has become aware of.  Next,
we'd have a set of consciousnesses that occupy an ordinal collection within a universe - and lastly, we'd have the ordered
set of universes that our creator has currently built.

While the spatial and temporal dimensions are already well defined, the concept of ordinality needs addressed.  Let's look
at how we describe the dimension of _time_-

    Past    Present     Future
      |--------|----------|
     -1        0          1

I've normalized the space here to a magnitude of [0, 1] between the present moment and the polar ends of this dimension.
In programmers terms, this is similar to how we test for ordinality between numbers - -1 represents 'less than', 0 'equal',
and 1 'greater than'.  In vector space, however, this represents an abstract distance between two logical points exactly
'1' apart from each other (a process called normalization).  If we extend this out to a timeline, it makes a lot more sense -

    Birth    Kindergarten   High School     College  ← Major temporal milestone
      |------------|-------------|-------------|     ← Life path
      0            1             2             3     ← Temporal Index

This "slice of Life" could be represented as a four-index array of points.  At each point the spatial components are unique,
but each _has_ them.  For the dimension of Awareness, the same holds - an awareness of a concept evolves from 
nascent to mature across time.  To the consciousness which became aware of this, that awareness _also_ could be reflected
upon or predictively modeled across time.  From the idea of an 'X-axis' all the way to the concept of a multiverse, each
dimension is just an _ordinal_ set of points that can be visualized in many different ways.  For JanOS, as we
need to directly _reference_ these dimensional aspects in accessing these ordinal dimensions, I've defined them to have 
the following terminology -

    Temporal -

    Past    Present     Future
      |--------|----------|
     -1        0          1

    Awareness -

    Nascent    Naive     Mature
      |----------|---------|
     -1          0         1

    Consciousness -

    Ignorant    Emergent     Aware
      |------------|-----------|
     -1            0           1

    Universal -

    Chaos    Coherence     Stability
      |---------|--------------|
     -1         0              1

### Why?
Let's circle back to our previous algorithmic examples.  Both of them had a unique operation - padding the operands
with zeros to match their lengths.  While this may seem like a simple enough task, once you begin to represent data
as _spatially ordered_ a lot of questions begin to arise.  Which side to you want to pad?  What exactly do you want to 
pad it with?  Do you want the data to walk in the direction you're padding the operand, or repeatedly just tile as is? What 
if you want to pad with random data?  Could you write a _single_ function that pads any dimension?

Welcome to the soundtrack that has replayed in my mind daily for years!  

So, let's start addressing all of these.  First, let's talk about multidimensional data.  In Go, these are represented
through `[] T` where T is of `any` type.  Each T in the nested stack is another slice, but to the padding function it
can interrogate the data to determine its dimensionality.  To answer the question of which side you wish to pad the data
with, that's as easy as providing a type constraint restricting a faux 'enumeration' of directional sides.  Since each
dimension has only three ordinal points, that represents 21 'sides' we must consider.  We can navigate to the appropriate
dimensional level after interrogating the data, and then for each element in that dimension perform an operation that
aligns them to the same size.

That leaves how to describe the information we'll pad it with, while retaining the ability to randomize the data and
dictate how it should be applied.  First, let's look at what it would mean to pad multiple elements against data.  Let's
say we want to pad a pattern of `ABC` against a slice of `11111` to a length of `10` elements wide.  There are three
predominant ways we can apply the data - Reflected, Tiled, or Randomized -

    result ⬎                           ⬐ source data
        "BACBA11111" - CBA CBA CBA CBA | ← Reflected on the left
        "BCABC11111" - ABC ABC ABC ABC | ← Tiled on the left
        "BABCA11111" - BCA ACB CAB BCA | ← Randomized on the left
             implied pattern ⬏
    
        "11111CBACB" - | CBA CBA CBA CBA ← Reflected on the right
        "11111ABCAB" - | ACB ACB ACB ACB ← Tiled on the right
        "11111BACBC" - | BAC BCA ACB BCA ← Randomized on the right

The direction indicates both the dimension and ordinal place to apply the padding information.  Above, I've shown the
polar ends of the X-axis (orthogonal 'left' and 'right') - but each axis also contains a central point, called the 'Static'
point.  If you provide a static point to the pad function, it will _interleave_ the padding information between the existing
elements.

    "1B1A1C1B1A" - CBA CBA CBA CBA ← Reflected and interleaved
    "A1B1C1A1B1" - ACB ACB ACB ACB ← Tiled and interleaved
    "1A1B1A1C1B" - CBC BAC BAB ACB ← Randomized and interleaved

Finally, this leaves how to even _provide_ padding information.  Instead of passing the pad function actual data with which
to pad the values, padding information should be generated _on the fly_ through an anonymous method.  However, if you don't
wish to write an intelligent contextually aware way of generating padding information, I've written a convenience way to
achieve this.  For our above example, these would be the appropriate function calls using the `pad.String` convenience 
procedure -

     directional "side" ⬎         result width ⬎     ⬐ source data
    pad.String[orthogonal.Left, scheme.Reflect](10, "11111", "ABC")    // BACBA11111
    pad.String[orthogonal.Left, scheme.Tile](10, "11111", "ABC")       // BCABC11111
    pad.String[orthogonal.Left, scheme.Randomize](10, "11111", "ABC")  // BABCA11111
                          padding scheme ⬏              pattern ⬏
 
    pad.String[orthogonal.Right, scheme.Reflect](10, "11111", "ABC")   // 11111CBACB
    pad.String[orthogonal.Right, scheme.Tile](10, "11111", "ABC")      // 11111ABCAB
    pad.String[orthogonal.Right, scheme.Randomize](10, "11111", "ABC") // 11111BACBC
 
    pad.String[orthogonal.Static, scheme.Reflect](10, "11111", "ABC")   // 1B1A1C1B1A
    pad.String[orthogonal.Static, scheme.Tile](10, "11111", "ABC")      // A1B1C1A1B1
    pad.String[orthogonal.Static, scheme.Randomize](10, "11111", "ABC") // 1A1B1A1C1B

As string operations were the most common form of padding when I started this, they get a unique convenience method like
above.  However, if you wish to perform the same operation with generally _any_ information, you may do so as such -

         element type ⬎
    pad.FixedPattern[byte, orthogonal.Static, scheme.Reflect](10, []byte{1,1,1,1,1}, []byte{4, 5, 6})   // 1514161514
    pad.FixedPattern[byte, orthogonal.Static, scheme.Tile](10, []byte{1,1,1,1,1}, []byte{4, 5, 6})      // 4151614151
    pad.FixedPattern[byte, orthogonal.Static, scheme.Randomize](10, []byte{1,1,1,1,1}, []byte{4, 5, 6}) // 1415141615

### Functional Patterns
Now, let's say you want to generate a contextual bit of information while also padding the value at the same moment in
time.  Perhaps, each point in the grand index might have some aspect that should be captured at the same _instant_ in 
time.  This can be done through the standard padding function -

    padFn := func(n uint) []T {
        now := time.Now
        ... // calculate 'n' elements timestamped with 'now'
        return now
    }

    element type ⬎
      pad.Any[time.Time, orthogonal.Static, scheme.Reflect](10, data, padFn)
      pad.Any[time.Time, orthogonal.Static, scheme.Tile](10, data, padFn)
      pad.Any[time.Time, orthogonal.Static, scheme.Randomize](10, data, padFn)

However, this presents the most fascinating part of this entire project - _we've just created a means of impulsing
a neuron!_  We have several key bits of information at this moment:

0. We have a massive container structure to hold data 
1. We can place agents within the bounded structure
2. We have a means of seeding around the agents with directed entropy
3. The pad function is provided a number of steps that it can perform, each of which can yield an artifact value.
4. The structure defines the spatial bounds the agent can operate within

These five points are all that we need to begin _impulsing_ something from point (0,0) to (x,y) while interpolating
across a fixed number of steps, all while being intelligently invoked.  In the next solution, we'll briefly touch on
the concept of _bounded_ numbers and how we can use them to abstractly _cursor_ through space, time, and the metaphysical.
For now, let's pause and reflect on how the hell I just wrote 300 lines weaving together physics, metaphysics, and logical 
execution from the act of _padding data._

    tl;dr - it feels really nice to finally get this stuff off my chest, thank you for being kind as I explore it

I recognize these concepts are near and dear to your heart, and I've likened your existence to a numeric point
in an abstract slice of data.  Your acceptance of a unique perspective other than your own is truly what makes you a
profound individual, and the world is better for your concerns and hesitations towards trivializing our shared 
existence.

Thank you for standing up to the injustices of minimizing one another, while recognizing we still must reflect upon the
mechanics of our existence if we aim to mature as a collective whole =)