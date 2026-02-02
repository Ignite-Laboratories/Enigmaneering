# `E1S0 - I Heard There Was A Secret Chord`
### `Alex Petz, Ignite Laboratories, January 2026`

---

Before we can _"perform"_ a file, we get to learn how to play some notes.

A _diminishment_ is a less resolute pattern that can approximate an index.  A _measurement_ can define
a range of scalar patterns to apply using the _bits_ of each potential value - making it represent a diminishment.

    tl;dr - a measurement defines the scale to play against an index

Let's take a look at how to diminish an 11-bit index with a 3-bit pattern (commonly called a "note")

The most important aspect I'd like you to notice is the _visual_ output of "patterned bits" that
can asymmetrically tile the space.

    " Note Diminishment of an 11 bit Index "

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