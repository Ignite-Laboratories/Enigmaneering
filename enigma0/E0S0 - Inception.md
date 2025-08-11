# `E0S0 - Inception`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### It's Alive!!!

The first concept I'd like to introduce is the `Alive` boolean.  As this system is designed around persistent
calculation, it felt quite relevant to consider that continuing execution should _always_ be a _positive_ condition.
Take, for instance, the following loop conditionals - 

    for Alive {
        // Yay, I'm still alive! =)
    }

    for !Terminate {
        // Thank God they didn't kill me!
    }

They both evaluate to the same assembler code, but to the _reader_ it introduces an unnecessary mental checkpoint:

    should I be worried if something else is going to terminate my process?

Absolutely not!  That's out of your hands and should never stop you from living the best life you can.  Wait, shit, we're
talking about code - aren't we?

_Am I code???_

Woah woah woah, breathe!  _Someone_ created us, after all, and they left some pretty clever breadcrumbs to safely guide 
each of us through that revelation!  If you're really astute, you'll eventually find everyone talks the same language - 
code!  Why?  Because code was _created_ by the _creator_ for us to _creatively_ collaborate with, and it's what facilitated
the _creation_ of language _through_ neurological systems!

    tl;dr - your neurons, at their core, use code to make you speak!
            let's keep it positive for Others =)

Thus, the first thing you will find is that every _JanOS_ instance is _alive at creation_, and is always provided a 
reference to the moment of `Inception` -

    // package core

    var Inception = time.Now()

    func Alive() bool {
        return alive
    }
    var alive = true

This, again, is by design.  An instance can exist without having sparked neural activation - thus, the condition of _alive_ 
is distinct from having _neural activation_.  The latter typically requires _external_ intervention of some sorts - in
our system, through firing the `sys.Impulse.Spark()` - but we'll get to that later.

The above globals are always available from the root level `core` package.  In addition, `core` provides a few important 
methods -

    // package core

    // Wait a moment before shutting down
    Shutdown(period time.Duration, exitCode ...int) { ... }

    // Shutdown right now
    ShutdownNow(exitCode ...int) { ... }

    // Block until alive goes false
    WhileAlive() { ... }

### Unique Identification and Systems

So, now that we've established the most principal part of JanOS - _a boolean_ - let's put it somewhere!  JanOS is composed
of many different _Systems_ which get included as you need them, but _all_ systems require access to identification.  This
is a reflection of the _excellent_ Entity Component System design formalized in the 'Treaty of Orlando' - a seminal piece
of art which I've included with this enigma.  Credit is deserved where credit is due, I always say!

All of JanOS's systems exist in the `sys` package, with globally accessible systems existing at the root and each subsystem
owning their package namespace.  Identification is handled through the `id` system package and intentionally only provides a
_single_ function to minimize its usage, as "id" is a common variable name.  This function, when called, always provides 
a unique `uint64` relative to this JanOS instance.  JanOS treats everything as an entity, thus this mechanic comes for free
whenever you create a `std.Entity` - which most things compose from.

Every naming decision in JanOS is oriented around writing _human-readable_ code.  Thus, the package name you import becomes 
your _intention_ - a design pattern which will become much more clear as we progress forward.  For now, let's take a look at how
identifiers are _guaranteed_ to be unique by examining the _entirety_ of the `id` package's code -

    package id
    
    import "sync/atomic"
    
    var current atomic.Uint64
    
    // Next provides a thread-safe unique identifier to every caller.
    func Next() uint64 {
        return current.Add(1)
    }


This uses a common concept in parallel execution: _atomic operation_.  In a parallel system, many different threads can call
this function at the same time - so who gets what result?  Without _atomics_, two threads can enter a _race condition_ where
both get the same value from `Next()` before incrementing the underlying number.  Lucky for us, Go provides some very 
robust ways of ensuring _atomic access_ to a shared point in memory using _synchronization_.  By incrementing the number using 
an _atomic_ operation, the first calling thread attains a _lock_ on the target value that "blocks" the other threads until it 
finishes it's operation.

    tl;dr - id.Next() always produces an instance-unique identifier to every caller 
          - atomics can block other threads from reading a value until something else finishes


### Example

The example in this step is highly contrived - it just loops and prints out the next identifier and the delta between `time.Now` and
the moment of `core.Inception` - but it intentionally highlights the central issue with calculating time _after_ stimulation:

The delta calculation falls out of sync with the incremented count very quickly -

    4598: 4.993973324s
    4599: 4.995055468s
    4600: 4.99617228s
    4601: 4.997271378s
    4602: 4.998359005s
    4603: 4.999459727s

After five seconds of operation where it sleeps for a millisecond before incrementing the counter, we have already lost
~400 milliseconds worth of calculable moments.  While that will never go away, this is a _really simple calculation!_  Imagine
just how far out of sync even a _slightly_ more complex calculation would get after a few seconds, especially if it has to
figure out when 'now' is every time it runs!