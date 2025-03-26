package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"time"
)

func main() {
	// Run for 5 seconds
	go core.Shutdown(time.Second * 5)

	// While alive...
	for core.Alive {
		// Get the time since JanOS was initialized
		delta := time.Since(core.Inception)

		// Print out the result
		fmt.Printf("%d: %v\n", core.NextID(), delta)

		// Sleep for a millisecond
		time.Sleep(time.Millisecond)
	}
}
