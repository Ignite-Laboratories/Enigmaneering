package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
	"github.com/ignite-laboratories/core/temporal"
	"time"
)

// Observe the last impulse of the core impulse engine
var observer = temporal.NewObservation[core.Runtime](core.Impulse, condition.Always, false, &core.Impulse.Last)

func main() {
	// Print the timeline every 2 seconds
	loopFreq := 0.5
	core.Impulse.Loop(printTimeline, condition.Frequency(&loopFreq), false)

	// Lower the impulse frequency to 4hz
	core.Impulse.MaxFrequency = 4

	// Make it so
	core.Impulse.Spark()
}

func printTimeline(ctx core.Context) {
	// Copy the timeline data
	observer.Mutex.Lock()
	data := make([]temporal.Data[core.Runtime], len(observer.Timeline))
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
