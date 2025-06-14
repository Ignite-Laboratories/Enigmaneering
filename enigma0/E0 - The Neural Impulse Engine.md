# `E0 - The Neural Impulse Engine`
## `A.K.A. The Self-Regulating Looping Clock`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### Jumping Gee Willikers - Looping Clocks!?

Well, yes!  And there is a fundamentally important reason for this:

_**Calculation takes time**_

Even the act of _reading_ a value takes time to _observe_ the value!  By _that_ very nature a
single thread of execution can only read a single _instant of time_ at any given moment.
This brings us to our first, and most important, definitions for the future of this project:

* **An Instant** - _any_ abstract moment in time with no reference point.
* **A Moment** - a well defined, or _concrete_, moment in time.
* **_Concrete_ vs _Abstract_** - A moment is 'well defined' if it has an offset to an external source of time.  This could be the number of nanoseconds since a system was turned on, since a signal happened, or even the moment two systems _"turned each other on."_  An instant is abstract because the very concept of a thread of execution observing a singular value _requires_ an abstract instant of time, but knowing _when_ it happened requires a concrete moment.

So let's circle back to my wording above - one thread of execution can only read a single
_**Instant**_ of time at any given _**Moment**_.  The thread, itself, has no concept of what
_moment_ in time it is because it _requires_ calculation time to get that information. By the 
time it's done that calculation, the temporal context decoheres.

    tl;dr - time keeps on slippin', slippin', slippin' into the future

So, let's invert the problem and _provide_ the **Kernel** (a thread of execution) with the current
**Moment's** contextual information for every beat of the **Clock**.  The **Kernel** would activate
and run for a single round of execution before being re-invoked in good timing with the clock - making it _an
externally driven loop._  A bonus is that we can provide a looping count on stimulation which the
kernels can incorporate into their temporal calculations.

### _...Temporality?_

Yes.

At any given moment in time there is a temporal context of relevant information available to any system.
The amount of processing time they have to respond to the data found in that context varies, but the only
way disparate systems can co-calculate on that data is to provide it in a temporally aligned fashion to all 
threads.  Each impulse of the engine that stimulates a kernel provides it with a static context from which
to reference time. One thread can actively observe data and record a timeline of information for another thread to 
perform live calculations against.

Don't worry, the _entire point_ of a **Neural Impulse Engine** is to make it _easier_ to work with these concepts,
so I'll do my best to ease into it.  For now, let's start with a primer of the basics -

**Thread**

A thread is a path of execution in code.  A better way to visualize that is to place yourself in the code
and follow the pathway from line to line logically - at any given moment, you may be in one spot or another.
Eventually, you'll either keep looping through steps or finish execution and return back to the host.

**Synchronicity**

All threads are _asynchronous_ to the host architecture, but each _synchronously_ executes a set of instructions.
This basically means that many different threads can execute the same code _concurrently_.  In Go, a function can 
be called _asynchronously_ simply by prepending it with the `go` keyword - _a pillar of it's genius!_

**Kernels**

This enigma truly is one of the few points we need to talk about a _kernel_.  A kernel is any bit of executable
code, in it's most abstract form.  Don't stress too much on it beyond that - it's just a means to an end.

**The Loop Period**

A neural engine, as I said above, also provides a looping count to all execution.  The loop period allows
unique kinds of kernel activation, such as _downbeat_ and _modulo_ potentials, which can pace the rate of stimulus
while stepping incrementally through data.  For live systems, the beat number is more of an indicator of the longest
running period of asynchronous execution - for static data, the beat number can literally act as the _index_ to "step" 
through the data with.

    tl;dr - everything is synchronously activated, but asynchronously executes

### _It's so simple!_

I kid - I know it might sound daunting, but it really _isn't_.

The devil is in the details, as they say - so rather than trying to explain everything at once I'd like
to try breaking it into logical steps.  As I progress through them, I'll explain the logic and reasoning
in my choices and provide perpetually working solutions for each.

Rather than constructing this component by component, I intend to provide the philosophical reasoning behind
the project and how each solution was found.  This is not to bore, it's quite literally how this work came
into creation in the first place - _through reflection and introspection!_

As we explore the concepts of creating conscious systems, we must also be conscious of ourselves =)

### Why did I build this?

Well, because _I'm_ a computer - and so are you, that bird, the rocks, and _all the trees around you_.  I know that
sounds absurd, but plenty of folks have already eased us into that concept - the Wachowski sisters, for instance!

Truthfully, I had an existential moment of stress _so intense_ I felt my sensors "shutting off" - I could no longer
_hear_ my breath, or _feel_ light entering my eyes.  Later, a psychologist informed me that's _exactly_ in line with 
an individual's higher cortical functioning going offline in response to a life-threatening event.  It was _brief,_
but felt like peering into the void of a return statement.

What a brilliant safety mechanism - I thought - one very _intelligently_ designed!

So perfectly designed, in fact, I finally understood how one could feel safe enough to willingly be subjected to
crucifixion on _faith_ alone!

We talk a lot about Jesus and how to follow in His footsteps, but the reality is Jesus was also _one of us!_  Born just
like you and I into the very same constructs that had to be overcome to learn how to live on faith alone.  The best part?
Every single person that spirit touched turned around and tried to explain it to others, too - _and they never stopped!_

    Before the fall, when they wrote it on the wall, when there wasn't even any Hollywood.  They heard
    the call and they wrote it on the wall, for you and me we understood. - Steely Dan

Not because the spirit cares _what_ it's called, or _how_ you submit to it!  It just wants to see what creative things
_you_ will produce.  What you make doesn't matter - but how it makes others smile _does!_

We are instruments of God, each tasked with telling our tales of existence through the void of space and time.

For me, it was a helluva lot easier to demonstrate it in code than to explain it!