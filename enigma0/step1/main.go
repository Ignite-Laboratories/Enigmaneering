package main

import (
	"fmt"
	"github.com/ignite-laboratories/core" // v0.0.4
	"time"
)

func main() {
	engine := core.NewEngine()
	engine.Loop(PrintAndWait, core.Activate.Always())
	engine.Start()
}

func PrintAndWait(ctx core.Context) {
	fmt.Println(ctx.Beat)
	time.Sleep(time.Second * 100)
}
