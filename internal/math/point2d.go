package math

// Point2D is a two-dimensional point.
type Point2D Vector2D

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

// Vector2D is a two-dimensional vector used for position or velocity.
type Vector2D struct {
	X int
	Y int
}
