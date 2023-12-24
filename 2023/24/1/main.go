package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	var trajectories []*math.Ray2D
	for _, line := range io.ReadLines() {
		rawPosAndVelocity := strings.Split(line, " @ ")
		posXYZ := io.ParseInts(rawPosAndVelocity[0], ", ")
		velocityXYZ := io.ParseInts(rawPosAndVelocity[1], ", ")
		trajectories = append(trajectories, &math.Ray2D{LineSegment2D: math.LineSegment2D{A: math.Point2D{X: posXYZ[0], Y: posXYZ[1]}, B: math.Point2D{X: posXYZ[0] + velocityXYZ[0], Y: posXYZ[1] + velocityXYZ[1]}}})
	}
	targetArea := &math.Rectangle2D{TopLeft: math.Point2D{X: 200000000000000, Y: 200000000000000}, BottomRight: math.Point2D{X: 400000000000000, Y: 400000000000000}}
	count := 0
	for _, a := range trajectories {
		for _, b := range trajectories {
			if a == b {
				continue
			}

			if a.Crosses(b, targetArea) {
				count++
			}
		}
	}
	fmt.Println(count)
}
