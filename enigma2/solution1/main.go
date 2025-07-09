package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E2S1 - The Index Subdivision Printer

This prints the binary pattern interval subdivisions of a requested index.

The output columns are:

  Interval   Pattern      Synthesized Point        Value     Point Distance
     (1)     [0 0 1]  [00100100 10010010 010010]  (599186)    Δ 599186
*/

var indexWidth = 44
var patternWidth = 8

func main() {
	lastPoint := tiny.NewPhrase()
	for i := 0; i < 1<<patternWidth; i++ {
		// Get the pattern interval's bits
		patternBits := tiny.From.Number(i, patternWidth)

		// Synthesize a point from the pattern interval
		point := tiny.Synthesize.Pattern(indexWidth, patternBits...)

		// Calculate the absolute delta to the synthetic point
		delta, _ := point.Subtract(lastPoint)

		// Store the point in order to calculate the next iteration's delta
		lastPoint = point

		// Print out the results
		fmt.Printf("(%*d) %v %v (%*d) Δ %d\n", tiny.GetBase10MaxWidth(patternWidth), i, patternBits, point.PadLeftToLength(indexWidth).Align(), tiny.GetBase10MaxWidth(indexWidth), point.Int(), delta.Int())
	}
}
