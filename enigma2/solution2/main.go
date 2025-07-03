package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
	"math/big"
)

/**
E2S2 - The Midpoint Averager

This synthesizes random data and takes the average of how many bits it drops relative to
the midpoint of its containing index.
*/

var bitLength = 1024
var iterations = 1 << 12

func main() {
	average := 0
	for i := 0; i < iterations; i++ {
		data := tiny.Synthesize.RandomPhrase(bitLength / 8)
		midpoint := tiny.Synthesize.Midpoint(bitLength)
		delta := new(big.Int).Sub(data.AsBigInt(), midpoint.AsBigInt())

		average += bitLength - len(delta.Text(2))
	}
	average /= iterations
	fmt.Printf("Average Bit Drop: %d\n", average)
}
