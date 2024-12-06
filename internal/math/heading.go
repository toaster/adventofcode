package math

// Heading directions
const (
	North Heading = '^'
	East  Heading = '>'
	South Heading = 'v'
	West  Heading = '<'
)

// Heading is a direction someone might face on a two dimensional map (e.g., a Plan2D).
type Heading rune

// Facing returns the point that one looks at with this heading and the given position.
func (h Heading) Facing(pos Point2D) Point2D {
	switch h {
	case North:
		return pos.SubtractXY(0, 1)
	case East:
		return pos.AddXY(1, 0)
	case South:
		return pos.AddXY(0, 1)
	case West:
		return pos.SubtractXY(1, 0)
	}
	return pos
}

func (h Heading) String() string {
	return string(h)
}

// TurnRight returns the heading that would result in a turn right of 90“.
func (h Heading) TurnRight() Heading {
	switch h {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}
	return h
}
