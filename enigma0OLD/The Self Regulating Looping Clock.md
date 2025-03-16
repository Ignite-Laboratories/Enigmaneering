# The Self-Regulating Looping Clock
#### Alex Petz, Ignite Laboratories, March 2025

### Jumping Gee Willies - Looping Clocks!?

Well, yes!  And there is a fundamentally important reason for this:

_**Calculation takes time**_

Even the act of _reading_ a value takes time to record the value!  By _that_ very nature a
single thread of execution can only read a single _instant of time_ at any given moment.
This brings us to our first, and most important, definitions for the future of this project:

* **An Instant** - _any_ abstract moment in time with no reference point.
* **A Moment** - a _well defined_ moment in time.
* **"Well Defined" vs "Abstract"** - A moment is 'well defined' if it has an offset to an external source of time.  This could be the number of nanoseconds since the system was turned on, or some higher-level system being observed.  An instant is abstract because the very concept of a thread of execution holding a value _requires_ an instant of time, but knowing _when_ it happened requires a reference point.

So let's circle back to my wording above - a single thread of execution can only read a single
_**Instant**_ of time at any given _**Moment**_.  The thread, itself, has no concept of what
the current _moment_ in time is because it requires calculations to get that information - typically
by reading `time.Now()` and then talking to other functions to get contextual information.

So, let's invert the problem and _provide_ the **Kernel** (a thread of execution) with the current
**Moment's** contextual information for every beat of the **Clock**.  The **Kernel** would activate
and run for a single calculation before being re-invoked - making it an externally driven _loop._

Now, we create a master _**Clock**_ that spawns off a new **Kernel** for every beat.  You'll notice
I'm specifically using the word **Beat**, rather than _tick_ - that's because this clock is special!  _It loops!_
It doesn't just provide the _actual_ **Moment** in time to the **Kernel** on invocation, it also
provides a **Beat** value indicating how high the **Clock** is currently counting up to - it's **Period**.

*Phew* - that's a lot of bolded terminology!  Let's sum them up -

* **Kernel** - An abstract thread of execution.
* **Clock Period** - A value the clock counts up to before looping the _beat_ back to 0.
* **Beat** - The current value of the clock's internal counter.
* **Clock** - A loop that drives the creation of kernels and provides them with temporal context. 

Truly, the _entire_ system hinges upon the above principle!  I know it might sound daunting, but it isn't.
The devil is in the details, as they say - so rather than trying to explain everything at once I'd like
to try breaking it into logical steps.