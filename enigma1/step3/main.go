package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/host"
)

type Coordinates struct {
	X int
	Y int
}

var mouser = temporal.NewCalculation[Coordinates](core.Impulse, condition.Always, false, GetCoordinates)
var analyzer = temporal.NewAnalysis[Coordinates, any, any](core.Impulse, condition.Always, false, PrintCoordinates, mouser)

func main() {
	core.Impulse.MaxFrequency = 16
	core.Impulse.Spark()
}

func PrintCoordinates(ctx core.Context, cache *any, data []temporal.Data[Coordinates]) any {
	points := make([]Coordinates, len(data))
	for i, v := range data {
		points[i] = v.Point
	}
	fmt.Println(points)
	return nil
}

func GetCoordinates(ctx core.Context) Coordinates {
	x, y, _ := host.Mouse.GetCoordinates()
	return Coordinates{x, y}
}
