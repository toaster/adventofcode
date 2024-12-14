package math

import math2 "math"

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
			l.incline = math2.Inf(-1)
		} else {
			l.incline = math2.Inf(1)
		}
		l.y0 = math2.NaN()
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
	if math2.IsInf(l.incline, 0) {
		crossX = float64(l.A.X)
	} else if math2.IsInf(other.incline, 0) {
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
	if math2.IsInf(l.incline, 0) {
		return math2.NaN()
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
