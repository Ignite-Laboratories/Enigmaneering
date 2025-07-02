package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E1S0 - The Index Printer

This simply prints an entire index worth of bits using the provided bit width.

Side Quest: Set the bit width to something large and witness the behavior of binary counting in real time.
*/

var bitWidth = 4

func main() {
	maxValue := (1 << bitWidth) - 1
	for i := maxValue; i >= 0; i-- {
		bits := tiny.From.Number(i, bitWidth)

		switch {
		case i == 0:
			fmt.Printf("%v  Light Side\n", bits)
		case i == (maxValue/2)+1:
			fmt.Printf("%v  Midpoint\n", bits)
		case i == maxValue:
			fmt.Printf("%v  Dark Side\n", bits)
		default:
			fmt.Printf("%v\n", bits)
		}
	}
}
