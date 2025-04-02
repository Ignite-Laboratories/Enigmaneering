package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/host/window"
)

func main() {
	for i := 0; i < 7; i++ {
		window.Create()
	}
	core.Impulse.StopWhen(window.StopPotential)
	core.Impulse.Spark()
}
