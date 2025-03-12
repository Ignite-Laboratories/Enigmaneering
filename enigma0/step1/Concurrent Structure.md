# Concurrent Structure
#### Alex Petz, Ignite Laboratories, March 2025

### Parallel Kernel Execution

Now, we currently have a _single_ action running at any moment.  That simply won't do - and, more importantly,
the actions should not slow each other down!  Luckily, that is solvable using a simple `sync.WaitGroup`.  The
concept of a `WaitGroup` is simple - save off a count of the number of actions running and then let each action
decrement that count when finished.  All while the main thread just waits for all actions to complete on every
beat of the clock.

    package main
    
    import (
        "fmt"
        "sync"
    )
    
    type Clock struct {
        Period  int
        Beat    int
        Actions []func(int, int)
    }
    
    var Alive = true
    var waitGroup *sync.WaitGroup
    
    func main() {
        waitGroup = &sync.WaitGroup{}
        clock := Clock{
            Period: 1000,
            Actions: []func(int, int){
                Action,
                Action,
                Action,
                Action,
                Action,
            },
        }
    
        for Alive {
            waitGroup.Add(5)
            for i, action := range clock.Actions {
                if Potential(clock.Beat) {
                    go action(i, clock.Beat)
                }
            }
            waitGroup.Wait()
    
            clock.Beat++
            if clock.Beat >= clock.Period {
                clock.Beat = 0
            }
        }
    }
    
    func Potential(beat int) bool {
        if beat == 0 {
            return true
        }
        waitGroup.Done()
        return false
    }
        
    func Action(id int, beat int) {
        fmt.Printf("Action #%d - Beat %d\n", id, beat)
        waitGroup.Done()
    }


I promise the length of code will shrink as we start to give this project some structure - for now let's examine
what changed here.  First, the clock now can be given several different actions (though they do the same thing).
In this contrived example, we are providing the clock with a pool of five actions that fire on the **Downbeat** of
the clock.  I literally mean that in the musical sense!  These are _measures_ of rhythmic execution and it greatly
helps to frame it in that context.