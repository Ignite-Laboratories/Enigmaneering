package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
	"math/big"
)

/**
E1S3E1 - The 1-Bit Averager

This calculates the average distance in from the end that yields
an overall bit drop after many recursive midpointing operations.
*/

var bitLength = 128
var cycles = 1 << 12

func main() {
	average := 0
	for i := 0; i < cycles; i++ {
		average += recurse()
	}
	average /= cycles

	fmt.Printf("Average remaining steps after reaching a 1-bit delta: %d\n", average)
}

func recurse() int {
	data := tiny.Synthesize.RandomPhrase(bitLength / 8)

	path := tiny.NewPhrase()
	delta := data.AsBigInt()

	for i := bitLength; i > 1; i-- {
		midpoint := tiny.Synthesize.Midpoint(i)
		delta = new(big.Int).Sub(midpoint.AsBigInt(), delta)

		sign := tiny.Zero
		if delta.Sign() < 0 {
			sign = tiny.One
		}
		path = path.AppendBits(sign)
		delta = delta.Abs(delta)

		// Bailout from the recursion when a 1-bit value is reached.
		if len(delta.Text(2)) < 2 {
			return i
		}
	}

	return data.BitLength()
}
