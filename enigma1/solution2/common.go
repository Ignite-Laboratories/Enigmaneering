package common

import (
	"github.com/ignite-laboratories/tiny"
	"math/big"
)

/**
E1S2 - The Delta Encoder

This creates a single pattern approximation of a random target value and outputs the delta as a binary value.
*/

func Approximate(patternWidth int, target tiny.Phrase) (bestPattern tiny.Phrase, bestDelta *big.Int) {
	t := target.AsBigInt()

	bestDelta = t
	bestPattern = tiny.Synthesize.Zeros(target.BitLength())
	for i := 0; i <= (1<<patternWidth)-1; i++ {
		// Create the initial pattern bits
		bits := tiny.From.Number(i, patternWidth)

		// Synthesize the full pattern
		p := tiny.Synthesize.Pattern(target.BitLength(), bits...)

		// Get the delta value
		delta := new(big.Int).Sub(t, p.AsBigInt())
		if delta.CmpAbs(bestDelta) <= 0 {
			// Save off the best result
			bestPattern = p
			bestDelta = delta
		}
	}

	return bestPattern, bestDelta
}
