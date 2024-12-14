package math

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

// ParsePoint3D parses a Point3D from a set of values separated by the given separator.
func ParsePoint3D(input string, separator string) Point3D {
	values := io.ParseInts(input, separator)
	return Point3D{X: values[0], Y: values[1], Z: values[2]}
}

// Point3D is a three-dimensional point.
type Point3D struct {
	X int
	Y int
	Z int
}

// Add adds another three-dimensional coordinate to this one.
func (p Point3D) Add(other Point3D) Point3D {
	return Point3D{p.X + other.X, p.Y + other.Y, p.Z + other.Z}
}

// AddXYZ adds dimension values to this point.
// This works like Add but saves the Point3D initialization if you only have the single dimensions at hand.
func (p Point3D) AddXYZ(x, y, z int) Point3D {
	return Point3D{p.X + x, p.Y + y, p.Z + z}
}

// IsGreaterThan returns whether any dimension of this point is greater than the respective dimension of the other point.
func (p Point3D) IsGreaterThan(other Point3D) bool {
	return p.X > other.X || p.Y > other.Y || p.Z > other.Z
}

// IsLessThan returns whether any dimension of this point is less than the respective dimension of the other point.
func (p Point3D) IsLessThan(other Point3D) bool {
	return p.X < other.X || p.Y < other.Y || p.Z < other.Z
}

// Neighbours returns the neighbours (excluding diagonal) of the point. The area is limited by the given Cuboid.
func (p Point3D) Neighbours(area Cuboid) (neighbours []Point3D) {
	if p.X > area.FrontBottomLeft.X {
		neighbours = append(neighbours, p.AddXYZ(-1, 0, 0))
	}
	if p.X < area.BackTopRight.X {
		neighbours = append(neighbours, p.AddXYZ(1, 0, 0))
	}
	if p.Y > area.FrontBottomLeft.Y {
		neighbours = append(neighbours, p.AddXYZ(0, -1, 0))
	}
	if p.Y < area.BackTopRight.Y {
		neighbours = append(neighbours, p.AddXYZ(0, 1, 0))
	}
	if p.Z > area.FrontBottomLeft.Z {
		neighbours = append(neighbours, p.AddXYZ(0, 0, -1))
	}
	if p.Z < area.BackTopRight.Z {
		neighbours = append(neighbours, p.AddXYZ(0, 0, 1))
	}
	return
}

// String returns a string representation of the point.
func (p Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

// Subtract subtracts another three-dimensional coordinate from this one.
// The result is the coordinate of this point relative to the other one.
func (p Point3D) Subtract(other Point3D) Point3D {
	return Point3D{p.X - other.X, p.Y - other.Y, p.Z - other.Z}
}
