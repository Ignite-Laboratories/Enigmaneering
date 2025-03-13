# 2 - Logical Execution
#### Alex Petz, Ignite Laboratories, March 2025

## Activation Orchestration

Now, this current design isn't very exciting - _yet!_

Here's where we get to build some bridges between computation and neurology.  Technically, we are
mimicking the design of neurons in the brain at this point - if you consider that a neuron is just
an impulse fired under the right conditions: an **Action Potential**.  The characteristic
of the impulse and how it relates with the rest of the system is entirely based around contextual 
_timing_, which we will explore later.  For now, let's start defining some kernel characteristics
that are analogous to their neuron counterparts -

### 1 - Atomic Operation
Neurons fire in a unique fashion - a condition is met, they fire, then they enter a _refractory_
period before restarting the cycle.  Biologically, this makes complete sense as there are many 
different conditions that facilitate the neuron's existence.  Computationally, there is another 
perspective that will become clearer as we proceed forward a few steps!

    tl;dr - Kernels should not activate again while their last activation is still executing.

Implementing this requires some housekeeping to keep these enigmatic steps from getting 
out of hand - let's explore those first before circling back to atomic operation.

### 2 - Uniquely Identified
As contrived as this sounds, every single thing in this universe is uniquely identified.  This
is actually really simple, if you consider that uniqueness comes from the _context_  in which you 
observe a cluster of like things!  _**The fact they exist**_ provides enough clarity to _observe_ 
they are different entities - no matter how similar they are or if they hold the same identified
value.  

_**Because**_ of observance, there will _always_ be a uniquely identifiable perspective of any set!  
As they say: God knows exactly how many hairs are on your head - well, yes, because within a split 
second that can be _observed_ by an intelligently designed system.  We coexist in such a system =)

    tl;dr: 
        1 - Everything in the universe can be uniquely identified.
        2 - Identification is a side effect of observance.

This part is really simple to implement - 

    package core
    
    import "sync/atomic"
    
    // Alive globally keeps any long-running routine alive until it is set to false.
    var Alive = true
    
    // masterId holds the last provided entity identifier value.
    var masterId uint64
    
    // NextID provides a thread-safe unique entity identifier to every caller.
    func NextID() uint64 {
        return atomic.AddUint64(&masterId, 1)
    }

The `NextID` function provides a thread-safe way to get a unique number in the context of the 
currently executing program.  At the end of the last step I lamented that the code is getting 
really long, so this is a prime opportunity to start moving code into packages as we progress 
forward. The first package is `core` and it holds only the most critical components of `JanOS`,
such as the above code.  I'll note when we move code into a package - for now, it's just what's above.

Since this is the first step that references an external package I'd also like to note that any 
steps that do so will reference a _snapshot_ of that package which will never disappear from
its repository.  This is through the wonders of `git tags` and their built-in integration with `Go`!

### Implementation

To implement this we need to start defining some more structure around the design.  First,
we need to define what a **Kernel** and an **Action Potential** are - as well as a better
way to provide **Context**.

Context is pretty self-explanatory, but I should note that it provides pointer references 
between the clock and the kernel executing the current action or potential function.  This
allows the executing function to reference clock details or to grab its own kernel identifier. 

    type Context struct {
        Moment    time.Time
        Beat      int
        Clock     *Clock
        Kernel    Kernel
        waitGroup *sync.WaitGroup
    }

Next we'll need to make some tweaks under the clock's hood - it needs some better structure 
and an easier way to interact with it.  On top of that, the clock shouldn't be worried about 
the `if potential { action() }` mechanic as that is the _kernel's_ job.  Thus, let's also 
promote the `Kernel` to its own type that the clock executes - 

    type Clock struct {
        ID          uint64
        LoopPeriod  int
        Beat        int
        Kernels     []Kernel
    }
    
    func (c *Clock) Start() {
        var wg sync.WaitGroup
    
        for core.Alive {
            var ctx Context
            ctx.Moment = time.Now()
            ctx.Beat = c.Beat
            ctx.Clock = c
            ctx.waitGroup = &wg
    
            wg.Add(len(c.Kernels))
            for _, k := range c.Kernels {
                ctx.Kernel = k
                go k.Execute(ctx)
            }
            wg.Wait()
    
            c.Beat++
            if c.Beat >= c.LoopPeriod {
                c.Beat = 0
            }
        }
    }

    type Kernel interface {
        Execute(ctx Context)
        GetID() uint64
    }
    
    func (c *Clock) AddKernel(action func(ctx Context), potential func(ctx Context) bool) {
        ap := &actionPotential{
            id:        core.NextID(),
            action:    action,
            potential: potential,
        }
        c.Kernels = append(c.Kernels, ap)
    }

`Kernel` is a _handler_ for interfacing with executing contexts.  Currently, it only has two
functions, but that'll evolve shortly.  The most important thing to note is _this_ is the
publicly exported way to work with execution contexts, whereas the concrete implementation is
in the `actionPotential` -

    type actionPotential struct {
        id        uint64
        executing bool
        action    func(ctx Context)
        potential func(ctx Context) bool
    }
    
    func (ap *actionPotential) Execute(ctx Context) {
        if !ap.executing && ap.potential(ctx) {
            ap.executing = true
            go func() {
                ap.action(ctx)
                ap.executing = false
            }()
        }
        ctx.waitGroup.Done()
    }
    
    func (ap *actionPotential) GetID() uint64 {
        return ap.id
    }

Here's where we have implemented the concept of a neuron that cannot be invoked again until
it's current round of execution has completed.  By doing so, every single `Kernel` added
to the clock represents a solitary neuron in a network of many getting stimulated across time.

**KEY:** The _kernel_ manages neurological activation.

Finally, let's explore a contrived example that demonstrates a set of actions that only activate 
once the last round of execution ends - even though they are stimulated on every beat of the clock -

    func main() {
        clock := NewClock(1024)
        clock.AddKernel(Action, Potential)
        clock.AddKernel(Action, Potential)
        clock.AddKernel(Action, Potential)
        clock.AddKernel(Action, Potential)
        clock.AddKernel(Action, Potential)
        clock.Start()
    }
    
    func Potential(ctx Context) bool {
        // Run on every beat of the clock
        return true
    }
    
    func Action(ctx Context) {
        fmt.Printf("Action #%d - Beat #%d\n", ctx.Kernel.GetID(), ctx.Beat)
        time.Sleep(1 * time.Second)
    }

From here you'll notice in the console output immediately that the first activation has all kernels 
fire on beat 0, but subsequent activations happen on chaotically misaligned beats - welcome to what 
sparked off all of this work!  This, my friends, is the host operating system's _task scheduler_, 
combined with _entropy_.  No matter how much you try to align these things, they will _only_ align 
through clever use of the clock as a feedback loop and coordination point.

It's not too complex, but that's the next step's challenge! =)