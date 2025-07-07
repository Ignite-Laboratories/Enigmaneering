package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
	"math/big"
	"strings"
)

/**
E2S2 - The Midpoint Averager and Waveform Printer

This walks every point in the provided index width and calculates the average of how many bits it drops relative to
the index's midpoint.  While doing so, it emits out a visual representation of the number of bits that have dropped
from the midpointing operation.
*/

var indexWidth = 7

func main() {
	limit := 1 << indexWidth
	header := tiny.Print.IndexWidth(indexWidth)
	spaceBetween := 4

	fmt.Printf(" ⬐ Numeric Index Representation\n")
	fmt.Printf("%s%*s%s\n", header, spaceBetween, "", tiny.Print.IndexWidth(indexWidth))

	average := 0
	for i := limit - 1; i >= 0; i-- {
		data := tiny.Synthesize.Point(i, indexWidth)
		midpoint := tiny.Synthesize.Midpoint(indexWidth)
		delta := new(big.Int).Sub(data.AsBigInt(), midpoint.AsBigInt())
		deltaStr := new(big.Int).Abs(delta).Text(2)

		waveStr := "▏"
		waveSize := indexWidth - len(deltaStr)
		if waveSize > 0 {
			waveStr = strings.Repeat("█", waveSize)
		}

		fmt.Printf("|%*s|    |%-*s|\n", indexWidth, deltaStr, indexWidth, waveStr)

		average += indexWidth - len(deltaStr)
	}
	average /= limit

	fmt.Printf("%s%*s%s\n", header, spaceBetween, "", tiny.Print.IndexWidth(indexWidth))
	fmt.Printf("%*s ⬑ Bit Drop Waveform \n", indexWidth+2+spaceBetween, "")
	fmt.Printf("\n Average Bit Drop: %d\n", average)
}
