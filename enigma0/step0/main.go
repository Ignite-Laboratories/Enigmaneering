package main

import (
	"fmt"
	"github.com/ignite-laboratories/core" // v0.0.2
	"time"
)

// This example is quite simple -

func main() {
	go core.DelayKill(time.Second * 5)
	for core.Alive {
		delta := time.Since(core.Inception)
		fmt.Printf("%d: %v\n", core.NextID(), delta)
		time.Sleep(time.Millisecond)
	}
}
