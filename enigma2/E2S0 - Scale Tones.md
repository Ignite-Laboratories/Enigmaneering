# `E2S0 - Scale Tones`
### `Alex Petz, Ignite Laboratories, February 2026`

---

A diminishment has a mathematical formulation, but it's so much simpler than that because _**it's just repeating
a pattern of bits across another index.**_

To highlight that, let's take a look at how to diminish an 11-bit index with a 3-bit diminishment interval.

The most important aspect I'd like you to notice is the _visual_ output of "patterned bits" that
can _**asymmetrically**_ tile the space.

    " 3 Bit Diminishment of an 11 bit Index "

    let 𝑛 = The index width
    let 𝑤 = The pattern width
    let 𝑝 = The pattern value
    let  𝑆𝑦𝑛(𝑛, 𝑤, 𝑝) ↦ ⌊(2ⁿ / (2ʷ - 1)) * 𝑝⌋
    let 𝑑𝑆𝑦𝑛(𝑛, 𝑤, 𝑝) ↦ 𝑣(𝑛, 𝑤, 𝑝) - 𝑣(𝑛, 𝑤, 𝑚𝑎𝑥(𝑝 - 1, 0))
    where 𝑚𝑎𝑥(𝑎, 𝑏) returns the larger of 𝑎 and 𝑏 
 
              ⬐ Intervals             ⬐ Synthesized Potentials
              𝑝                 𝑆𝑦𝑛(11,3,𝑝)                       ⬐𝑑𝑆𝑦𝑛(11,3,𝑝)  
      (0) | 0 0 0 |   | 0 0 0   0 0 0   0 0 0   0 0 | (   0  ) + 292
      (1) | 0 0 1 |   | 0 0 1   0 0 1   0 0 1   0 0 | (  292 ) + 293
      (2) | 0 1 0 |   | 0 1 0   0 1 0   0 1 0   0 1 | (  585 ) + 292
      (3) | 0 1 1 |   | 0 1 1   0 1 1   0 1 1   0 1 | (  877 ) + 293
      (4) | 1 0 0 |   | 1 0 0   1 0 0   1 0 0   1 0 | ( 1170 ) + 292
      (5) | 1 0 1 |   | 1 0 1   1 0 1   1 0 1   1 0 | ( 1462 ) + 293
      (6) | 1 1 0 |   | 1 1 0   1 1 0   1 1 0   1 1 | ( 1755 ) + 292
      (7) | 1 1 1 |   | 1 1 1   1 1 1   1 1 1   1 1 | ( 2047 )
          |←  𝑤  →|   |←              𝑛            →|     ⬑ Potential Values
              ⬑ 3                 11 ⬏

Here, 3-bits (a binary "note") provide us with _8_ potential scale tones (an octave) to repeat across any index.

Fidelity space is unique because the plot of "equal-width sub-measurements ↦ ordinal index position"
generates a _waveform._

[insert example waveform here]

If the target index's bit-width isn't divisible by the diminishment's bit width, a quasi-periodic waveform 
will emerge from the interval when plotted against the target.

[insert quasi-periodic waveform atop of the last]

As such, I theorize that music theory can best approximate a waveform in fidelity space - just as a curious 
child can pick a tune out by ear on a piano.  Diminishments can be ordered to find the "chord progression" of 
a file, giving a "close enough" approximation of the target waveform.  In doing so, the magnitude to reach the
target from the diminishment interval will decrease enough for ideal midpointing.

Humanity has _eons_ of research invested in describing how **sounds** best fit against one another, making it a 
perfect toolkit for performing waveform approximation =)

    tl;dr - music theory can be used to approximate enough of a file that you'd only need to encode the delta needed to reach it