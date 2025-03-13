# 0 - The Impulse Engine
#### Alex Petz, Ignite Laboratories, March 2025

## Make it so

Alright, the very first thing we get to build is the engine that drives kernel impulses.  As we already
established, this is a _looping clock_.  So - let's build exactly that!

First things first - we need our clock to have periodicity with which to count up to.  This is quite straightforward,
but importantly leverages a _Go routine_ to create concurrency by using the `go` keyword when the action is invoked.
This allows each action to run in an isolated scope _on every beat of the clock_, even if the last action hasn't 
yet finished running.

    type Clock struct {
        LoopPeriod int
        Beat       int
        Action     func(int)
    }
    
    var Alive = true
    func main() {
        clock := Clock{LoopPeriod: 1024, Action: Action}
        
        for Alive {
            go clock.Action(clock.Beat)

            clock.Beat++
            if clock.Beat >= clock.LoopPeriod {
                clock.Beat = 0
            }
        }
    }
    
    func Action(beat int) {
        fmt.Printf("Beat %d\n", beat)
    }

#### Key Principle -
The first principle I'm establishing is the concept of an `Alive` boolean controlling the program's continued
execution.  Many systems use a `Terminate` paradigm, but `Alive` provides a _positive_ condition for continuing 
execution - **_which is idiomatic to read!_**  Code should _always_ be written for readability, wherever 
possible.  From left to right `"if I'm still alive then do something"` is intuitive because if it's executing
it's still alive!  Whereas `"if I shouldn't be terminating then do something"` causes a mental pause to consider
if other things out of our scope should be handled - a _micro-anxiety_ we don't need to put on our readers.

I promise this is important!  Every ounce of how _any_ system _(yourself included!)_ observes and consumes complex
information is relevant in a conscious toolkit.

### Cool beans, it can count - whoop-dee-doo!

The above is only _half_ of the equation - the other half of the _action_ is its _potential_ to be activated.
That's what makes it an **Action Potential** execution pipeline, after all!  Luckily, building in potential
is quite simple, too.  The potential function can be as complex as desired, and the clock can provide as much
context to the action potential functions as desired:

    type Clock struct {
        LoopPeriod int
        Beat       int
        Action     func(int)
    }
    
    var Alive = true
    func main() {
        clock := Clock{LoopPeriod: 1000, Action: Action}
    
        for Alive {
            if Potential(clock.Beat) {
                go clock.Action(clock.Beat)
            }
    
            clock.Beat++
            if clock.Beat >= clock.LoopPeriod {
                clock.Beat = 0
            }
        }
    }
    
    func Potential(beat int) bool {
        if beat == 0 {
            return true
        }
        return false
    }
    
    func Action(beat int) {
        fmt.Printf("Beat #%d\n", beat)
    }

Truly - _that is all it takes to begin building an action-potential execution pipeline!_

It does **not** require dropping to low level languages like _C_ to begin building executive systems =)