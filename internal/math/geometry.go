package math

import (
	"fmt"
	"math"
	"sort"

	"github.com/toaster/advent_of_code/internal/io"
)

// OrientationCount3D is the number of distinct orientations of a three-dimensional object which can
// rotate around each axis in 90° steps.
const OrientationCount3D = 24

// ParsePoint3D parses a Point3D from a set of values separated by the given separator.
func ParsePoint3D(input string, separator string) Point3D {
	values := io.ParseInts(input, separator)
	return Point3D{X: values[0], Y: values[1], Z: values[2]}
}

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

// Line2D describes a two-dimensional line.
type Line2D struct {
	LineSegment2D
	formulaComputed bool
	incline         float64
	y0              float64
}

// Crosses returns whether the line crosses another one in a specific rectangular area.
func (l *Line2D) Crosses(other *Line2D, area *Rectangle2D) bool {
	tlx := float64(area.TopLeft.X)
	tly := float64(area.TopLeft.Y)
	brx := float64(area.BottomRight.X)
	bry := float64(area.BottomRight.Y)
	crosses, crossX, crossY := l.crossPoint(other)
	return crosses && crossX >= tlx && crossX <= brx && crossY >= tly && crossY <= bry
}

func (l *Line2D) computeFormula() {
	if l.A.X == l.B.X {
		if l.A.Y < l.B.Y {
			l.incline = math.Inf(-1)
		} else {
			l.incline = math.Inf(1)
		}
		l.y0 = math.NaN()
	} else {
		l.incline = float64(l.B.Y-l.A.Y) / float64(l.B.X-l.A.X)
		l.y0 = -(float64(l.A.X))*l.incline + float64(l.A.Y)
	}
	l.formulaComputed = true
}

func (l *Line2D) crossPoint(other *Line2D) (bool, float64, float64) {
	if !l.formulaComputed {
		l.computeFormula()
	}

	// parallel
	if l.incline == other.incline {
		return false, 0, 0
	}

	// crossing at x == 0
	if l.y0 == other.y0 {
		return true, 0, l.y0
	}

	var crossX float64
	if math.IsInf(l.incline, 0) {
		crossX = float64(l.A.X)
	} else if math.IsInf(other.incline, 0) {
		crossX = float64(other.A.X)
	} else {
		crossX = (other.y0 - l.y0) / (l.incline - other.incline)
	}
	return true, crossX, l.yAtX(crossX)
}

func (l *Line2D) yAtX(x float64) float64 {
	if !l.formulaComputed {
		l.computeFormula()
	}
	if math.IsInf(l.incline, 0) {
		return math.NaN()
	}
	return l.incline*x + l.y0
}

// LineSegment2D describes a two-dimensional line segment.
type LineSegment2D struct {
	A Point2D
	B Point2D
}

// LinkedLineSegment2D is a linked LineSegment2D, i.e. usable for a linked list like a contour.
type LinkedLineSegment2D struct {
	LineSegment2D
	Previous *LinkedLineSegment2D
	Next     *LinkedLineSegment2D
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

// Ray2D represents a two-dimensional ray, i.e. a Line2D where point A is the starting point and the
// ray goes into direction of point B.
type Ray2D Line2D

// Crosses returns whether the ray crosses another one in a specific rectangular area.
func (r *Ray2D) Crosses(other *Ray2D, area *Rectangle2D) bool {
	crosses, crossX, crossY := (*Line2D)(r).crossPoint((*Line2D)(other))
	if !crosses {
		return false
	}

	if crossX < float64(area.TopLeft.X) || crossX > float64(area.BottomRight.X) || crossY < float64(area.TopLeft.Y) || crossY > float64(area.BottomRight.Y) {
		return false
	}

	return r.reaches(crossX, crossY) && other.reaches(crossX, crossY)
}

func (r *Ray2D) reaches(pX float64, pY float64) bool {
	if r.B.X > r.A.X && pX > float64(r.A.X) {
		return true
	}
	if r.B.Y > r.A.Y && pY > float64(r.A.Y) {
		return true
	}
	if r.B.X < r.A.X && pX < float64(r.A.X) {
		return true
	}
	if r.B.Y < r.A.Y && pY < float64(r.A.Y) {
		return true
	}

	return false
}

// Rectangle2D describes a two-dimensional rectangle.
type Rectangle2D struct {
	TopLeft     Point2D
	BottomRight Point2D
}

// Contains returns whether the Rectangle2D completely contains the other Rectangle2D.
func (r Rectangle2D) Contains(other Rectangle2D) bool {
	return other.TopLeft.X >= r.TopLeft.X && other.BottomRight.X <= r.BottomRight.X &&
		other.TopLeft.Y >= r.TopLeft.Y && other.BottomRight.Y <= r.BottomRight.Y
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
