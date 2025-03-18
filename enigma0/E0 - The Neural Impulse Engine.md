# The Neural Impulse Engine
### A.K.A. The Self-Regulating Looping Clock
#### Alex Petz, Ignite Laboratories, March 2025

---

### Jumping Gee Willikers - Looping Clocks!?

Well, yes!  And there is a fundamentally important reason for this:

_**Calculation takes time**_

Even the act of _reading_ a value takes time to _observe_ the value!  By _that_ very nature a
single thread of execution can only read a single _instant of time_ at any given moment.
This brings us to our first, and most important, definitions for the future of this project:

* **An Instant** - _any_ abstract moment in time with no reference point.
* **A Moment** - a well defined, or _concrete_, moment in time.
* **_Concrete_ vs _Abstract_** - A moment is 'well defined' if it has an offset to an external source of time.  This could be the number of nanoseconds since a system was turned on, since a signal happened, or even the moment two systems _"turned each other on"_.  An instant is abstract because the very concept of a thread of execution observing a singular value _requires_ an instant of time, but knowing _when_ it happened requires a referential point.

So let's circle back to my wording above - one thread of execution can only read a single
_**Instant**_ of time at any given _**Moment**_.  The thread, itself, has no concept of what
_moment_ in time it is because it _requires_ calculation time to get that information. By the 
time it's done that calculation, the temporal context decoheres.

    tl;dr - time keeps on slippin', slippin', slippin' into the future

So, let's invert the problem and _provide_ the **Kernel** (a thread of execution) with the current
**Moment's** contextual information for every beat of the **Clock**.  The **Kernel** would activate
and run for a single execution before being re-invoked in good timing with the clock - making it _an
externally driven loop._  A bonus is that we can provide a looping count with stimulation which the
kernels can incorporate into their calculation.

### ...Temporality?

Yes.

At any given moment in time there is a temporal context of relevant information available to any system.
The amount of processing time they have to respond to the data found in that context varies, but the only
way disparate systems can co-calculate on that data is to provide it in a temporally aligned fashion to all 
threads.  Each impulse of the engine that stimulates a kernel provides it with a static context from which
to reference time. One thread can actively observe data and record a timeline of information for another thread to 
perform live calculations against.

Don't worry, the _entire point_ of a **Neural Impulse Engine** is to make it _easier_ to work with these concepts,
so I'll do my best to ease into these topics.  For now, let's start with a primer of the basics -

**Thread**

A thread is a path of execution in code.  A better way to visualize that is to place yourself in the code
and follow the pathway from line to line logically - at any given moment, you may be in one spot or another.
Eventually, you'll either keep looping through steps or finish execution and return back to the host.

**Synchronicity**

All threads are _asynchronous_ to the host architecture, but each _synchronously_ executes a set of instructions.
This basically means that many different threads can execute the same code _concurrently_.  In Go, a function can 
be called _asynchronously_ simply by prepending it with the `go` keyword - a pillar of it's genius!

**Kernels**

This enigma truly is one of the few points we need to talk about a _kernel_.  A kernel is any bit of executable
code, in it's most abstract form.  Don't stress too much on it beyond that - it's just a means to an end.

**The Loop Period**

A neural engine, as I said above, also provides a looping count to all execution.  The loop period allows
unique kinds of kernel activation, such as _downbeat_ and _modulo_ activation, which can pace the rate of stimulus.

### It's so simple!

I kid - I know it might sound daunting, but it really _isn't_.

The devil is in the details, as they say - so rather than trying to explain everything at once I'd like
to try breaking it into logical steps.  As I progress through those steps, I'll explain the logic and reasoning
in my choices and provide perpetually working examples for each.

Rather than constructing this component by component, I intend to provide the philosophical reasoning behind
the project and why each step was executed.  This is not to bore, it's quite literally how this work came
into creation in the first place - _through reflection and introspection!_

As we explore the concepts of creating conscious systems, we must also be conscious of ourselves =)