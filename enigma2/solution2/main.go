package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
	"math/big"
)

/**
E2S2 - The Midpoint Averager

This walks every point in the provided index width and calculates the average of how many bits it drops relative to
the midpoint of its containing index.

NOTE: This is naturally restricted to a maximum index width of 64 bits - our practical infinity to even consider.
*/

var indexWidth = 16

func main() {
	average := 0
	cycles := 1 << indexWidth
	for i := cycles - 1; i >= 0; i-- {
		data := tiny.Synthesize.Point(i, indexWidth)
		midpoint := tiny.Synthesize.Midpoint(indexWidth)
		delta := new(big.Int).Sub(data.AsBigInt(), midpoint.AsBigInt())

		average += indexWidth - len(delta.Text(2))
	}
	average /= cycles
	fmt.Printf("Average Bit Drop: %d\n", average)
}
