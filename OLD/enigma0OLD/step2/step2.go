package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"sync"
	"time"
)

func main() {
	clock := NewClock(1000000)
	clock.AddKernel(Action, Potential)
	clock.AddKernel(Action, Potential)
	clock.AddKernel(Action, Potential)
	clock.AddKernel(Action, Potential)
	clock.AddKernel(Action, Potential)
	clock.Start()
}

func Potential(ctx Context) bool {
	// Run on every beat of the clock
	return true
}

func Action(ctx Context) {
	fmt.Printf("Action #%d - beat #%d\n", ctx.Kernel.GetID(), ctx.Beat)
	time.Sleep(1 * time.Second)
}

type Clock struct {
	ID         uint64
	LoopPeriod int
	Beat       int
	kernels    []Kernel
}

func NewClock(period int) Clock {
	return Clock{
		ID:         core.NextID(),
		LoopPeriod: period,
	}
}

func (c *Clock) Start() {
	var wg sync.WaitGroup

	for core.Alive {
		var ctx Context
		ctx.Moment = time.Now()
		ctx.Beat = c.Beat
		ctx.Clock = c
		ctx.waitGroup = &wg

		wg.Add(len(c.kernels))
		for _, k := range c.kernels {
			ctx.Kernel = k
			go k.Execute(ctx)
		}
		wg.Wait()

		c.Beat++
		if c.Beat >= c.LoopPeriod {
			c.Beat = 0
		}
	}
}

func (c *Clock) AddKernel(action func(ctx Context), potential func(ctx Context) bool) {
	ap := &actionPotential{
		ID:        core.NextID(),
		action:    action,
		potential: potential,
	}
	c.kernels = append(c.kernels, ap)
}

// Kernel is how others interface with the project as the public API and should be exported, thus capitalized
type Kernel interface {
	Execute(ctx Context)
	GetID() uint64
	IsExecuting() bool
}

// actionPotential is only used by the clock for execution and shouldn't be exported, thus lowercase
type actionPotential struct {
	ID        uint64
	executing bool
	action    func(ctx Context)
	potential func(ctx Context) bool
}

func (ap *actionPotential) Execute(ctx Context) {
	if !ap.executing && ap.potential(ctx) {
		ap.executing = true
		go func() {
			ap.action(ctx)
			ap.executing = false
		}()
	}
	ctx.waitGroup.Done()
}

func (ap *actionPotential) GetID() uint64 {
	return ap.ID
}

func (ap *actionPotential) IsExecuting() bool {
	return ap.executing
}

type Context struct {
	Moment    time.Time
	Beat      int
	Clock     *Clock
	Kernel    Kernel
	waitGroup *sync.WaitGroup
}
