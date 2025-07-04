package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E2S0 - The Index Printer

This simply prints an entire index worth of bits using the provided bit width.

Side Quest: Set the bit width to something large and witness the behavior of binary counting in real time.
*/

var indexBitWidth = 4
var maxValue = (1 << indexBitWidth) - 1

func main() {
	for i := maxValue; i >= 0; i-- {
		printIndex(i, tiny.From.Number(i, indexBitWidth)...)
	}
}

func printIndex(i int, bits ...tiny.Bit) {
	switch {
	case i == 0:
		fmt.Printf("%v ← Light Side\n", bits)
	case i == (1<<indexBitWidth)/2:
		fmt.Printf("%v ← Midpoint\n", bits)
	case i == maxValue:
		fmt.Printf("%v ← Dark Side\n", bits)
	default:
		fmt.Printf("%v\n", bits)
	}
}
