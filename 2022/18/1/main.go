package main

import (
	"fmt"
	"math"

	"github.com/toaster/advent_of_code/internal/io"
	aocmath "github.com/toaster/advent_of_code/internal/math"
)

func main() {
	points := map[aocmath.Point3D]bool{}
	for _, line := range io.ReadLines() {
		coords := io.ParseInts(line, ",")
		points[aocmath.Point3D{X: coords[0], Y: coords[1], Z: coords[2]}] = true
	}
	area := aocmath.Cuboid{
		FrontBottomLeft: aocmath.Point3D{X: math.MinInt64, Y: math.MinInt64, Z: math.MinInt64},
		BackTopRight:    aocmath.Point3D{X: math.MaxInt64, Y: math.MaxInt64, Z: math.MaxInt64},
	}
	surface := 0
	for p := range points {
		for _, n := range p.Neighbours(area) {
			if !points[n] {
				surface++
			}
		}
	}
	fmt.Println(surface)
}
