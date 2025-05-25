# `E1S1 - Fuzzy Approximation`
### `Alex Petz, Ignite Laboratories, May 2025`

---

### Astronomically Large Numbers
Every single file in existence is an _Astronomically Large Number_ - represeneted in binary.  Because of this,
there are many different ways to reach a target number through mathematics - and I intend to exploit that to 
_algorithmically_ reach the target from a seed value.

To do so, we need to _distill_ or _reduce_ the number down to its seed value.  I'd like to try doing this
through synthesizing scaled approximations of the number using a few simple steps.

First, we need the absolute simplest way to get _close enough_ to our target value in a minimal number
of bits.  To do this, I'd like to establish a "practical infinity" unique to computational models: 2⁶⁴

Why?  Because at the current moment in time there is no reason humanity will ever need to create a
singular file that's roughly 2 exabytes long!  For all intents and purposes, we can treat this as the
upper boundary that addresses every unique file that humanity will ever _desire_ to create.

    tl;dr - I'm daring some cheeky bastard to eventually justify a ~2 exabyte file's existence

This means in 64 bits we can make the absolute fuzziest approximation of the target data, because
it's the minimum number of bits necessary to synthesize the closest value. Let's say that a file is 
`n` _bits_ long - that means its value exists between `2ⁿ⁻¹` and `2ⁿ-1`. `2ⁿ⁻¹` is considered the 
_Light_ binary representation of the target value, while `2ⁿ-1` is considered the _Dark_ representation.

For Example:

                    |---10 Bits Wide----|
               512:[ 1 0 0 0 0 0 0 0 0 0 ] <- Light (2⁹)
    Target ->  715:[ 1 0 1 1 0 0 1 0 1 1 ] <- Grey
              1023:[ 1 1 1 1 1 1 1 1 1 1 ] <- Dark (2¹⁰-1)

In the above case, that means we would synthesize _512_ to fuzzily reach the target value.

### The Timeline
From here, we will start representing the stored bits towards approximation as the _Timeline._  Before we start
_any_ approximation we have critical operation that first has to be performed: trimming leading zeros.

Data is stored in byte form, meaning the MSBs of the initial byte may or may not be 1 - for our approximation
to work, we need to trim any leading zeros off and store the amount removed.  Since the chances of seeing a 0
in the first byte of a file is minimal, we will store the value to trim in a _Note_ (3 bits); however, we will
treat the value `7` as a unique condition where an extra bit _should_ be read to indicate if the 8th bit of the
first byte was also trimmed. For example:


                             𝧘 Zero trim width
                         |   6   |                       
         Timeline ->     [ 1 1 0 ] -> Trim 0 zeros

                                    𝧘 Continuation value
                         |   7   | 0 |                       
         Timeline ->     [ 1 1 1 | 0 ] -> Trim 7 zeros

                         |   7   | 1 |                       
         Timeline ->     [ 1 1 1 | 1 ] -> Trim 8 zeros

The timeline is shown in _Phrase_ form above, where the `|` character represents a break between _Measurements_
in the phrase of binary information.  A measurement is simply a variable width slice of up to 32 bits, while a
phrase is merely a slice of measurements.  Both types provide lots of structure around working with this kind
of data, which we'll explore as we progress forward.

For our current contrived example, we'll assume there was no trimming to be had on the timeline.

### Hierarchical Scaling
Now, if we continue following the approximation pattern then we eventually would re-create the file - but we'd be 
storing the steps _extremely inefficiently!_  While 64 bits to kick off the value is an important first step, we
need to start making some much tighter approximations.  To do so, we'll start halving our workable bit-width
and using that restriction to narrow down our search.

    tl;dr - subdivide the power of 2 range by the number of points your current width allows

What does that mean?  Well, let's take our above example and continue walking through the process, but we'll
perform it at an appropriate "scale".  

_NOTE: When I say "at scale" I mean that we will start with an initial approximation of 4 bits wide, rather than 64._

    Step 0:

           Target -> 715:[ 1 0 1 1 0 0 1 0 1 1 ] <- Grey
    Approximation -> 512:[ 1 0 0 0 0 0 0 0 0 0 ] <- Light (2⁹)
            Delta -> 203:[     1 1 0 0 1 0 1 1 ] <- Grey

                                       𝧘 Light bit width
                         |   0   |    9    |                          
         Timeline ->     [ 0 0 0 | 1 0 0 1 ]


Now, let's use a _crumb_ to approximate the current _Delta_ value (203).  The delta value exists between `2⁷`
and `2⁸-1`, yielding an address space of `127` which we we'll subdivide using a _Crumb_ (2 bits).

    Step 1:

                     128:[ 1 0 0 0 0 0 0 0 ] <- Light (2⁷)
           Source -> 203:[ 1 1 0 0 1 0 1 1 ] <- Grey
                     255:[ 1 1 1 1 1 1 1 1 ] <- Dark (2⁸-1)

            Target -> 75:[   1 0 0 1 0 1 1 ] <- Source - 2⁷
                            { Range 0-127 }
     Approximation ->  2:[             0 1 ] <- Third quadrant of Range (((Range+1)/4)*2)
             Delta -> 11:[         1 0 1 1 ] <- Target - Quadrant (64)

                         |   0   |    9    |  2  |                        
         Timeline ->     [ 0 0 0 | 1 0 0 1 | 0 1 ]