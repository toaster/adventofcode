package math

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// OrientationCount3D is the number of distinct orientations of a three-dimensional object which can
// rotate around each axis in 90° steps.
const OrientationCount3D = 24

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

// Adjacents returns the adjacent positions (including diagonal) of the point.
func (p Point2D) Adjacents() (adjacents []Point2D) {
	adjacents = append(adjacents, p.AddXY(-1, -1))
	adjacents = append(adjacents, p.AddXY(0, -1))
	adjacents = append(adjacents, p.AddXY(1, -1))
	adjacents = append(adjacents, p.AddXY(-1, 0))
	adjacents = append(adjacents, p.AddXY(1, 0))
	adjacents = append(adjacents, p.AddXY(-1, 1))
	adjacents = append(adjacents, p.AddXY(0, 1))
	adjacents = append(adjacents, p.AddXY(1, 1))
	return
}

// DirectAdjacents returns the adjacent positions (excluding diagonal) of the point.
func (p Point2D) DirectAdjacents() (adjacents []Point2D) {
	adjacents = append(adjacents, p.AddXY(0, -1))
	adjacents = append(adjacents, p.AddXY(-1, 0))
	adjacents = append(adjacents, p.AddXY(1, 0))
	adjacents = append(adjacents, p.AddXY(0, 1))
	return
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

// Neighbours returns the neighbours (excluding diagonal) of the point. The area is limited by the given Rectangle2D.
func (p Point2D) Neighbours(area Rectangle2D) (neighbours []Point2D) {
	if p.X > area.TopLeft.X {
		neighbours = append(neighbours, p.AddXY(-1, 0))
	}
	if p.X < area.BottomRight.X {
		neighbours = append(neighbours, p.AddXY(1, 0))
	}
	if p.Y > area.TopLeft.Y {
		neighbours = append(neighbours, p.AddXY(0, -1))
	}
	if p.Y < area.BottomRight.Y {
		neighbours = append(neighbours, p.AddXY(0, 1))
	}
	return
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
