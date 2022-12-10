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

// Point2D is a two-dimensional point.
type Point2D struct {
	X int
	Y int
}

// Add adds another two-dimensional coordinate to this one.
func (p Point2D) Add(other Point2D) Point2D {
	return Point2D{p.X + other.X, p.Y + other.Y}
}

// AddXY adds dimension values to this point.
// This works like Add but saves the Point2D initialization if you only have the single dimensions at hand.
func (p Point2D) AddXY(x, y int) Point2D {
	return Point2D{p.X + x, p.Y + y}
}

// IsAdjacent returns whether the point is adjacent to the other (including diagonal).
func (p Point2D) IsAdjacent(other Point2D) bool {
	return other.X > p.X-2 && other.X < p.X+2 && other.Y > p.Y-2 && other.Y < p.Y+2
}

// IsGreaterThan returns whether any dimension of this point is greater than the respective dimension of the other point.
func (p Point2D) IsGreaterThan(other Point2D) bool {
	return p.X > other.X || p.Y > other.Y
}

// IsLessThan returns whether any dimension of this point is less than the respective dimension of the other point.
func (p Point2D) IsLessThan(other Point2D) bool {
	return p.X < other.X || p.Y < other.Y
}

// ManhattanDistance returns the Manhattan Distance of the point to the other point.
func (p Point2D) ManhattanDistance(other Point2D) int {
	return AbsInt(p.X-other.X) + AbsInt(p.Y-other.Y)
}

// Subtract subtracts another two-dimensional coordinate from this one.
// The result is the coordinate of this point relative to the other one.
func (p Point2D) Subtract(other Point2D) Point2D {
	return Point2D{p.X - other.X, p.Y - other.Y}
}

// SubtractXY subtracts dimension values from this point.
// This works like Subtract but saves the Point2D initialization if you only have the single dimensions at hand.
func (p Point2D) SubtractXY(x, y int) Point2D {
	return Point2D{p.X - x, p.Y - y}
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

// IsGreaterThan returns whether any dimension of this point is greater than the respective dimension of the other point.
func (p Point3D) IsGreaterThan(other Point3D) bool {
	return p.X > other.X || p.Y > other.Y || p.Z > other.Z
}

// IsLessThan returns whether any dimension of this point is less than the respective dimension of the other point.
func (p Point3D) IsLessThan(other Point3D) bool {
	return p.X < other.X || p.Y < other.Y || p.Z < other.Z
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
