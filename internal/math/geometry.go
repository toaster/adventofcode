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
