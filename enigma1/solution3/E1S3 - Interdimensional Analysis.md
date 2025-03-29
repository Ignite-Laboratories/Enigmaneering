# `E1S3 - Interdimensional Analysis`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### Real Time Integration
The next feature of the `temporal` package is `Integration` - let's change the neural loop that drives our
print function to one created by `temporal` -

    Before -
	    core.Impulse.Loop(printTimeline, when.Frequency(std.HardRef(1.0).Ref), false)

    After -
	    temporal.Integration(core.Impulse, when.Frequency(std.HardRef(1.0).Ref), false, false, printTimeline, incrementer)

This dimension _naturally_ performs the duplicate filtration that we just implemented - providing an _integral
dimension,_ rather than a neural activation.

A temporal integration requires the print function to change its signature to an `Integral` -

    type Integral[TIn any, TOut any, TCache any] func(core.Context, *TCache, []TIn) TOut

Integrals are provided a slice of temporal data points to calculate a single point from.  In addition to this,
they are provided a _cache_ reference that can be used to pass information between integral activations.  That's
not as relevant at this point, but it will come in handy much later.

    func printTimeline(ctx core.Context, cache *any, data []std.Data[int]) int {
        total := 0
        for _, v := range data {
            total += v.Point
        }
    
        fmt.Printf("%v - %v\n", data, total)
        return total
    }

Our printTimeline _integral_ now returns back an `int`, making the timeline of the integration dimension of type
`*Dimension[int, any]`.  The integral dimension provides the latest found "total" of the source dimension,
since our print function adds the latest found data together.  This mechanism provides a way of _calculating_
the "area under a curve" as fast as possible, if not throttled with a slower frequency potential.

The next thing you'll notice is that our print function no longer has to reference the incrementer directly,
allowing it to neatly be provided as a _local_ variable in context at creation -

    func main() {
        var incrementer = temporal.Differential(core.Impulse, when.Always, false, increment)
        temporal.Integration(core.Impulse, when.Frequency(std.HardRef(1.0).Ref), false, false, printTimeline, incrementer)
        core.Impulse.MaxFrequency = 4
        core.Impulse.Spark()
    }

### Output
Let's take a look at the output for a moment -

    [] - 0
    [1 2 3 4] - 10
    [5 6 7 8] - 26
    [9 10 11 12 13] - 55
    [14 15 16] - 45
    [17 18 19 20] - 74
    [21 22 23 24] - 90

The left represents the temporal information found by the integral at each activation.  The right represents the
total of all the values found, which is also our integral dimension's timeline values.  Right now, we are
forcing the two dimensions to operate in good timing by using a slower frequency than the target dimension - as
we step forward we will begin to work with the concept of _resonant frequency_ calculation.  For now, it's as
obvious as it sounds - dimensions can activate at a sympathetic frequency to each other, facilitating logical
ordered calculation across time.