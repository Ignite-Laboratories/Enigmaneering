package main

import (
	"fmt"
	"sync"
)

type Clock struct {
	LoopPeriod int
	Beat       int
	Actions    []func(int, int)
}

var Alive = true
var waitGroup *sync.WaitGroup

func main() {
	waitGroup = &sync.WaitGroup{}
	clock := Clock{
		LoopPeriod: 1024,
		Actions: []func(int, int){
			Action,
			Action,
			Action,
			Action,
			Action,
		},
	}

	for Alive {
		waitGroup.Add(len(clock.Actions))
		for i, action := range clock.Actions {
			if Potential(clock.Beat) {
				go action(i, clock.Beat)
			}
		}
		waitGroup.Wait()

		clock.Beat++
		if clock.Beat >= clock.LoopPeriod {
			clock.Beat = 0
		}
	}
}

func Potential(beat int) bool {
	if beat == 0 {
		return true
	}
	waitGroup.Done()
	return false
}

func Action(id int, beat int) {
	fmt.Printf("Action #%d - Beat #%d\n", id, beat)
	waitGroup.Done()
}
