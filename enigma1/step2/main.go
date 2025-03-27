package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
	"github.com/ignite-laboratories/core/temporal"
	"time"
)

var incrementer = temporal.NewCalculation[int](core.Impulse, condition.Always, false, increment)

func main() {
	// Print the timeline every second
	loopFreq := 1.0
	core.Impulse.Loop(printTimeline, condition.Frequency(&loopFreq), false)

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

var lastMoment time.Time

func printTimeline(ctx core.Context) {
	// Copy the timeline data
	incrementer.Mutex.Lock()
	data := make([]temporal.Data[int], len(incrementer.Timeline))
	copy(data, incrementer.Timeline)
	incrementer.Mutex.Unlock()

	// Trim duplicates
	trimCount := 0
	for _, v := range data {
		if v.Context.Moment.After(lastMoment) {
			break
		}
		trimCount++
	}
	data = data[trimCount:]

	// Get the point values
	values := make([]int, len(data))
	for i, v := range data {

		values[i] = v.Point
		lastMoment = v.Context.Moment
	}

	// Print the stats
	fmt.Printf("%v\n", values)
}
