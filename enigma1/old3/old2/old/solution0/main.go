package main

import (
	"fmt"
	"github.com/ignite-laboratories/tiny"
	"math/big"
)

var powerMap = make(map[int]int)
var oddMap = make(map[string]int)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("Iteration %d -------------\n", i)
		Reduce(tiny.Synthesize.RandomPhrase(1048576, 8))
		fmt.Printf("Power Map: %v\n", powerMap)
		CheckOddMap()
	}
}

func CheckOddMap() {
	found := false
	for k, v := range oddMap {
		if v > 1 {
			found = true
			fmt.Printf("Found a non-unique number: %v\n", k)
		}
	}
	if !found {
		fmt.Println("No duplicate numbers found.")
	}
}

func Reduce(data tiny.Phrase) {
	str := data.StringBinary()
	power := 0

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '0' {
			power++
		} else {
			break
		}
	}

	powerMap[power] += 1

	val := new(big.Int)
	val.SetString(str, 2)

	if val.Cmp(big.NewInt(1)) == 0 {
		return
	}

	if power == 0 {
		oddMap[str] += 1
		val.Mul(val, big.NewInt(3))
		val.Add(val, big.NewInt(1))
	} else {
		val.Div(val, big.NewInt(int64(power*2)))
	}

	Reduce(tiny.NewPhraseFromString(val.Text(2)))
}
