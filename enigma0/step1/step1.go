package main

import (
	"fmt"
	"sync"
)

type Clock struct {
	Period  int
	Beat    int
	Actions []func(int, int)
}

var Alive = true
var waitGroup *sync.WaitGroup

func main() {
	waitGroup = &sync.WaitGroup{}
	clock := Clock{
		Period: 1000,
		Actions: []func(int, int){
			Action,
			Action,
			Action,
			Action,
			Action,
		},
	}

	for Alive {
		waitGroup.Add(5)
		for i, action := range clock.Actions {
			if Potential(clock.Beat) {
				go action(i, clock.Beat)
			}
		}
		waitGroup.Wait()

		clock.Beat++
		if clock.Beat >= clock.Period {
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
	fmt.Printf("Action #%d - Beat %d\n", id, beat)
	waitGroup.Done()
}
