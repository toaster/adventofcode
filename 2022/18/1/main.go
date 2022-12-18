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
	r := aocmath.Range{Start: math.MinInt64, End: math.MaxInt64}
	surface := 0
	for p := range points {
		for _, n := range p.Neighbours(r, r, r) {
			if !points[n] {
				surface++
			}
		}
	}
	fmt.Println(surface)
}
