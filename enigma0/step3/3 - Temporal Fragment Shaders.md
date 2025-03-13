# 3 - Temporal Fragment Shaders
#### Alex Petz, Ignite Laboratories, March 2025

## Fuzzy Timelines and Temporal Frames 
I hope I thoroughly have established that _reading_ a value _takes time!_  If not, perhaps my colloquialisms
just strike the right chord with you - who knows - but it's crucial to understand!  Why?  Because that's _why_
we created the impulse engine - _to drive the observation of time._

So, how could a chaotic batch of execution coalesce an observed signal?  Through a **Temporal Frame Buffer**.
It's gonna take a few steps to get there, but in this one we will lay the foundation down.

### Let's dive in =)

First, however, let's soup up the impulse engine a little bit - 

    type Clock struct {
        ID          uint64
        Beat        int
        LoopPeriod  int
        BPL         int
        Kernels     []Kernel
    }

The first thing we are going to define is a new term: **BPL**.  This is the number of **Beats Per Loop** of the
clock.  Currently, the clock has a fixed `LoopPeriod` - but that'll be changing as well.  The next thing we are
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
            // Now should always be saved as the first step of the loop
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
of BeatsPerLoop every 1024 beats.  There is no rhyme or reason to this number other than it's not too small
or too big and seems to provide a reasonable figure - at this point, there isn't really a way to know if it
is or isn't accurate!  But it gives a good ballpark figure, from my experience.

Next, we will add some functionality to the `actionPotential` system -

    type actionPotential struct {
        ID             uint64
	
        // lastTrigger is the last beat's moment this actionPotential was activated.
        lastTrigger    time.Time
        
        // lastCompletion is the last moment in time this actionPotential finished execution.
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
            ap.lastTrigger = ctx.Moment
            
            go func() {
                ap.action(ctx)

                // ...and record when this kernel finished executing
                ap.lastCompletion = time.Now()

                ap.executing = false
            }()
        }
        ctx.waitGroup.Done()
    }

We've added two values to the structure - `lastTrigger` and `lastCompletion`.  Both of these are being set
by the kernel during and immediately after execution.

### Phew - that's a lotta code!
Don't worry, after this step much of the code will be migrated into the `impulse` package - for now, let's
continue!  Now that we have some data points to record we need to create a structure to hold the data for us -
