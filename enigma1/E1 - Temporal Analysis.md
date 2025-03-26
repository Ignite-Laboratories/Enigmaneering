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
and a _Stimulator_.  The Trimmer is activated as a loop that actively trims timeline entries older than the observational
_Window_ of the dimension.  The Stimulator is activated impulsively to populate entries on the dimensional timeline. The 
last notable thing a dimension contains is a synchronization object - _Mutex_ - which we will see heavily used throughout 
these steps.

### core.Temporal

The majority of the work in these steps will leverage the final package of core - `temporal` - this is a container of helper
methods for _creating_ and _working with_ dimensions.