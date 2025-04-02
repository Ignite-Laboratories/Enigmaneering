# `E1S2 - Duplicate Filtration`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### Leveraging Temporal Context

There is only a _very_ small tweak that allows duplicate filtration - walk the data and use
the provided temporal context to trim off _this_ activation's duplicates -

    var lastMoment time.Time
    
    func printTimeline(ctx core.Context) {
	    // Copy the timeline data

        ...
    
        // Trim duplicates
        trimCount := 0
        for _, v := range data {
            if v.Context.Moment.After(lastMoment) {
                break
            }
            trimCount++
        }
        data = data[trimCount:]
    
        ...

        // Get the point values
        // Print the stats
    }

Legitimately, _it's that simple_ with this kind of architecture =)

The data is always provided in a _temporally aligned_ fashion, meaning walking from left-to-right allows the system
to 'break' early when the desired condition is met.

Here's the new output -

    []
    [1 2 3 4]
    [5 6 7 8]
    [9 10 11 12 13]
    [14 15 16 17]
    [18 19 20 21]
    [22 23 24 25]
    [26 27 28]

Now, our print function will _only_ process new timeline entries and ignores anything it's already handled!

That's all fine and dandy, but the print function is excessively complex - let's make it simpler in the next solution =)