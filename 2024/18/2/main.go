package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	var bytes []math.Point2D
	maxX := 0
	maxY := 0
	lines := io.ReadLines()
	after := io.ParseInt(lines[0])
	for _, line := range lines[2:] {
		b := io.ParseInts(line, ",")
		bytes = append(bytes, math.Point2D{X: b[0], Y: b[1]})
		if b[0] > maxX {
			maxX = b[0]
		}
		if b[1] > maxY {
			maxY = b[1]
		}
	}
	mem := math.NewPlan2D(maxX+1, maxY+1)
	for i := 0; i < after; i++ {
		mem = mem.AddObstacle(bytes[i])
	}
	for i := after; i < len(bytes); i++ {
		mem = mem.AddObstacle(bytes[i])
		distance := math.FindShortestDistance(
			func(pos math.Point2D) []math.Point2D {
				return mem.Neighbours(pos)
			},
			math.Point2D{},
			math.Point2D{X: maxX, Y: maxY},
		)
		if distance < 0 {
			fmt.Printf("%d,%d\n", bytes[i].X, bytes[i].Y)
			break
		}
	}
}
