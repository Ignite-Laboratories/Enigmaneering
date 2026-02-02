# `E1S0 - I Heard There Was A Secret Chord`
### `Alex Petz, Ignite Laboratories, January 2026`

---

Before we can _"perform"_ a file, we get to learn how to play some notes.

An _index_ can define a range of scalar patterns provided by each potential's _individual bits_ - making it represent
the scale tones with which to "diminish" a larger index into a less resolute approximation.

Let's take a look at how to diminish an 11-bit index with a 3-bit pattern.

The most important aspect I'd like you to notice is the _visual_ output of "patterned bits" that
can asymmetrically tile the space.

    " 3 Bit Diminishment of an 11 bit Index "

    let 𝑛 = The index width
    let 𝑤 = The pattern width
    let 𝑝 = The pattern value
    let  𝑣(𝑛, 𝑤, 𝑝) ↦ ⌊(2ⁿ / (2ʷ - 1)) * 𝑝⌋
    let 𝑑𝑣(𝑛, 𝑤, 𝑝) ↦ 𝑣(𝑛, 𝑤, 𝑝) - 𝑣(𝑛, 𝑤, 𝑚𝑎𝑥(𝑝 - 1, 0))
    where 𝑚𝑎𝑥(𝑎, 𝑏) returns the larger of 𝑎 and 𝑏 
 
                                      ⬐ Synthesized Potentials
              𝑝                   𝑣(11,3,𝑝)                       ⬐𝑑𝑣(11,3,𝑝)  
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

3 bits (a binary "note") provides 8 scale tones (an octave) to repeat across any index.

Fidelity space is unique because the plot of "equal-width sub-measurements ↦ ordinal index position"
generates a _waveform._ Mismatched bit-widths between the diminishment interval and target index also causes
a quasi-periodic waveform to emerge from the binary note, when plotted.

[insert example plot here]

As such, I theorize that music theory can best approximate a waveform in fidelity space 
just as a curious child can pick a tune out by ear on a piano.  Numbers can be ordered to
find the "chord progression" of a file, giving a "close enough" approximation of the potential
to cross a useful threshold regarding its magnitude's bit-width.  

Humanity has _eons_ of research invested in the primitives to describe how **sounds** best fit 
against one another, making it a perfect toolkit for performing waveform approximation =)