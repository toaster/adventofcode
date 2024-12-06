package math

// RoutePoint describes a point of a two-dimensional Route.
type RoutePoint struct {
	Point2D
	EnterHeading Heading
}

// Route describes a two-dimensional route consisting of RoutePoint.
type Route struct {
	isLoop  bool
	points  []RoutePoint
	passed  map[RoutePoint]bool
	visited map[Point2D]bool
}

// Add adds the Point2D with the Heading to the route.
func (r *Route) Add(p Point2D, h Heading) {
	if r.passed == nil {
		r.passed = map[RoutePoint]bool{}
	}
	if r.visited == nil {
		r.visited = map[Point2D]bool{}
	}

	rp := RoutePoint{Point2D: p, EnterHeading: h}
	if r.passed[rp] {
		r.isLoop = true
		return
	}

	r.visited[p] = true
	r.passed[rp] = true
	r.points = append(r.points, rp)
}

// IsLoop returns whether the route is a loop.
func (r *Route) IsLoop() bool {
	return r.isLoop
}

// PointCount returns the number of distinct points of the route.
func (r *Route) PointCount() int {
	return len(r.visited)
}

// Points returns all points covered by the route in arbitrary order.
func (r *Route) Points() []Point2D {
	points := make([]Point2D, len(r.visited))
	for p := range r.visited {
		points = append(points, p)
	}
	return points
}
