package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
	"math/big"
)

/**
E1S3E0 - The Breadcrumb Printer

This prints out the breadcrumb data from recursively midpointing a randomly generated value.
*/

var bitLength = 1024
var iterations = 1 << 12

func main() {
	data := tiny.Synthesize.RandomPhrase(bitLength / 8)

	for i := 0; i < iterations; i++ {
		midpoint := tiny.Synthesize.Midpoint(bitLength)
		delta := new(big.Int).Sub(data.AsBigInt(), midpoint.AsBigInt())

	}
	fmt.Printf("Average Bit Drop: %d\n", average)
}
