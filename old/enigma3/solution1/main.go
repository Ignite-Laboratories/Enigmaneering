package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/host/window"
)

func main() {
	for i := 0; i < 7; i++ {
		window.Create(std.XY[int]{X: 640, Y: 480})
	}
	core.Impulse.StopWhen(window.StopPotential)
	core.Impulse.Spark()
}
