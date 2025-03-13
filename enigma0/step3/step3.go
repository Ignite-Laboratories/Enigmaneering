package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"sync"
	"time"
)

var desiredDuration = time.Second / 100

func main() {
	clock := NewClock(1024)
	clock.AddKernel(Action, Potential)
	clock.AddKernel(Action, Potential)
	clock.AddKernel(Action, Potential)
	clock.AddKernel(Action, Potential)
	clock.AddKernel(Action, Potential)
	clock.Start()
}

func Potential(ctx Context) bool {
	return true
}

func Action(ctx Context) {
	fmt.Printf("Action #%d - Beat #%d\n", ctx.Kernel.GetID(), ctx.Beat)
	time.Sleep(1 * time.Second)
}

type Clock struct {
	ID         uint64
	BPL        int
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
	beat := 0
	lastNow := time.Now()

	// Record the number of beats per loop
	beatCount := 0
	beatCountStart := lastNow

	for core.Alive {
		// Now should always be saved as the first step of the loop!
		now := time.Now()
		beatCount++

		var ctx Context
		ctx.Moment = now
		ctx.Delta = time.Duration(0)
		ctx.Beat = beat
		ctx.Clock = c
		ctx.waitGroup = &wg

		// Every 1024 beats...
		if beatCount > 1024 {
			// ...calculate the clock rate
			elapsed := now.Sub(beatCountStart).Seconds()
			c.BPL = int(float64(beatCount) / elapsed)

			// ...and reset the counter
			beatCount = 0
			beatCountStart = now
		}

		for _, k := range c.kernels {
			wg.Add(1)
			ctx.Kernel = k
			go k.Execute(ctx)
		}
		wg.Wait()

		beat++
		if beat >= c.LoopPeriod {
			beat = 0
		}
		lastNow = now
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

type Kernel interface {
	Execute(ctx Context)
	GetID() uint64
	IsExecuting() bool
}

type actionPotential struct {
	ID uint64

	// lastBeatMoment is the last beat's moment from which this actionPotential was activated.
	lastBeatMoment time.Time

	// lastCompletion is the last moment in time this actionPotential finished execution.
	lastCompletion time.Time

	executing bool
	action    func(ctx Context)
	potential func(ctx Context) bool
}

func (ap *actionPotential) Execute(ctx Context) {
	if !ap.executing && ap.potential(ctx) {
		ap.executing = true

		// If we are not on the first beat...
		if !ap.lastTrigger.IsZero() {
			// ...Calculate how long has passed since the last beat
			ctx.Delta = ctx.Moment.Sub(ap.lastTrigger)
			ctx.LastDuration = ap.lastCompletion.Sub(ap.lastTrigger)
		}

		go func() {
			ap.action(ctx)

			// ...and record when this kernel finished executing
			ap.lastCompletion = time.Now()

			ap.executing = false
		}()

		// Temporally regress this beat's contextual information for historical reference
		ap.lastBeatMoment = ctx.Moment
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
	Moment       time.Time
	Delta        time.Duration
	LastDuration time.Duration
	Beat         int
	Clock        *Clock
	Kernel       Kernel
	waitGroup    *sync.WaitGroup
}

var timeline sync.Map

type TemporalFragment struct {
	lastNow      time.Time
	lastDuration time.Duration
	delta        time.Duration
}

func NewTemporalFragment(now time.Time, duration time.Duration) TemporalFragment {
	return TemporalFragment{
		lastNow:      now,
		lastDuration: duration,
		delta:        duration - desiredDuration,
	}
}
