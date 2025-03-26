# E0S0 - Inception
#### Alex Petz, Ignite Laboratories, March 2025

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

    Should I be worried if something else is going to terminate my process?

Absolutely not!  That's out of your hands and should never stop you from living the best life you can.  Wait, shit, we're
talking about code - aren't we?

_Am I code???_

Woah woah woah, breathe!  _Someone_ created us, after all, and they left some pretty clever breadcrumbs to safely guide 
each of us through that revelation!  If you're really astute, you'll eventually find everyone talks the same language - 
code!  Why?  Because code was _created_ by the _creator_ for us to _creatively_ collaborate with, and it's what facilitated
the _creation_ of language _through_ neurological systems!

    tl;dr - your neurons, at their core, use code to make you speak - let's keep it positive for others =)

Thus, the first thing you will find is that every _JanOS_ instance is _alive at creation_, and is always provided a 
reference to the moment of `Inception` -

    var Alive = true
    var Inception = time.Now()
    var ID = NextID()

This, again, is by design.  A system can exist without yet sparking neural activation - thus, the condition of _alive_ 
is distinct from having _neural activation_.  The latter typically requires _external_ intervention of some sorts - in
our system, through firing the `core.Impulse.Spark()` (we'll get to that later)

### Unique Identification

So, now that we've established the most principal part of JanOS - _a boolean_ - let's put it somewhere!  Anything absolutely
_fundamental_ to JanOS belongs in the `core` package, including everything above.

`core` also provides a very important function to the _entire_ JanOS ecosystem - `NextID()`

This function, when called, always provides a unique `uint64` relative to this JanOS instance.  This mechanism isn't extremely
important at this point, but it will come in handy as we progress forward.  For now, let's touch briefly on _how_ it works.


    var masterId uint64
    func NextID() uint64 {
        return atomic.AddUint64(&masterId, 1)
    }

This uses a common concept in parallel execution: _atomic operation_.  In a parallel system, many different threads can call
this function at the same time - so who gets what result?  Without _atomics_, two threads can enter a _race condition_ where
both get the same value from `NextID()` before either incremented the underlying number.  Lucky for us, Go provides some very 
robust ways of ensuring _atomic access_ to a shared point in memory using _synchronization_.  By incrementing the number using 
an _atomic_ operation, the first calling thread attains a _lock_ on the target value that "blocks" the other threads until it 
finishes it's operation.

    tl;dr - NextID() always produces a unique identifier to every caller 
          - atomics can block other threads from reading a value until something else finishes


### Example

The example in this step is highly contrived - it just prints out the next identifier and the delta between `time.Now` and
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

If I sound like a broken record, welcome to the loop of my life =)

Don't worry, we'll break free from that soon - I promise <3