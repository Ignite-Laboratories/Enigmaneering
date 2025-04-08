# `E4 - Neural Recycling`
### `Alex Petz, Ignite Laboratories, April 2025`

---

### Performance
Before we proceed _any_ further we need to address a blatant issue with this architecture.

Here's what the typical CPU load of serial execution normally looks like -

<picture>
<img alt="JanOS Logo" src="assets/4%20-%20Traditional%20Performance.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>

Here's what the it looks like the second we turn on our neural architecture -

<picture>
<img alt="JanOS Logo" src="assets/4%20-%20Neural%20Performance.png" style="display: block; margin-left: auto; margin-right: auto;">
</picture>

While that isn't exactly _high_ on the charts, it's still quite thread intensive!  That's
because our only way of activating neurons _stimulatively_ requires creating a new Go
routine for each and every impulse of the engine - which is _highly_ inefficient.  This is
a solvable issue, and one I've considered since long before I even wrote a line of impulse
code.  

### Cascading Neural Activation
To resolve it, we need to leverage the _looping_ activation in a cascading fashion.  Since
accomplishing this requires extra structure to coordinate multiple looping neurons, this
is our first `spark` structure!  At this point, we're _orchestrating_ many neurons to fire
in good timing with one another - which is the exact point of the `spark` module.

So what does it mean to fire a loop in a cascading fashion?  Well, imagine that on every impulse
you can check if a prior neuron is available to be re-activated - rather than always creating a 
new Go routine.  The neuron, itself, supports this kind of functionality already - but a _cluster_
of neurons has no awareness of each other.  An orchestrator would either _create_ a neuron on
each impulse, if none are ready to re-activate, or recycle the existing one.  This allows the
neural count to reach an equilibrium while supporting stimulative activation, all while recycling
as many Go routines as possible.