package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
)

var incrementer = temporal.Calculation(core.Impulse, when.Always, false, increment)

func main() {
	// Print the timeline every second
	core.Impulse.Loop(printTimeline, when.Frequency(std.HardRef(1.0).Ref), false)

	// Lower the impulse frequency to 4hz
	core.Impulse.MaxFrequency = 4

	// Make it so
	core.Impulse.Spark()
}

var value = 0

func increment(ctx core.Context) int {
	value++
	return value
}

func printTimeline(ctx core.Context) {
	// Copy the timeline data
	incrementer.Mutex.Lock()
	data := make([]std.Data[int], len(incrementer.Timeline))
	copy(data, incrementer.Timeline)
	incrementer.Mutex.Unlock()

	// Get the point values
	values := make([]int, len(data))
	for i, v := range data {
		values[i] = v.Point
	}

	// Print the stats
	fmt.Printf("%v\n", values)
}
