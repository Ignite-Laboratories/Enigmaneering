package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"time"
)

// Observe the last impulse of the core impulse engine
var observer = temporal.Observer(core.Impulse, when.Always, false, std.Target(&core.Impulse.Last))

func main() {
	// Print the timeline every half second
	core.Impulse.Loop(printTimeline, when.Frequency(std.HardRef(0.5).Ref), false)

	// Lower the impulse frequency to 4hz
	core.Impulse.MaxFrequency = 4

	// Make it so
	core.Impulse.Spark()
}

func printTimeline(ctx core.Context) {
	// Copy the timeline data
	observer.Mutex.Lock()
	data := make([]std.Data[core.Runtime], len(observer.Timeline))
	copy(data, observer.Timeline)
	observer.Mutex.Unlock()

	// Get the point values
	values := make([]time.Duration, len(data))
	for i, v := range data {
		values[i] = v.Point.Duration + v.Point.RefractoryPeriod
	}

	// Print the stats
	fmt.Printf("%v\n", values)
}
