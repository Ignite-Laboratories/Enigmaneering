package main

import (
	"enigma1/solution2"
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E1S2.0 - The Delta Encoder

This performs a single round of synthetic approximation and then prints the binary target,
approximation, and delta values.
*/

var indexWidth = 22
var patternWidth = 3

func main() {
	target := tiny.Synthesize.RandomPhrase(indexWidth)
	bestPattern, bestDelta := common.Approximate(patternWidth, target)

	fmt.Println(target.StringBinary())
	fmt.Println(bestPattern.StringBinary())
	fmt.Println(bestDelta.Text(2))
}
