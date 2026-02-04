# `E1S2 - I Heard There Was a Secret Chord`
### `Alex Petz, Ignite Laboratories, February 2026`

---

Now, I'd like to touch on how to play "scale tones" across an index using diminishment.  If it's helpful, you can
consider this analogous to signal theory - as repeated diminishments are layering waveforms on top of one another.  

You have two axes of freedom - tones can be applied _sequentially_ or _concurrently._  

**Sequential** application breaks the target index into similarly sized bit-widths where each tone is applied locally.  If 
an even bit-width division can't be made, the bit-widths should round to the nearest integer and the final instance may 
be smaller.

This is the _time domain._

**Concurrent** application recursively spreads each tone across the resolution identified by the last - more densely
approximating the target potential.

This is the _frequency domain._

These two modes of operation can be composed together to generate a rich set of tones I call the target's "binary
overtones."  These harmonics can be very advantageous!  This is because the ratio of bit-widths between the diminishment
interval and target index is analogous the _frequency_ of a musical tone.  This means a "perfect fifth" (1.5:1) or "minor
third" (1.2:1) can _**mathematically**_ be represented as the number of times a pattern repeats across
the target index's bit-width.

Effectively, this means that a less-wide scale tone is a _"higher note"_ in the index, as it "oscillates" at a higher frequency.

Naturally - this also means a 1-bit interval is the _highest possible binary frequency._

