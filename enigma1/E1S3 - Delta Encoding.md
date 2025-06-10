# `E1S1 - Delta Encoding`
## `A.K.A. - "Middle Out Compression"`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### What if you read binary from the top?
Every great project starts with a simple question.  Let's explore this one a bit!

Binary gives us a few unique qualities that most numbers _don't._  For instance, the bit width (which we'll 
call 𝑛) directly defines the absolute largest value that number could _not_ be: 2ⁿ. It also gives us the 
smallest value it _could_ be: 2ⁿ-¹.

The value exists _between_ those boundaries.

Let's call 2ⁿ-¹ the _light_ value, and (2ⁿ)-1 the _dark_ value.  (I apologize, superscript seems to have difficulties
with the '-' character)

         2ⁿ [ 1 0 0 0 0 0 0 0 0 0 ] (512)[256] <- Upper
     (2ⁿ)-1   [ 1 1 1 1 1 1 1 1 1 ] (511)[255] <- Dark
              [ 1 1 1 0 0 0 0 0 0 ] (448)[192] <- Upper Quarter
              [ 1 1 0 1 0 0 1 0 1 ] (421)[165] <- Target
              [ 1 1 0 0 0 0 0 0 0 ] (384)[128] <- Mid
              [ 1 0 1 0 0 0 0 0 0 ] (320)[64]  <- Lower Quarter
       2ⁿ-¹   [ 1 0 0 0 0 0 0 0 0 ] (256)[0]   <- Light (or Lower)

The target _clearly_ exists only 37 above the mid point, but we still _store_ it as 421 above _0!_

Why!?

Well, '37' doesn't tell us anything meaningful unless we know _where_ to apply it to.
This exists 37 steps _above_ the mid point between the powers of two that bound the target in, and
there are only _four_ different directions we could even walk:

    Up from the lower bound
    Down from the mid point
    Up from the mid point
    Down from the upper bound

Each of those directions represents a _quarter_ of the address space, which is coincidentally _two bits_
of information.  This allows us to start building an encoding scheme:

                   ⬐ The remainder bits
    [ ⁰⁄₁ ⁰⁄₁ ] [ ⁰⁄₁ ... ]
         ⬑ The focus crumb key
    
    Key | Meaning
     00 | The remainder is read as up from the lower bound
     01 | The remainder is read as down from the mid-point
     10 | The remainder is read as up from the mid-point
     11 | The remainder is read as down from the upper bound

So, let's go back to our previous example and write out the encoded value of 421:

    [ 1 1 0 1 0 0 1 0 1 ] <- Target Measurement
    | 1 0 - 1 0 0 1 0 1 | <- Encoded Phrase

Immediately, we have gained a _single bit_ of reduction!  This is _fantastic_ - but don't start quarter splitting
every byte you find quite yet: _the length has changed!_  Your next measurement would not be readable because you
wouldn't know _when_ it starts!  However, this particular design has an exploit: it works _better_ at scale =)

This particular solution simply shows how to delta encode _any_ length binary information and spits out the resulting
binary waveforms.  I want you to grasp that this is a very _cyclic_ process, and that we are taking advantage of this
particular aspect of numbers.

