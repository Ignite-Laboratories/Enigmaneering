package main

import (
	"github.com/ignite-laboratories/tiny"
	"math/big"
)

/**
E1S3 - The Refinement Pathway Generator

This generates a refinement pathway for a randomly generated source, then prints and compares it
against the target binary information.
*/

type Approximation struct {
	Signature tiny.Phrase
	Target    tiny.Phrase
	Upper     tiny.Phrase
	Lower     tiny.Phrase
	Delta     *big.Int
}

var patternWidth = 12
var indexWidth = 44

func main() {
	target := tiny.Synthesize.RandomPhrase(indexWidth)
}

func (a Approximation) Refine(width int) {

}
