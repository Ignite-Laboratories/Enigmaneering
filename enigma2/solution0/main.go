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

var width = 7
var limit = 1 << width

func main() {
	fmt.Println(tiny.Print.IndexWidth(width))

	for i := limit - 1; i >= 0; i-- {
		printIndex(i, tiny.From.Number(i, width)...)
	}

	fmt.Println(tiny.Print.IndexWidth(width))
}

func printIndex(i int, bits ...tiny.Bit) {
	bitStr := tiny.Print.BetweenPipes(bits...)

	switch {
	case i == 0:
		fmt.Printf("%v ← Light Side\n", bitStr)
	case i == limit/2:
		fmt.Printf("%v ← Midpoint\n", bitStr)
	case i == limit-1:
		fmt.Printf("%v ← Dark Side\n", bitStr)
	default:
		fmt.Printf("%v\n", bitStr)
	}
}
