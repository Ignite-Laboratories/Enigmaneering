package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
)

func init() {
	go core.Impulse.Spark() // Make it so
}

func main() {
	fmt.Println("Press enter to trigger a stimulation")
	for core.Alive {
		// Press the enter key to read from stdin
		_, _ = fmt.Scanln()

		// Trigger a stimulation
		core.Impulse.Trigger(PrintParity, when.Always, true)
	}
}

func PrintParity(ctx core.Context) {
	fmt.Printf("Impulse moment %v\n", ctx.Moment)
}
