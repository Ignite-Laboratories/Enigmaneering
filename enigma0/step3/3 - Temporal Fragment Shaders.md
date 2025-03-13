# 3 - Temporal Fragment Shaders
#### Alex Petz, Ignite Laboratories, March 2025

## Fuzzy Timeline Population 
I hope I thoroughly have established that _reading_ a value _takes time!_  If not, perhaps my colloquialisms
just strike the right chord with you - who knows - but it's crucial to understand!  Why?  Because that's _why_
we created the impulse engine - _to drive the observation of time._

So, how could a chaotic batch of executed kernels coalesce an observed signal?  Through **Temporal Fragment Shading**.
While we're operating in `Temporal Space`, the concept of _shader programs_ still applies due to a special aspect of their
relativity - each moment in time is _relative_ to the downbeat moment of the loop, just as pixels are relative to their 
physical location on a screen.  As Einstein taught us - _everything is relative_

Effectively, that makes the _loop itself_ the definition of the shortest observable window of time.  It defines the relative
start and end points of _calculable_ observation.  Each beat of the clock represents an index of time to be observed, which
is contextually passed to a freshly activated temporal fragment shader (or, in other words, a _kernel_ or a _neuron_)

To do this, we will construct a way to hold fragmented information and a timeline from which many threads can read and write.

### Let's dive in!

First, however, let's soup up the impulse engine a little bit - 

    type Clock struct {
        ID          uint64
        Beat        int
        LoopPeriod  int
        BPL         int
        Kernels     []Kernel
    }

The first thing we are going to define is a new term: **BPL**.  This is the number of **Beats Per Loop** of the
clock.  Currently, the clock has a fixed `LoopPeriod` - but that'll soon change as well.  The next thing we are
going to do is add some extra data points to the `Context` -

    type Context struct {
	    // Delta is the amount of time that has passed since the Kernel's last impulse.
        Delta time.Duration

	    // LastDuration is the amount of time that the Kernel last took to finish execution.
        LastDuration time.Duration

        Moment time.Time
        Beat int
        Clock *Clock
        Kernel Kernel
        waitGroup *sync.WaitGroup
    }

Here we have provided two new values that can be used to perform some intelligent analysis.  The next part
involves making some tweaks to the clock.  

    func (c *Clock) Start() {
        var wg sync.WaitGroup
        beat := 0
        lastNow := time.Now()
    
        // Record the number of beats per loop
        beatCount := 0
        beatCountStart := lastNow
    
        for core.Alive {
            // Now should always be saved as the first step of the loop!
            now := time.Now()
            beatCount++
    
            var ctx Context
            ctx.Moment = now
            ctx.Delta = time.Duration(0)
            ctx.Beat = beat
            ctx.Clock = c
            ctx.waitGroup = &wg
    
            // Every 1024 beats...
            if beatCount > 1024 {
                // ...calculate the clock rate
                elapsed := now.Sub(beatCountStart).Seconds()
                c.BPL = int(float64(beatCount) / elapsed)
                
                // ...and reset the counter
                beatCount = 0
                beatCountStart = now
            }
    
            for _, k := range c.kernels {
                wg.Add(1)
                ctx.Kernel = k
                go k.Execute(ctx)
            }
            wg.Wait()
    
            beat++
            if beat >= c.LoopPeriod {
                beat = 0
            }
            lastNow = now
        }
    }

This is pretty straightforward - we have added an internal counter to the clock that can calculate the number
of `BeatsPerLoop` every 1024 beats.  There is no rhyme or reason to this number other than it's not too small
or too big and seems to provide a reasonable figure - at this point, there isn't really a way to know if it
is or isn't accurate!  1024 gives a good ballpark figure, from my experience.

Next, we will add some functionality to the `actionPotential` system -

    type actionPotential struct {
        ID             uint64
	
        // lastBeatMoment is the last beat's moment from which this type of Kernel was activated.
        lastBeatMoment    time.Time
        
        // lastCompletion is the last moment in time this type of Kernel finished execution.
        lastCompletion time.Time
        
        executing      bool
        action         func(ctx Context)
        potential      func(ctx Context) bool
    }

    func (ap *actionPotential) Execute(ctx Context) {
        if !ap.executing && ap.potential(ctx) {
            ap.executing = true
    
            // If we are not on the first beat...
            if !ap.lastTrigger.IsZero() {
                // ...Calculate how long has passed since the last beat
                ctx.Delta = ctx.Moment.Sub(ap.lastTrigger)
                ctx.LastDuration = ap.lastCompletion.Sub(ap.lastTrigger)
            }
    
            go func() {
                ap.action(ctx)
    
                // ...and record when this kernel finished executing
                ap.lastCompletion = time.Now()
    
                ap.executing = false
            }()
    
            // Temporally regress this beat's contextual information for historical reference
            ap.lastBeatMoment = ctx.Moment
        }
        ctx.waitGroup.Done()
    }

We've added two values to the structure - `lastBeatMoment` and `lastCompletion`.  Both of these are set
by the kernel during and immediately after execution.  I want to note my choice of terminology - we are
_regressing_ the momentary information temporally.  The act of observing these regressions is what provides
integration of the data through calculus, which mathematically operates to an infinite resolution.  But we
are defining *executed math* - in this system, we can consider we've hit the literal boundary of 
`Calculatable Time` and are _always_ integrating from the most resolute information possible back up. While
more calculations could be performed on the data, the *recorded* data is already as detailed as it could ever 
possibly be!

This makes the data points naturally align in a multidimensional matrix because each index's offset always 
represents the same index in time.  In a computer, this allows the data to rapidly be processed in parallel,
which is effectively _executing_ the concepts of calculus through integration.

    tl;dr - Recorded data is temporally aligned for easy parallel calculation

### Phew - that's a lotta code!
Don't worry, after this step much of the code will be migrated into its own package - for now, let's
continue!  Now that we have some data points to record we need to create some structures to hold the data -

    var timeline sync.Map
    
    type TemporalFragment struct {
        lastNow      time.Time
        lastDuration time.Duration
        delta        time.Duration
    }
    
    func NewTemporalFragment(now time.Time, duration time.Duration) TemporalFragment {
        return TemporalFragment{
            lastNow:      now,
            lastDuration: duration,
            delta:        duration - pulseDuration,
        }
    }

The `Timeline` variable is a `sync.Map` as it provides thread safe access to read and write against, but 
it's effectively just a slice of `TemporalFragment` slices - effectively making each dictionary entry
a _track_ of time in a looping timeline.  Visually, it helps to consider a digital audio workstation and
each kernel as a unique instrument that records its specific track of data.  It can only record one instant
at a time, but when many instruments are recording they can fill in a complete set of information.