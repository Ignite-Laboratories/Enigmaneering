# `E1S4 - Mouse Antics`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### Derived Dimensions
As we continue to explore how to derive dimensions from one another we need something more interesting
to calculate off of.  Rather than continuing with such contrived examples, let's start working towards
a goal: intelligent reactions to stimulated input!  The easiest input to capture would be
the _mouse position,_ with one minor caveat - Go doesn't natively provide a way to capture the mouse
information!  Thankfully, JanOS has another module that provides such access - `host`

If you'd like to read more about how it's implemented, check out `enigma2` =)

The `host` module provides a package called `mouse` with two key members -

    SampleCoordinates()
    SampleRate

Both of these provide global access to general mouse sampling operations - and I'll break down
how to work in this paradigm as we progress through this solution.  For now, I'm going to
start demonstrating the kinds of derived dimensions that can be made from such a simple input.