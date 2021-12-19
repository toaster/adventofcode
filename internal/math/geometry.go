package math

import (
	"fmt"
	"sort"

	"github.com/toaster/advent_of_code/internal/io"
)

// OrientationCount3D is the number of distinct orientations of a three-dimensional object which can
// rotate around each axis in 90° steps.
const OrientationCount3D = 24

// TransformOrientation transforms a three-dimensional coordinate according to one of the possible
// three-dimensional orientations (see OrientationCount3D).
func TransformOrientation(p Point3D, orientation int) Point3D {
	switch orientation {
	case 0:
		return p
	case 1:
		return Point3D{p.X, p.Z, -p.Y}
	case 2:
		return Point3D{p.X, -p.Y, -p.Z}
	case 3:
		return Point3D{p.X, -p.Z, p.Y}
	case 4:
		return Point3D{-p.Z, p.Y, p.X}
	case 5:
		return Point3D{-p.X, p.Y, -p.Z}
	case 6:
		return Point3D{p.Z, p.Y, -p.X}
	case 7:
		return Point3D{p.Y, -p.X, p.Z}
	case 8:
		return Point3D{-p.X, -p.Y, p.Z}
	case 9:
		return Point3D{-p.Y, p.X, p.Z}
	case 10:
		return Point3D{p.Y, p.Z, p.X}
	case 11:
		return Point3D{-p.X, p.Z, p.Y}
	case 12:
		return Point3D{-p.Y, p.Z, -p.X}
	case 13:
		return Point3D{p.Z, -p.X, -p.Y}
	case 14:
		return Point3D{-p.X, -p.Z, -p.Y}
	case 15:
		return Point3D{-p.Z, p.X, -p.Y}
	case 16:
		return Point3D{p.Z, -p.Y, p.X}
	case 17:
		return Point3D{-p.Z, -p.Y, -p.X}
	case 18:
		return Point3D{-p.Y, -p.X, -p.Z}
	case 19:
		return Point3D{p.Y, p.X, -p.Z}
	case 20:
		return Point3D{-p.Y, -p.Z, p.X}
	case 21:
		return Point3D{p.Y, -p.Z, -p.X}
	case 22:
		return Point3D{-p.Z, -p.X, p.Y}
	case 23:
		return Point3D{p.Z, p.X, p.Y}
	default:
		io.ReportError("", fmt.Errorf("invalid 3D orientation: %d", orientation))
		return Point3D{}
	}
}

// Point2D is a two-dimensional point.
type Point2D struct {
	X int
	Y int
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

// String returns a string representation of the point.
func (p Point3D) String() string {
	return fmt.Sprintf("(%5d, %5d, %5d)", p.X, p.Y, p.Z)
}

// Subtract subtracts another three-dimensional coordinate from this one.
// The result is the coordinate of this point relative to the other one.
func (p Point3D) Subtract(other Point3D) Point3D {
	return Point3D{p.X - other.X, p.Y - other.Y, p.Z - other.Z}
}

// Sortable3DPoints is a sortable slice of Point3D.
type Sortable3DPoints []Point3D

var _ sort.Interface = (Sortable3DPoints)(nil)

func (s Sortable3DPoints) Len() int {
	return len(s)
}

func (s Sortable3DPoints) Less(i, j int) bool {
	a := s[i]
	b := s[j]
	return a.X < b.X || (a.X == b.X && (a.Y < b.Y || (a.Y == b.Y && a.Z < b.Z)))
}

func (s Sortable3DPoints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
