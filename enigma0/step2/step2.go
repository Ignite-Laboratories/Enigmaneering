package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"sync"
	"time"
)

type Clock struct {
	ID      uint64
	Period  int
	Beat    int
	Kernels []Kernel
}

func NewClock(period int) Clock {
	return Clock{
		ID:     core.NextID(),
		Period: period,
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

		wg.Add(len(c.Kernels))
		for _, k := range c.Kernels {
			ctx.Kernel = k
			go k.Execute(ctx)
		}
		wg.Wait()

		c.Beat++
		if c.Beat >= c.Period {
			c.Beat = 0
		}
	}
}

func (c *Clock) AddKernel(action func(ctx Context), potential func(ctx Context) bool) {
	ap := &actionPotential{
		id:        core.NextID(),
		action:    action,
		potential: potential,
	}
	c.Kernels = append(c.Kernels, ap)
}

// Kernel is how others interface with the project as the public API and should be exported, thus capitalized
type Kernel interface {
	Execute(ctx Context)
	GetID() uint64
}

// actionPotential is only used by the clock for execution and shouldn't be exported, thus lowercase
type actionPotential struct {
	id        uint64
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
	return ap.id
}

type Context struct {
	Moment    time.Time
	Beat      int
	Clock     *Clock
	Kernel    Kernel
	waitGroup *sync.WaitGroup
}

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
	fmt.Printf("Action #%d - Beat %d\n", ctx.Kernel.GetID(), ctx.Beat)
	time.Sleep(1 * time.Second)
}
