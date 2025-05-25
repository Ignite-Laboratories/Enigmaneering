# `E1 - Temporal Analysis`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### Reverse Calculus

The purpose of this engine is to create _feedback loops_ - seriously, that is the _heart and soul_ of intelligence!
How?  Through logical real-time dimensional analysis - also known as _calculus._  Don't worry, we are technically 
working in the _reverse perspective_ of traditional calculus, and it's a _lot_ easier to get your feet wet in this 
direction.  We won't be breaking out any fancy mathematical equations - _yet_ - instead, we are building a calculus _toolkit._

So - what makes this a _reverse_ perspective?  

Because an integral calculates to _infinity_, but an impulse engine provides us with a _finite limit_ from which we
_sample_ time.  While we _can_ calculate to the limit, the system itself provides a mechanic of _real-time point calculation_
which can then be analyzed as a set that yields a single output - meaning we are working backwards from the integral by
building in the calculations to _live_ generate it.  The level of resolution and fidelity of the data to be analyzed is
a new dynamic of this system that must be taken into account - _I'm just the messenger!_

I promise, my work aims to make all of this _easier_ =)

### Another Dimension, Another Dimension, Another Dimension

The most primitive component of dimensional analysis is the `Dimension` - a structure that contains both the _Current_ 
value and a _Timeline_ of historical values that this dimension has held.  A dimension contains two neurons - a _Trimmer_
and a _Stimulator_.  The Trimmer is activated as a loop that constantly trims timeline entries older than the observational
_Window_ of the dimension - by default, 2 seconds.  The Stimulator is typically activated _impulsively_ to populate entries on 
the dimensional timeline, though more advanced analyzers can leverage _looping_ stimulators. The last notable thing a dimension 
contains is a synchronization object - _Mutex_ - which should be locked anytime the dimension's timeline is being modified.

### core.temporal

The majority of the work in these solutions leverages the next package of core - `temporal` - this is a container of helper
methods for _creating_ and _working with_ dimensions.  As we walk through the solutions, we'll explore the temporal package.

### core.std

The final package in the `core` module is `std` - this contains some commonly used types across the system.  For the majority
of this enigma, we will be leveraging the standard _coordinate_ structures - `XY[T] - XYZ[T] - XYZW[T]` - these require a type
constraint of `core.Numeric`, which constrains them to `integer` _and_ `floating point` types.  We'll also be leveraging
a type called `HardRef` that allows for _inline_ references to fixed values, rather than referencing another variable.

Obviously, if your system requires a _variable_ then a _hard_ reference simply will not do - but for _prototyping_, it's 
very handy!

Here's an example -

    Without HardRef -
        var freq = 16.0
        core.Impulse.Loop(Stimulate, when.Frequency(&freq), true)

    With HardRef -
        core.Impulse.Loop(Stimulate, when.Frequency(std.HardRef(16.0).Ref), true)

Finally, we have a standard method signature for retrieving a reference to a target variable, called a `TargetFunc` -

    core.std -
        type TargetFunc[TValue any] func() *TValue
        
        func Target[TValue any](val *TValue) TargetFunc[TValue] {
            return func() *TValue {
                return val
            }
        }

The method `std.Target(val)` provides a quick shorthand for creating a `TargetFunc`

We're in "bat country" - everything is a pointer and how we retrieve and pass them around is _important!_  The tools
in the `std` package have been very carefully curated to help you navigate this journey _effectively_.  I won't detail
every type in this package here, but as we are introduced to them I'll explain their design and purpose.