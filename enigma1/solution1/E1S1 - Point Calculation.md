# `E1S1 - Point Calculation`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### What if you want a more complex observation?

We'll evolve the last example a little bit by replacing the _observer_ with a _differential_ - in this
example it just increments a value -

    var incrementer = temporal.Differential(core.Impulse, when.Always, false, increment)

    var value = 0
    func increment(ctx core.Context) int {
        value++
        return value
    }

Here, the terminology has changed to a _differential_.  Differentials return some kind of function
that calculates a value when activated, the complexity of which is up to you.  The terminology is
intentional - the incrementer _dimension_ represents a differential data set, calculated point
by point.  The incrementer here isn't doing any  _differential calculation,_ per-say, but that's 
not because it _can't!_

Let's look at the output timeline -

    []
    [1 2 3 4 5]
    [2 3 4 5 6 7 8 9]
    [6 7 8 9 10 11 12]
    [9 10 11 12 13 14 15 16 17]
    [13 14 15 16 17 18 19 20]
    [17 18 19 20 21 22 23 24]
    [21 22 23 24 25 26 27 28]
    [26 27 28 29 30 31 32]

Right away you'll notice that there are _duplicate_ entries from our print function.  That's because
the dimension, itself, provides a _rolling window of observance_.  By default, all dimensions are
created with a 2 second observable window - but you can adjust that with its `Window` field.  The
print function isn't doing anything besides printing whatever is in the timeline buffer, so it
isn't able to filter out the duplicates - we'll handle that in the next solution.