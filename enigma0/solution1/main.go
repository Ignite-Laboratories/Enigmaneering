package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

func main() {
	// Keep the beat counting for out a few seconds at a time
	core.Impulse.Loop(Hold, when.Always, false)

	// Print out even beats
	core.Impulse.Stimulate(func(ctx core.Context) {
		PrintParity(ctx, "Even")
	}, when.Beat.Even, false)

	// Print out odd beats
	core.Impulse.Stimulate(func(ctx core.Context) {
		PrintParity(ctx, "Odd")
	}, when.Beat.Odd, false)

	// Limit the impulse rate to 4hz
	core.Impulse.MaxFrequency = 4.0

	// Alternatively, increase impulse resistance
	//core.Impulse.Resistance = 800000000

	// Make it so
	core.Impulse.Spark()
}

func Hold(ctx core.Context) {
	time.Sleep(time.Second * 5)
}

func PrintParity(ctx core.Context, parity string) {
	fmt.Printf("%d - %v\n", ctx.Beat, parity)
}
