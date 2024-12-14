package math

import (
	"fmt"
	"strings"
)

// ParseCuboid parses a Cuboid from a set of coordinates separated by the given separator with the
// single dimensions separated by commas.
func ParseCuboid(input string, separator string) *Cuboid {
	edges := strings.Split(input, separator)
	a := ParsePoint3D(edges[0], ",")
	b := ParsePoint3D(edges[1], ",")
	if a.IsGreaterThan(b) {
		a, b = b, a
	}
	return &Cuboid{FrontBottomLeft: a, BackTopRight: b}
}

// Cuboid represents a cuboid, stupid.
type Cuboid struct {
	FrontBottomLeft Point3D
	BackTopRight    Point3D
}

func (c *Cuboid) String() string {
	return fmt.Sprintf("%v..%v", c.FrontBottomLeft, c.BackTopRight)
}

// Equals returns whether this cuboid is equal to the given one.
func (c *Cuboid) Equals(other *Cuboid) bool {
	return c.FrontBottomLeft == other.FrontBottomLeft && c.BackTopRight == other.BackTopRight
}

// Intersect returns the intersection of this cuboid with another one or nil if they don’t intersect.
func (c *Cuboid) Intersect(other *Cuboid) *Cuboid {
	leftX := MaxInt(c.FrontBottomLeft.X, other.FrontBottomLeft.X)
	rightX := MinInt(c.BackTopRight.X, other.BackTopRight.X)
	if leftX > rightX {
		return nil
	}

	bottomY := MaxInt(c.FrontBottomLeft.Y, other.FrontBottomLeft.Y)
	topY := MinInt(c.BackTopRight.Y, other.BackTopRight.Y)
	if bottomY > topY {
		return nil
	}

	frontZ := MaxInt(c.FrontBottomLeft.Z, other.FrontBottomLeft.Z)
	backZ := MinInt(c.BackTopRight.Z, other.BackTopRight.Z)
	if frontZ > backZ {
		return nil
	}

	return &Cuboid{
		FrontBottomLeft: Point3D{
			X: leftX,
			Y: bottomY,
			Z: frontZ,
		},
		BackTopRight: Point3D{
			X: rightX,
			Y: topY,
			Z: backZ,
		},
	}
}

// IsInside returns whether a point is part of this cuboid.
func (c *Cuboid) IsInside(p Point3D) bool {
	return !p.IsLessThan(c.FrontBottomLeft) && !p.IsGreaterThan(c.BackTopRight)
}

// NonIntersectingGroup returns the slice of cuboids that remains if you remove the intersecting
// parts of the given cuboids from this cuboid.
func (c *Cuboid) NonIntersectingGroup(cuboids []*Cuboid) []*Cuboid {
	nonIntersecting := []*Cuboid{c}
	for _, other := range cuboids {
		var newNonIntersecting []*Cuboid
		for _, cub := range nonIntersecting {
			newNonIntersecting = append(newNonIntersecting, cub.nonIntersecting(other)...)
		}
		nonIntersecting = newNonIntersecting
	}
	return nonIntersecting
}

// Size returns the size of this cuboid which is its volume.
func (c *Cuboid) Size() int {
	if c == nil {
		return 0
	}

	return (c.BackTopRight.X - c.FrontBottomLeft.X + 1) * (c.BackTopRight.Y - c.FrontBottomLeft.Y + 1) * (c.BackTopRight.Z - c.FrontBottomLeft.Z + 1)
}

func (c *Cuboid) nonIntersecting(other *Cuboid) []*Cuboid {
	if c.Equals(other) {
		return nil
	}

	i := c.Intersect(other)
	if i == nil {
		return []*Cuboid{c}
	}

	var cuboids []*Cuboid
	if c.FrontBottomLeft.X < i.FrontBottomLeft.X {
		cuboids = append(cuboids, &Cuboid{
			FrontBottomLeft: c.FrontBottomLeft,
			BackTopRight: Point3D{
				X: i.FrontBottomLeft.X - 1,
				Y: c.BackTopRight.Y,
				Z: c.BackTopRight.Z,
			},
		})
	}
	if c.BackTopRight.X > i.BackTopRight.X {
		cuboids = append(cuboids, &Cuboid{
			FrontBottomLeft: Point3D{
				X: i.BackTopRight.X + 1,
				Y: c.FrontBottomLeft.Y,
				Z: c.FrontBottomLeft.Z,
			},
			BackTopRight: c.BackTopRight,
		})
	}
	if c.FrontBottomLeft.Y < i.FrontBottomLeft.Y {
		cuboids = append(cuboids, &Cuboid{
			FrontBottomLeft: Point3D{
				X: i.FrontBottomLeft.X,
				Y: c.FrontBottomLeft.Y,
				Z: c.FrontBottomLeft.Z,
			},
			BackTopRight: Point3D{
				X: i.BackTopRight.X,
				Y: i.FrontBottomLeft.Y - 1,
				Z: c.BackTopRight.Z,
			},
		})
	}
	if c.BackTopRight.Y > i.BackTopRight.Y {
		cuboids = append(cuboids, &Cuboid{
			FrontBottomLeft: Point3D{
				X: i.FrontBottomLeft.X,
				Y: i.BackTopRight.Y + 1,
				Z: c.FrontBottomLeft.Z,
			},
			BackTopRight: Point3D{
				X: i.BackTopRight.X,
				Y: c.BackTopRight.Y,
				Z: c.BackTopRight.Z,
			},
		})
	}
	if c.FrontBottomLeft.Z < i.FrontBottomLeft.Z {
		cuboids = append(cuboids, &Cuboid{
			FrontBottomLeft: Point3D{
				X: i.FrontBottomLeft.X,
				Y: i.FrontBottomLeft.Y,
				Z: c.FrontBottomLeft.Z,
			},
			BackTopRight: Point3D{
				X: i.BackTopRight.X,
				Y: i.BackTopRight.Y,
				Z: i.FrontBottomLeft.Z - 1,
			},
		})
	}
	if c.BackTopRight.Z > i.BackTopRight.Z {
		cuboids = append(cuboids, &Cuboid{
			FrontBottomLeft: Point3D{
				X: i.FrontBottomLeft.X,
				Y: i.FrontBottomLeft.Y,
				Z: i.BackTopRight.Z + 1,
			},
			BackTopRight: Point3D{
				X: i.BackTopRight.X,
				Y: i.BackTopRight.Y,
				Z: c.BackTopRight.Z,
			},
		})
	}
	return cuboids
}
