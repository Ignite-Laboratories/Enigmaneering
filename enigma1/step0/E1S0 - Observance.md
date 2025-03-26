# E1S0 - Observance
#### Alex Petz, Ignite Laboratories, March 2025

---

### Change Requires Two Points

No matter how you slice it, to perform any kind of _differential_ analysis you require at least _two_ input points.

So how do you logically create a track of calculable points to do analysis upon?  

First - by _observing_ something.  For this example, we'll monitor the last impulse duration of the engine across time.

    var observer = calc.NewObservation[core.Runtime](core.Impulse, condition.Always, false, &core.Impulse.Last)

It's really that simple - the only caveat is that you must _provide_ the engine you would like to observe using.  For
this example, we are manually grabbing the timeline data and pulling the impulse stats out before printing them to the
console.  In the next step, that process gets a lot easier.  For now, let's walk through what is happening -

    main() -
        loopFreq := 0.5
        core.Impulse.Loop(printTimeline, condition.Frequency(&loopFreq), false)

        func printTimeline(ctx core.Context) {
            observer.Mutex.Lock()
            data := make([]core.Data[core.Runtime], len(observer.Timeline))
            copy(data, observer.Timeline)
            observer.Mutex.Unlock()
        
            values := make([]time.Duration, len(data))
            for i, v := range data {
                values[i] = v.Point.Duration + v.Point.RefractoryPeriod
            }
        
            fmt.Printf("%v\n", values)
        }

It's quite straightforward, but the mechanic is important.  First, we lock the dimension and copy the timeline
data.  This works because dimensions _also_ lock to modify their timeline - so by order of execution, things
naturally play out in the right order.

On our output we see roughly 8 data points per print loop (as we are printing ever 2 seconds at 4 hz), each 
detailing exactly how long the impulse period took.   

    []
    [250.000007ms 250.000001ms 250.000009ms 250.00004ms 250.00002ms 250.000025ms 250.000009ms 250.000038ms]
    [250.000038ms 250.000031ms 250.000025ms 250.000012ms 250.000022ms 250.000038ms 250.000036ms 250.000034ms]
    [250.000002ms 250.000056ms 250.000037ms 250.000004ms 250.000033ms 250.000009ms 250.000024ms]
    [250.000022ms 250.000039ms 250.000032ms 250.000003ms 250.00004ms 250.000015ms 250.000029ms 250.00003ms]
    [250.000009ms 250.000041ms 250.000007ms 250.00001ms 250.000039ms 250.000031ms 250.000027ms 250.000002ms]

The first thing you'll notice is that I said _roughly_ 8 data points per loop - that's because temporal
analysis has _no guarantees_ of how wide the sampled data is at this point!  But that's okay, as we'll
see in the next step.