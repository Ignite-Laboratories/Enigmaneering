package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
	"math"
	"math/big"
)

/**
E2S3E0 - The Breadcrumb Printer

This prints out the breadcrumb data from recursively midpointing a randomly generated value.
*/

var bitLength = 32

func main() {
	data := tiny.Synthesize.RandomPhrase(bitLength / 8)
	dataStr := data.StringBinary()

	fmt.Printf("%*s 🡨 Data\n", 3+paddedBitWidth+len(dataStr), dataStr)

	path := tiny.NewPhrase()
	delta := data.AsBigInt()

	// NOTE: We walk to -above- one bit remaining as the minimum footprint of
	// the delta is 1 bit and we cannot synthesize the midpoint of one bit.

	step := 0
	for i := bitLength; i > 1; i-- {
		midpoint := tiny.Synthesize.Midpoint(i)
		delta = new(big.Int).Sub(midpoint.AsBigInt(), delta)

		sign := tiny.Zero
		if delta.Sign() < 0 {
			sign = tiny.One
		}
		path = path.AppendBits(sign)
		delta = delta.Abs(delta) // KEY: We take the absolute value of the delta!

		fmt.Printf("(%*d) %v %v\n", paddedBitWidth, step, path.StringBinary(), delta.Text(2))
		step++
	}

	fmt.Printf(" ⬑ Step  Path ⬏ %*s\n", paddedBitWidth+len(dataStr)-13, "Delta ⬏")
}

// NOTE: This is just for pretty printing purposes =)
var paddedBitWidth = int(math.Floor(math.Log10(float64(bitLength)))) + 1
