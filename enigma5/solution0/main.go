package main

import (
	"enigma2/solution0/host"
	"fmt"
	"github.com/ignite-laboratories/core"
)

func main() {
	for core.Alive {
		fmt.Println(host.GetCoordinates())
	}
}
