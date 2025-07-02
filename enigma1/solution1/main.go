package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
	"math"
	"math/big"
)

/**
E1S1 - The Index Subdivision Printer

This prints the binary pattern subdivisions of a requested index

The output columns are:

 Subdivision Pattern     Synthesized Point         Value     Delta From Last Point
     (1)     [0 0 1]  [00100100 10010010 010010]  (599186)    Δ 599186

// TODO: Remove the dependency on big
*/

var patternWidth = 8
var indexWidth = 44

func main() {
	patternMax := 1 << patternWidth

	last := big.NewInt(0)
	for i := 0; i < patternMax; i++ {
		patternBits := tiny.From.Number(i, patternWidth)                   // Get the pattern's bits
		synthesized := tiny.Synthesize.Pattern(indexWidth, patternBits...) // Synthesize a point value from a pattern
		value := synthesized.AsBigInt()                                    // Convert it to a big.Int for arithmetic operations
		delta := new(big.Int).Sub(value, last)                             // Calculate the delta from the synthetic value

		// Print the result and store the created value for the next iteration
		maxDigits := int(math.Floor(math.Log10(float64(patternMax)))) + 1 // This just inline calculates how much to pad the value for printing
		fmt.Printf("(%*d) %v %v (%v) Δ %d\n", maxDigits, i, patternBits, synthesized, value, delta)
		last = value
	}
}
