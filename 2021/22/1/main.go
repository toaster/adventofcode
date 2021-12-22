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
			cube: &cuboid{
				frontBottomLeft: math.Point3D{X: xr[0], Y: yr[0], Z: zr[0]},
				backTopRight:    math.Point3D{X: xr[1], Y: yr[1], Z: zr[1]},
			},
			state: io.ParseBool(s[0]),
		})
	}

	initRegion := &cuboid{
		frontBottomLeft: math.Point3D{X: -50, Y: -50, Z: -50},
		backTopRight:    math.Point3D{X: 50, Y: 50, Z: 50},
	}
	cubes := map[math.Point3D]bool{}
	for _, step := range steps {
		if !isInside(step.cube.frontBottomLeft, initRegion) {
			continue
		}
		for x := step.cube.frontBottomLeft.X; x <= step.cube.backTopRight.X; x++ {
			for y := step.cube.frontBottomLeft.Y; y <= step.cube.backTopRight.Y; y++ {
				for z := step.cube.frontBottomLeft.Z; z <= step.cube.backTopRight.Z; z++ {
					p := math.Point3D{X: x, Y: y, Z: z}
					if step.state {
						cubes[p] = true
					} else {
						delete(cubes, p)
					}
				}
			}
		}
	}
	fmt.Println(len(cubes))
}

func isInside(p math.Point3D, cube *cuboid) bool {
	return !p.IsLessThan(cube.frontBottomLeft) && !p.IsGreaterThan(cube.backTopRight)
}

type cuboid struct {
	frontBottomLeft math.Point3D
	backTopRight    math.Point3D
}

type rebootStep struct {
	cube  *cuboid
	state bool
}
