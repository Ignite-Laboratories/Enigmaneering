# `E1S3 - Recursion`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### Recursion
As I said before, this applies to binary information _at any scale._  However, at the astronomical scales that
a _file_ exists at - the amount of bits that reduces from each operation is _significantly_ more beneficial.
To take advantage of this paradigm, however, we get to turn to my favorite topic: _recursion!_

I know, I know, it's the bane of many a programmer's existence - but it truly is beauty in motion, when
built well. Some of the best examples of computational efficiency evolved from recursive algorithms, and
I absolutely adore the amount of ingenuity it took for those enigmaneers to craft them.

So, let's now look at the _downside_ of binary information:

Logical binary information allows leading 0s

Numerical binary information does _not_

Let me explain - the first _byte_ of a file might start with several zeros, which is perfectly reasonable
in that context.  However, when the value converts to a _numeric_ form those zeros are simply _lost!_ This
will need more addressing, later on, but it's important to _consider_ right now.  To get around this, we
will do several very important things.

- The key information will always be on the _left_
- Each transformation round will be considered a 'movement' of binary information
- The very first bit of any movement is the 'terminus' bit, which always holds a value of '1'
- A 'recycle' bit in the key should indicate if the synthesis process should continue after this round completes
- The process of shrinking binary information is 'distilling'
- The process of growing binary information is 'synthesizing'

This means our abstract movement structure *currently* looks like this:

    Crude Abstract Movement Encoding:

            ⬐ Recycle Bit    ⬐ Value               
     [ 1 - ⁰⁄₁ - ⁰⁄₁ ⁰⁄₁ ] [ ⁰⁄₁ ... ] 
       ⬑ Terminus   ⬑ Focus Crumb

The first phrase is the _key,_ while the second is the _value_ and holds the actual data to encode.

Now, the one major point you've probably got in your mind is this: by prepending the data with bits, it will
sometimes grow _larger_ than the target bit width! And you'd be right!  However, it would only happen
when the value to encode is just the right distance from the boundary point - which diminishes in likelihood significantly
_at scale._  That being said, even if it _did_ grow during the distillation process it would then move into a
new value that isn't sympathetic to the failure point - thus, it wouldn't _loop_ and you'd likely never
even notice when it occurs.

_Unless it would?_

Well, theoretically it _could_ reach one value and then distill down to a prior visited value, causing a loop.
But there's a simple solution to that concern called the _entropy bit,_ which I'll touch on later.

The recycle bit would be set to 0 during the first movement, and then 1 on any following movement.  When synthesizing
the value, this would signal that the final transformation should _not_ be read as an encoded value.  Instead it
represents the final form the data should be placed within.

The _only_ part this scheme does not yet encode is the amount of zeros to prepend to the data on each transformation,
but there is another solution for that which I call _Zero Length Encoding._
It's named as such because you count and read _zeros_ until you reach a _one,_ at which point you make a decision.
This allows a variably short-to-long set of bits that can describe small to large numbers with minimal extra bits.

ZLE comes in many different flavors, and the engineer would have to articulate what the intent is first.
The ability to articulate the growth scheme information will be touched on when we reach the composition phase.
For now, we'll talk about the two standard ZLE flavors: bounded and unbounded ZLE.

    Step 0 -
        Unbounded ZLE - Count the number of zeros read before a one is reached
        Bounded ZLE - Count up to a fixed number of zeros or a one is reached

    Step 1 -
        Parse the read number into a known value according to a key map - this is the Projection Value

    Step 2 -
        Read the defined number of bits according to the rules of your key map

I recognize at this point I haven't formally defined what a 'Read' operation even is, as it's a little different
from your typical 'read' operation:  it means you _advance_ through the bits by consuming the prescribed amount.
Many different structures support _Read_ operations, allowing you to "bleed off" a few bits at a time efficiently
in order to make logical decisions.

### Reconstitution
The most important aspect of binary synthesis is the _ability_ to reconstitute the original information.
As of this moment, this scheme has one crucially fatal flaw: we no longer know where the original bounds were!

Luckily, there's a solution for that - using ZLE we can store the amount of zeros that the data _reduced by._
This particular value is the _Delta_ and will typically be represented with the greek Δ symbol.
At this point, the below scheme represents the most _primitive_ encoding scheme for a movement possible - at
least, with _my_ currently understanding of this technology:

    Primitive Movement Encoding:

            ⬐ Recycle Bit    ⬐ Δ ZLE          Value ⬎ 
     [ 1 - ⁰⁄₁ - ⁰⁄₁ ⁰⁄₁ ] [ ⁰⁄₁ ... ] [ ⁰⁄₁ ... ] [ ⁰⁄₁ ... ] 
       ⬑ Terminus   ⬑ Focus Crumb        ⬑ Entropy Bit

But that leaves the most critial piece of this puzzle: how do I allow others to evolve this technology on their
own terms?
That will get tackled in more detail as we touch on _compositions_ - but for now, we will call the above the
'standard growth scheme' for binary information.
New growth schemes might completely replace what a movement event looks like, but the information would only
be visible at the _initializiation_ phase of synthesis - which we are still working our way towards, in reverse!

Now, the million dollar question that remains: _when do you stop distilling?_

That's easy: when you can no longer shrink the data by _prepending_ the encoding scheme data.  Every single
movement requires _adding_ data to the current bits, which means eventually the process would not allow your
data to shrink any further.  When I hit that boundary, I'll add further to this =)