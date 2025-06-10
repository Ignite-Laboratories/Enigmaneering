package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E1S2 - The Rectified Sinusoidal Waveform Printer

This alternates between incrementing and decrementing between sub-indices.

The first value in the output is the numeric representation, the second is the
delta value for that sub-index.
*/

var bitWidth = 10
var subdivisions = 7

func main() {
	oversizeValue := 1 << bitWidth
	subdivisionMax := oversizeValue / (subdivisions + 1)
	value := oversizeValue - 1

	for i := subdivisions; i >= 0; i-- {
		odd := i%2 == 1

		if odd {
			for ii := 0; ii < subdivisionMax; ii++ {
				bits := tiny.From.Number(ii)
				fmt.Printf("[%6d][%6d:%d]%v\n", value, ii, i, bits)
				value--
			}
		} else {
			for ii := subdivisionMax - 1; ii >= 0; ii-- {
				bits := tiny.From.Number(ii)
				fmt.Printf("[%6d][%6d:%d]%v\n", value, ii, i, bits)
				value--
			}
		}
	}
}
