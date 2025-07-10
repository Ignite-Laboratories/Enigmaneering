package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
)

/**
E2S2 - The Midpoint Averager

This walks every point in the provided index width and calculates the average of how many bits it drops relative to
the midpoint of its containing index.

NOTE: This is naturally restricted to a maximum index width of 64 bits - our practical infinity to even consider.
*/

var width = 16

func main() {
	for i := 0; i < 8; i++ {
		a := tiny.Synthesize.RandomBits(width).Align(64)[0]
		aV := a.Value()

		b := tiny.Synthesize.RandomBits(width).Align(64)[0]
		bV := b.Value()

		if aV > bV {
			result := aV - bV
			bits := tiny.NewMeasurementFromBits(tiny.From.Number(result, width)...)

			fmt.Printf("%v\t(%d)\n", a.StringBinary(), aV)
			fmt.Printf("%v\t(%d)\n", b.StringBinary(), bV)
			fmt.Println()
			fmt.Printf("%v\t(%d)\n", bits.StringBinary(), result)
		} else {
			result := bV - aV
			bits := tiny.NewMeasurementFromBits(tiny.From.Number(result, width)...)

			fmt.Printf("%v\t(%d)\n", b.StringBinary(), bV)
			fmt.Printf("%v\t(%d)\n", a.StringBinary(), aV)
			fmt.Println()
			fmt.Printf("%v\t(%d)\n", bits.StringBinary(), result)
		}

		fmt.Println()
	}
}
