package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	directions := strings.Split(strings.TrimSpace(io.ReadAll()), ", ")
	start := math.Point2D{}
	loc := start
	heading := 0
	for _, direction := range directions {
		distance := io.ParseInt(direction[1:])
		switch direction[0] {
		case 'R':
			heading++
		case 'L':
			heading += 3
		}
		heading %= 4
		switch heading {
		case 0:
			loc = loc.SubtractXY(0, distance)
		case 1:
			loc = loc.AddXY(distance, 0)
		case 2:
			loc = loc.AddXY(0, distance)
		case 3:
			loc = loc.SubtractXY(distance, 0)
		}
	}
	fmt.Println(start.ManhattanDistance(loc))
}
