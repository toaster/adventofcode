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
			cub: &cuboid{
				frontBottomLeft: math.Point3D{X: xr[0], Y: yr[0], Z: zr[0]},
				backTopRight:    math.Point3D{X: xr[1], Y: yr[1], Z: zr[1]},
			},
			state: io.ParseBool(s[0]),
		})
	}

	var onCuboids []*cuboid
	var offCuboids []*cuboid
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
		count += onCuboid.size()
	}
	fmt.Println(count)
}

func apply(cub *cuboid, addingCuboids []*cuboid, subtractingCuboids []*cuboid) ([]*cuboid, []*cuboid) {
	additionalAddingCuboids := cub.nonIntersectingGroup(addingCuboids)

	var newSubtractingCuboids []*cuboid
	for _, subtractingCuboid := range subtractingCuboids {
		newSubtractingCuboids = append(newSubtractingCuboids, subtractingCuboid.nonIntersectingGroup(additionalAddingCuboids)...)
	}
	return append(addingCuboids, additionalAddingCuboids...), newSubtractingCuboids
}

type cuboid struct {
	frontBottomLeft math.Point3D
	backTopRight    math.Point3D
}

func (c *cuboid) String() string {
	return fmt.Sprintf("%v..%v", c.frontBottomLeft, c.backTopRight)
}

func (c *cuboid) equals(other *cuboid) bool {
	return c.frontBottomLeft == other.frontBottomLeft && c.backTopRight == other.backTopRight
}

func (c *cuboid) intersect(other *cuboid) *cuboid {
	leftX := math.MaxInt(c.frontBottomLeft.X, other.frontBottomLeft.X)
	rightX := math.MinInt(c.backTopRight.X, other.backTopRight.X)
	if leftX > rightX {
		return nil
	}

	bottomY := math.MaxInt(c.frontBottomLeft.Y, other.frontBottomLeft.Y)
	topY := math.MinInt(c.backTopRight.Y, other.backTopRight.Y)
	if bottomY > topY {
		return nil
	}

	frontZ := math.MaxInt(c.frontBottomLeft.Z, other.frontBottomLeft.Z)
	backZ := math.MinInt(c.backTopRight.Z, other.backTopRight.Z)
	if frontZ > backZ {
		return nil
	}

	return &cuboid{
		frontBottomLeft: math.Point3D{
			X: leftX,
			Y: bottomY,
			Z: frontZ,
		},
		backTopRight: math.Point3D{
			X: rightX,
			Y: topY,
			Z: backZ,
		},
	}
}

func (c *cuboid) isInside(p math.Point3D) bool {
	return !p.IsLessThan(c.frontBottomLeft) && !p.IsGreaterThan(c.backTopRight)
}

func (c *cuboid) nonIntersecting(other *cuboid) []*cuboid {
	if c.equals(other) {
		return nil
	}

	i := c.intersect(other)
	if i == nil {
		return []*cuboid{c}
	}

	var cuboids []*cuboid
	if c.frontBottomLeft.X < i.frontBottomLeft.X {
		cuboids = append(cuboids, &cuboid{
			frontBottomLeft: c.frontBottomLeft,
			backTopRight: math.Point3D{
				X: i.frontBottomLeft.X - 1,
				Y: c.backTopRight.Y,
				Z: c.backTopRight.Z,
			},
		})
	}
	if c.backTopRight.X > i.backTopRight.X {
		cuboids = append(cuboids, &cuboid{
			frontBottomLeft: math.Point3D{
				X: i.backTopRight.X + 1,
				Y: c.frontBottomLeft.Y,
				Z: c.frontBottomLeft.Z,
			},
			backTopRight: c.backTopRight,
		})
	}
	if c.frontBottomLeft.Y < i.frontBottomLeft.Y {
		cuboids = append(cuboids, &cuboid{
			frontBottomLeft: math.Point3D{
				X: i.frontBottomLeft.X,
				Y: c.frontBottomLeft.Y,
				Z: c.frontBottomLeft.Z,
			},
			backTopRight: math.Point3D{
				X: i.backTopRight.X,
				Y: i.frontBottomLeft.Y - 1,
				Z: c.backTopRight.Z,
			},
		})
	}
	if c.backTopRight.Y > i.backTopRight.Y {
		cuboids = append(cuboids, &cuboid{
			frontBottomLeft: math.Point3D{
				X: i.frontBottomLeft.X,
				Y: i.backTopRight.Y + 1,
				Z: c.frontBottomLeft.Z,
			},
			backTopRight: math.Point3D{
				X: i.backTopRight.X,
				Y: c.backTopRight.Y,
				Z: c.backTopRight.Z,
			},
		})
	}
	if c.frontBottomLeft.Z < i.frontBottomLeft.Z {
		cuboids = append(cuboids, &cuboid{
			frontBottomLeft: math.Point3D{
				X: i.frontBottomLeft.X,
				Y: i.frontBottomLeft.Y,
				Z: c.frontBottomLeft.Z,
			},
			backTopRight: math.Point3D{
				X: i.backTopRight.X,
				Y: i.backTopRight.Y,
				Z: i.frontBottomLeft.Z - 1,
			},
		})
	}
	if c.backTopRight.Z > i.backTopRight.Z {
		cuboids = append(cuboids, &cuboid{
			frontBottomLeft: math.Point3D{
				X: i.frontBottomLeft.X,
				Y: i.frontBottomLeft.Y,
				Z: i.backTopRight.Z + 1,
			},
			backTopRight: math.Point3D{
				X: i.backTopRight.X,
				Y: i.backTopRight.Y,
				Z: c.backTopRight.Z,
			},
		})
	}
	return cuboids
}

func (c *cuboid) nonIntersectingGroup(cuboids []*cuboid) []*cuboid {
	nonIntersecting := []*cuboid{c}
	for _, other := range cuboids {
		var newNonIntersecting []*cuboid
		for _, cub := range nonIntersecting {
			newNonIntersecting = append(newNonIntersecting, cub.nonIntersecting(other)...)
		}
		nonIntersecting = newNonIntersecting
	}
	return nonIntersecting
}

func (c *cuboid) size() int {
	if c == nil {
		return 0
	}

	return (c.backTopRight.X - c.frontBottomLeft.X + 1) * (c.backTopRight.Y - c.frontBottomLeft.Y + 1) * (c.backTopRight.Z - c.frontBottomLeft.Z + 1)
}

type rebootStep struct {
	cub   *cuboid
	state bool
}
