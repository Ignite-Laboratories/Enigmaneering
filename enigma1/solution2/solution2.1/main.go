package main

import (
	"enigma1/solution2"
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E1S2.1 - The Bit Drop Averager

This performs a single round of synthetic approximation and then calculates the bit drop - then averages
the results of this operation over many rounds to calculate the average bit drop.
*/

var indexWidth = 22
var patternWidth = 4
var cycles = 1024

func main() {
	count := 0
	for i := 0; i < cycles; i++ {
		target := tiny.Synthesize.RandomPhrase(indexWidth)
		_, bestDelta := common.Approximate(patternWidth, target)

		bitDrop := len(target.StringBinary()) - len(bestDelta.Text(2)) - patternWidth
		count += bitDrop
	}
	fmt.Println("Average bit drop: ", count/cycles)
}
