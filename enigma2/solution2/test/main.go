package main

import (
	"fmt"
	"math"
)

/**
E2S2 - The Midpoint Averager

This walks every point in the provided index width and calculates the average of how many bits it drops relative to
the midpoint of its containing index.

NOTE: This is naturally restricted to a maximum index width of 64 bits - our practical infinity to even consider.
*/

var indexWidth = 9

func main() {
	//average := 0
	//cycles := 1 << indexWidth

	for i := 0; i < 64; i++ {
		indexWidth = i
		printHeader()
	}

	//for i := cycles - 1; i >= 0; i-- {
	//	data := tiny.Synthesize.Point(i, indexWidth)
	//	midpoint := tiny.Synthesize.Midpoint(indexWidth)
	//	delta := new(big.Int).Sub(data.AsBigInt(), midpoint.AsBigInt())
	//	deltaStr := delta.Text(2)
	//
	//	average += indexWidth - len(deltaStr)
	//}
	//average /= cycles
	//fmt.Printf("Average Bit Drop: %d\n", average)
}

func printHeader() {
	//fmt.Printf("|←%*v%s%*v→|\n", int(math.Max(math.Floor((float64(indexWidth)/2)-2), 0)), "", fmt.Sprintf("%*v", 1+int(indexWidth >= 4), indexWidth), int(math.Max(math.Ceil((float64(indexWidth)/2)-2), 0)), "")
	fmt.Printf("|←%*v%s%*v→|\n", int(math.Max(math.Floor((float64(indexWidth)/2)-2), 0)), "",
		fmt.Sprintf("%*v", 1+int(indexWidth >= 4), indexWidth),
		int(math.Max(math.Ceil((float64(indexWidth)/2)-2), 0)), "")
}
