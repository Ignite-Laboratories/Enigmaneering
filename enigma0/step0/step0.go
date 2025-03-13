package main

import "fmt"

type Clock struct {
	LoopPeriod int
	Beat       int
	Action     func(int)
}

var Alive = true

func main() {
	clock := Clock{LoopPeriod: 1024, Action: Action}

	for Alive {

		if Potential(clock.Beat) {
			go clock.Action(clock.Beat)
		}

		clock.Beat++
		if clock.Beat >= clock.LoopPeriod {
			clock.Beat = 0
		}
	}
}

func Potential(beat int) bool {
	if beat == 42 {
		return true
	}
	return false
}

func Action(beat int) {
	fmt.Printf("Beat %d\n", beat)
}
