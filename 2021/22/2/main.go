package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	var steps []*rebootStep
	for _, line := range io.ReadLines() {
		s := strings.Split(line, " ")
		r := strings.Split(s[1], ",")
		xr := io.ParseInts(r[0][2:], "..")
		yr := io.ParseInts(r[1][2:], "..")
		zr := io.ParseInts(r[2][2:], "..")
		steps = append(steps, &rebootStep{
			cub: &math.Cuboid{
				FrontBottomLeft: math.Point3D{X: xr[0], Y: yr[0], Z: zr[0]},
				BackTopRight:    math.Point3D{X: xr[1], Y: yr[1], Z: zr[1]},
			},
			state: io.ParseBool(s[0]),
		})
	}

	var onCuboids []*math.Cuboid
	var offCuboids []*math.Cuboid
	for _, step := range steps {
		cub := step.cub
		if step.state {
			onCuboids, offCuboids = apply(cub, onCuboids, offCuboids)
		} else {
			offCuboids, onCuboids = apply(cub, offCuboids, onCuboids)
		}
	}
	count := 0
	for _, onCuboid := range onCuboids {
		count += onCuboid.Size()
	}
	fmt.Println(count)
}

func apply(cub *math.Cuboid, addingCuboids []*math.Cuboid, subtractingCuboids []*math.Cuboid) ([]*math.Cuboid, []*math.Cuboid) {
	additionalAddingCuboids := cub.NonIntersectingGroup(addingCuboids)

	var newSubtractingCuboids []*math.Cuboid
	for _, subtractingCuboid := range subtractingCuboids {
		newSubtractingCuboids = append(newSubtractingCuboids, subtractingCuboid.NonIntersectingGroup(additionalAddingCuboids)...)
	}
	return append(addingCuboids, additionalAddingCuboids...), newSubtractingCuboids
}

type rebootStep struct {
	cub   *math.Cuboid
	state bool
}
