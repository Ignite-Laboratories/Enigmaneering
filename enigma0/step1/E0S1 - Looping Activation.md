# E0S0 - Looping Activation
#### Alex Petz, Ignite Laboratories, March 2025

---

### Activation?

Yes!  Remember before how I stated that this enigma is the only one we _really_ need to think about the **kernel** in?
Well, that's because the _entire concept_ of a neural impulse engine hinges on the concept of _activating_ execution.

Colloquially, the terms `activate` and `execute` are pretty interchangeable - but an _activation_ happens _cyclically_!
Neurons are microscopic little programmatic loops, at their core, constantly being stimulated with impulses from which
they perform the same calculations over and over again.  While the technicalities of activation will come later, it's
crucial to understand how activation happens across time.

<picture>
<img alt="JanOS Logo" src="../assets/E0S1D0 - Single Beat.svg" height="200" >
</picture>

The first concept of activation is the _impulse_.  For every loop of the engine, an impulse of execution is coordinated.
Immediately, the current moment of _time_ is recorded and provided to every kernel of execution through a `core.Context`
structure.  All non-executing kernels are invoked, the engine's current beat is incremented, and the process repeats.

<picture>
<img alt="JanOS Logo" src="../assets/E0S1D1 - Multiple Beats.svg" height="350" >
</picture>

Each cycle of the loop is an **impulse**.  Right now we are looking at the timeline of a loop currently at beat 42.
Abstractly, there is no concept of _time_ through this lens - but the provided `core.Context` given by the engine
allows every activated kernel to know the **impulse moment** it's activation stemmed from.  When combined with
a beat value, kernels can intelligently decide how to respond.  Is it an event beat?  An odd beat?  Could I skip
every 15 beats by modulo-ing the provided beat number with 16?

These are the kinds of decisions a neural impulse engine provides it's kernel activations! Let's look at a single
activation across time - this one is still executing and has yet to calculate a value:

<picture>
<img alt="JanOS Logo" src="../assets/E0S1D2 - Activation.svg" height="350" >
</picture>

The first thing you can see is that an activation does _not_ begin at the impulse moment!  That's because coordinating
the launch of many activations requires _time_.  Above, that is chocked up purely to _entropy_ - but, thanks to the
`core.Context` every activation is temporally aware of it's relative offset to the impulse moment at all times.

The next thing you'll see is that an activation, when plotted across time as such, represents a _track_ of execution
on a _timeline_.  I'll touch briefly on why this is relevant in a second.

The final thing you should notice is that an activation _can_ span across multiple impulses - if it is _asynchronously_
invoked.  If it's not, it's considered a _blocking activation_ that quite literally slows the rate of impulse.

<picture>
<img alt="JanOS Logo" src="../assets/E0S1D5 - Blocking Activation.svg" height="300" >
</picture>

Here, the activation finished and simply returned back the beat number it was passed.  Since it was instantiated
on impulse beat 42, the value is associated with that grid position on the X-axis representing time.

As noted, this activation style blocks the next impulse from happening until it completes.  This quality can be leveraged 
to add temporal _weight_ to the system, slowing it's overall execution to ease load on the host architecture.  We don't 
need to worry about that now, but it's important to keep in mind!

If an activation is triggered _asynchronously_ it can either be executed in one of two ways.  The first is a *stimulated*
activation for every single impulse.  In this setup, the activations fire as fast as possible without care for whether
the last activation has finished -

<picture>
<img alt="JanOS Logo" src="../assets/E0S1D4 - Asynchronous Activation.svg" height="300" >
</picture>

The second kind of asynchronous activation is a _looping activation_.  For this configuration, a single kernel is cyclically
executed on the following impulse after completion.  The period between completion and re-invocation is considered the 
_refractory period_ between neural activations.

<picture>
<img alt="JanOS Logo" src="../assets/E0S1D3 - Cyclic Activation.svg" height="350" >
</picture>

### Temporal Fragment Shading

Before we proceed any further I want to briefly tease you with what this kind of activation scheme can yield.  Just as pixel
fragment shaders are executed in _parallel_ with their positional context, our activations are provided with their _temporal
positional context_.  While these activations occur entirely _concurrently_, rather than necessarily in _parallel_, they
can still refer to a specific index on a temporal frame buffer which is coalesced by a looping activation.  I don't intend
to go too far into the weeds about that, yet - but visually it's easy to understand at this point in the process -

<picture>
<img alt="JanOS Logo" src="../assets/E0S1D6 - Logical Activation.svg" height="500" >
</picture>

Each activation knows which index of a timeline it should be "shading" in a value for.  A looping activation can observe
the data recorded by these temporal fragment shaders at it's own pace and the impulse window itself frames the range
it should calculate.  It knows which beat it activated on, and the beat the last activation ended on - thus, it can
always reliably retrieve the exact indexes that each round of fragment shading handled while it was processing the last
round of data.  This creates a variable window of observance dictated entirely by the performance of the host architecture.

    tl;dr - overall system load directly affects observational fidelity.