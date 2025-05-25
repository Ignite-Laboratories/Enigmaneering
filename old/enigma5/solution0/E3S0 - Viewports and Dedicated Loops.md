# `E2S0.0 - Viewports and Dedicated Loops`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### Dedicated Loops and Dimensional Destruction
Before we proceed _any_ further, we have to address a massive issue in multi-threaded design: _many state machines
require single threaded operation._  Chief of that list is our graphics pipeline - _OpenGL._  Inherently, that
presents a problem with firing _asynchronous_ rendering impulses - each one is a new Go routine with no guarantees
of which host OS thread it executes on.  To solve this issue, I've created a new kind of `temporal.Loop` called the
`temporal.DedicatedLoop` - it behaves _exactly_ as a traditional temporal loop does, but uses a channel as it's
dimensional _Cache._  This allows an isolated Go routine, pinned to a single thread with `runtime.LockOSThread()`,
to _block_ for impulses.

Here's the creation of a `Viewport` - 

    func New(engine *core.Engine, potential core.Potential, title *string, windowSize *std.XY[int]) *Viewport {
        // Spark off a new open GL context on an impulse thread
        v := &Viewport{}
        v.Title = title
        v.Size = windowSize
        v.impulse = make(chan core.Context)
        v.Dimension = temporal.DedicatedLoop(engine, potential, false, v.tick, v.cleanup)
        return v
    }

There is only _one_ caveat of a `DedicatedLoop` - as these mechanisms typically require memory cleanup, as well,
a dedicated loop is provided with a `cleanup` method at creation.  This method will get invoked when the dimension
is _destroyed_ - an operation that clears it's neurons from the impulse engine and cleans itself up.  If you
have no method to provide, `core.DoNothing()` is provided to target for your convenience.

### Disconnected Stoppage
The impulse engine, until now, has never needed to _stop._  It can be at any time through it's `Stop()` method,
but now that we are exploring opening up a _window,_ it would make sense for the window to stop the engine when
it closes.  But the window has _no_ awareness of whether _it_ should stop the impulse engine!  The rules behind
when to stop are something that the creating system should _imbue_ upon the structure.  Thus, a _signaling_
mechanic made a helluva lot more sense, so I added a mechanism to the engine where it will stop _itself_ when
a `Potential` is met.

Let's spark off a new `Viewport` window -

    func main() {
        title := "Hello, World!"
        size := std.XY[int]{X: 420, Y: 380}
        framerate := 30.0
        var porthole = viewport.New(core.Impulse, when.Frequency(&framerate), &title, &size)
    
        core.Impulse.StopWhen(std.PotentialTarget(&porthole.Destroyed))
        core.Impulse.Spark()
    }


The first thing you will notice is the terminology of a `porthole` - at this point, these structures are
meant to create little portholes into the dimensions that are currently available.  They will evolve into
more complex structures, but _those_ structures will _not_ be portholes while still utilizing the same
mechanic of the `viewport` package.

Thus - until those structures are built, the term `porthole` allows us to identify a _kind_ of viewport,
even if it's the only kind we have at this point.

It's a real "Ship of Theseus" situation, in reverse!

### Rendering
The concept of a viewport is simple - it just creates a new window, using `GLFW` (the open Graphics Library
Framework Windowing), and then provides an overridable `Render()` method to any types that embed it.  The render
method is fired as a _looping activation_ gated by an action potential, allowing fine tuned control over
the framerate.  Please refer back to `Enigma 0 Solution 1` for a better understanding of how that mechanic
works.