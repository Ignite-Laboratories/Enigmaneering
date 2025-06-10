package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E1S1 - The Sawtooth Waveform Printer

This simply prints an entire index worth of bits in their numerical form.

You can subdivide the index in order to visualize the ramping waveform pattern it creates.
*/

var bitWidth = 8
var subdivisions = 3

func main() {
	oversizeValue := 1 << bitWidth
	subdivisionMax := oversizeValue / (subdivisions + 1)

	for i := subdivisions; i >= 0; i-- {
		for ii := subdivisionMax - 1; ii >= 0; ii-- {
			bits := tiny.From.Number(ii)
			value := ii + (subdivisionMax * i)
			fmt.Printf("[%6d][%6d:%d]%v\n", value, ii, i, bits)
		}
	}
}
