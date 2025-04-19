# `E2 - Hydra`
### `Alex Petz, Ignite Laboratories, April 2025`

---

### Multi-headed Systems

The next major module of JanOS is `hydra` - which provides a toolkit for creating parallel
_systems_ that operate off of an impulse engine.  As of this writing, the only systems in use
are entirely _graphical_ - but that's for a good reason!  Graphical pipelines come with a typical 
requirement: _host thread affinity_.

What's thread affinity?  Well, in many low-level libraries pointers are used to access data that's
been "pinned" in place - if you hop to a different host thread, those pointers no longer have access
to that global address space.  Go routines are _not_ host threads - in fact, Go's scheduler can
intermittently move your routines between host threads.  Luckily, the language provides a workaround
in the `runtime` module - when your routine starts you can call `runtime.LockOSThread()` to _wire_ 
it to a singular host thread.  

This is wonderful, but it doesn't _quite_ help a neural threading architecture.  Currently, every 
impulse is a separate Go routine _entirely._  Even if you wired the routine on every impulse, it would 
only intermittently get _launched_ on the same host thread!

To solve this issue, I designed several different tools that are _heavily_ leveraged by Hydra.  Chief
of these are the `std.Synchro` and `core.System` - by the end of this enigma we should be able to 
display multiple graphical contexts of relevant information _on demand_ using a global framerate
that's _not_ regulated by v-sync!  This comes with a _fascinating_ new drawback where frames can
_phase_ in and out of sync with the monitor, rather than necessarily _tearing_.  Until I explore the
phenomenon further, I'm not sure if it's _truly_ an issue or if my code is too rudimentary to truly
demonstrate it.

    tl;dr - Hydra allows parallel programs to be started, impulsed, and stopped.

### Stoppable Loops
What makes a system different from a standard loop?  A system can be _stopped._  Up to now we
have created dimensions that are infinitely impulsed loops.  These structures will activate so long 
as their potentials return true - which is okay in many circumstances, but what if you wish to
_control_ the execution of these impulses?

that 
require an external driver to keep them running in good timing.  As we progress, I aim to demonstrate 
that a large number of rendering contexts can be driven off of an impulse - rather than regulating the 
framerate _themselves_.

To accomplish this we have to be able to drive _threads_ of data - especially if we want to access
technologies like OpenGL which rely upon a single-threaded host pipeline.  Currently, our impulses
are entirely _multithreaded_ and don't play nicely with such technologies.  As such, we need a way to
perform _inter-thread signaling_ - which inherently requires _decoupling_ the control systems.  One
thread should be able to ask another thread to stop and _wait_ for it to acknowledge that request -

    core -
        // System represents an stoppable looping structure.
        type System struct {
            Alive    bool
            Stopping bool
            Cleanup  func()
        }
        
        // Stop will signal the system to cease cycling, then block until it processes the request.
        func (sys *System) Stop() {
            sys.Stopping = true
            for core.Alive && sys.Alive {
                // Hold until finished
                time.Sleep(time.Millisecond) // Don't waste cycles =)
            }
            sys.Cleanup()
        }

For now, let's focus on the `Stop()` method.  This _very simple_ mechanic ensures that the calling
thread can wait for the system to finish whatever it needs to do _before_ it advances.  In the case
of graphics, that means giving the current frame a moment to finish drawing operations before calling
its `Cleanup()` method and subsequently destroying the window it's drawing to.
