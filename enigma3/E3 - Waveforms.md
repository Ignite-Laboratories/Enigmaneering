# `E3 - Waveforms`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### Bounded Dimensions

A `Waveform` is a unique kind of `Dimension` - one that has been constrained to only `core.Numeric` types.
That restricts these kinds of dimensions to _mathematical_ operations that use integer and floating-point
numbers.  This will come in heavy use as we progress forward, but for now I get to introduce the next
two critical modules to the `JanOS` ecosystem - `spark` and `glitter`

_Spark_ provides a toolkit for orchestrating dimensional systems, to put it eloquently.  The term intends
to express that you are _sparking a dimension into existence_, as that's literally what you're doing!

_Glitter_ provides a toolkit for rendering graphical _viewports_ of dimensional systems.  This includes
the basic structures to build your own renderable outputs, which is what this enigma will explore.

The goal of this particular enigma is to demonstrate creating many viewports coordinated by a singular
impulse engine, with each calculating its waveform data off of the others in real time.  The later,
as we've already tested, is absolutely feasible - the _rendering_ component, however, is what we're
exploring here.

### Time-Oriented vs Time-Consuming
There are two different perspectives of execution I am using this project to highlight, and this is an
excellent moment to talk about the differences.  _All_ execution is time-consuming, as the name implies,
but only _some_ execution is time-_oriented._  Simulations are a prime example of a time-_oriented_
system, as things like physics require the passing of a logical amount of time to perform their next 
calculations.  Rendering systems are a prime example of a time-_consuming_ system as they aim to _maximize_
their frames-per-second, rather than calculating temporal values.



For now, let's start creating some `viewports` =)