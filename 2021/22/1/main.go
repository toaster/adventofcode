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
			cuboid: &math.Cuboid{
				FrontBottomLeft: math.Point3D{X: xr[0], Y: yr[0], Z: zr[0]},
				BackTopRight:    math.Point3D{X: xr[1], Y: yr[1], Z: zr[1]},
			},
			state: io.ParseBool(s[0]),
		})
	}

	initRegion := &math.Cuboid{
		FrontBottomLeft: math.Point3D{X: -50, Y: -50, Z: -50},
		BackTopRight:    math.Point3D{X: 50, Y: 50, Z: 50},
	}
	cuboids := map[math.Point3D]bool{}
	for _, step := range steps {
		if !initRegion.IsInside(step.cuboid.FrontBottomLeft) {
			continue
		}
		for x := step.cuboid.FrontBottomLeft.X; x <= step.cuboid.BackTopRight.X; x++ {
			for y := step.cuboid.FrontBottomLeft.Y; y <= step.cuboid.BackTopRight.Y; y++ {
				for z := step.cuboid.FrontBottomLeft.Z; z <= step.cuboid.BackTopRight.Z; z++ {
					p := math.Point3D{X: x, Y: y, Z: z}
					if step.state {
						cuboids[p] = true
					} else {
						delete(cuboids, p)
					}
				}
			}
		}
	}
	fmt.Println(len(cuboids))
}

type rebootStep struct {
	cuboid *math.Cuboid
	state  bool
}
