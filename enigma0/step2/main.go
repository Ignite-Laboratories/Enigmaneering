package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

var printer = core.Impulse.Stimulate(PrintParity, when.Always, true)

func main() {
	core.Impulse.Loop(Toggle, when.Always, false)
	core.Impulse.MaxFrequency = 4
	core.Impulse.Spark()
}

func Toggle(ctx core.Context) {
	if printer.Muted {
		fmt.Printf("[%d] Un-muting\n", ctx.Beat)
	} else {
		fmt.Printf("[%d] Muting\n", ctx.Beat)
	}
	printer.Muted = !printer.Muted
	time.Sleep(time.Second * 2)
}

func PrintParity(ctx core.Context) {
	if ctx.Beat%2 == 0 {
		fmt.Printf("[%d] - Even\n", ctx.Beat)
	} else {
		fmt.Printf("[%d] - Odd\n", ctx.Beat)
	}
}
