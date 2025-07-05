package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E2S0 - The Index Printer

This simply prints an entire index worth of bits.

Side Quest: Set the bit width to something large and witness the behavior of binary counting in real time.
*/

var bitWidth = 4
var limit = (1 << bitWidth) - 1

func main() {
	for i := limit; i >= 0; i-- {
		printIndex(i, tiny.From.Number(i, bitWidth)...)
	}
}

func printIndex(i int, bits ...tiny.Bit) {
	switch {
	case i == 0:
		fmt.Printf("%v ← Light Side\n", bits)
	case i == (1<<bitWidth)/2:
		fmt.Printf("%v ← Midpoint\n", bits)
	case i == limit:
		fmt.Printf("%v ← Dark Side\n", bits)
	default:
		fmt.Printf("%v\n", bits)
	}
}
