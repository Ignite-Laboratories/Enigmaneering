package main

import "fmt"

type Clock struct {
	Period int
	Beat   int
	Action func(int)
}

var Alive = true

func main() {
	clock := Clock{Period: 1000, Action: Action}

	for Alive {

		if Potential(clock.Beat) {
			go clock.Action(clock.Beat)
		}

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
	return false
}

func Action(beat int) {
	fmt.Printf("Beat %d\n", beat)
}
