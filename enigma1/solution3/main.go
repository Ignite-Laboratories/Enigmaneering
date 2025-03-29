package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
)

func main() {
	var incrementer = temporal.Calculation(core.Impulse, when.Always, false, increment)
	temporal.Integration(core.Impulse, when.Frequency(std.HardRef(1.0).Ref), false, false, printTimeline, incrementer)
	core.Impulse.MaxFrequency = 4
	core.Impulse.Spark()
}

var value = 0

func increment(ctx core.Context) int {
	value++
	return value
}

func printTimeline(ctx core.Context, cache *any, data []std.Data[int]) int {
	total := 0
	for _, v := range data {
		total += v.Point
	}

	// Print the stats
	fmt.Printf("%v - %v\n", data, total)
	return total
}
