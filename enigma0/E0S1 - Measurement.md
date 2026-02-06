# `E0S1 - Measurement`
### `Alex Petz, Ignite Laboratories, January 2026`

---

Now that we have an infinite ruler, we need a way to take _measurements_ across it.

But _**What**_ are we measuring?

An onion can be described in many forms - _but it's still a point cloud of atoms in space **we** consider to be an onion!_

_Everything **is** data._

Data is just facts and statistics collected together for reference and analysis - technically, an _"onion"_ only
exists as an _"onion"_ because we **collectively** accepted that's what that cluster of atoms happens to be.

_That's wonderful!_

It means that _everything_ can be measured as data, and the manifested world around You is _intricately encoded._

Our infinite ruler only has a sense of _difference_ that arises when compared to another index.  To
calculate a difference, one must collapse the potentials into _measurements._

    " Measurement "
        A value recorded from an index's currently held potential.

In binary an index is a known _bit-width,_ while a measurement is the bit-width's contained _value._

As this represents a new conceptual dimension, the best term I've found to describe the resolution of an index is its _"depth."_

To model this in code, I've built a `measurement` type that byte-packs individual bits without enforcing byte-alignment.

    tl;dr - measurements are a singular indexed value

      why - this represents a value within a variable width register