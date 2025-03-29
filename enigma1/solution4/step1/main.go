package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/host/mouse"
	"math"
)

func init() {
	mouse.Reaction(core.Impulse, &mouse.SampleRate, VelocityReaction)
}

func main() {
	core.Impulse.Spark()
}

func VelocityReaction(ctx core.Context, old std.Data[std.XY[int]], current std.Data[std.XY[int]]) {
	delta := current.Point.X - old.Point.X
	deltaAbs := math.Abs(float64(delta))
	if deltaAbs > 100 {
		fmt.Println("Slow down!!!")
	}
}
