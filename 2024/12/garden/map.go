package garden

import (
	"slices"

	"github.com/toaster/advent_of_code/internal/math"
)

// ParseMap creates a new Map from the input lines.
func ParseMap(lines []string) Map {
	return Map(math.ParsePlan2D(lines, ' '))
}

// Map is the map of a garden (https://adventofcode.com/2024/day/12).
type Map math.Plan2D

// ComputeFencingCosts computes the fencing costs of the garden.
func (m Map) ComputeFencingCosts() int {
	regions := m.computeRegions()
	costs := 0
	for _, region := range regions {
		area := len(region)
		perimeter := 0
		for p := range region {
			perimeter += len(math.Plan2D(m).Neighbours(p, math.Plan2DNeighboursWithExternal(), math.Plan2DNeighboursWithLimiter(func(p math.Point2D) bool { return !region[p] })))
		}
		costs += area * perimeter
	}
	return costs
}

// ComputeFencingCostsWithBulkDiscount computes the fencing costs of the garden applying the bulk discount.
func (m Map) ComputeFencingCostsWithBulkDiscount() int {
	regions := m.computeRegions()
	costs := 0
	for _, region := range regions {
		area := len(region)
		var sides []side
		for p := range region {
			for _, h := range math.Headings {
				if _, ok := region[h.Facing(p)]; !ok {
					newS := sideFacing(p, h)
					done := false
					for i, s := range sides {
						if s.Contains(newS) {
							done = true
							break
						}
						if extended, ok := newS.Extends(s); ok {
							done = true
							sides[i] = extended
							for j, s := range sides {
								if compacted, ok := s.Extends(extended); ok {
									sides[i] = compacted
									sides = slices.Delete(sides, j, j+1)
									break
								}
							}
							break
						}
					}
					if !done {
						sides = append(sides, newS)
					}
				}
			}
		}
		costs += area * len(sides)
	}
	return costs
}

func sideFacing(pos math.Point2D, heading math.Heading) side {
	switch heading {
	case math.North:
		return side{
			rect: math.Rectangle2D{
				TopLeft:     pos,
				BottomRight: pos.AddXY(1, 0),
			},
			outside: heading,
		}
	case math.East:
		return side{
			rect: math.Rectangle2D{
				TopLeft:     pos.AddXY(1, 0),
				BottomRight: pos.AddXY(1, 1),
			},
			outside: heading,
		}
	case math.South:
		return side{
			rect: math.Rectangle2D{
				TopLeft:     pos.AddXY(0, 1),
				BottomRight: pos.AddXY(1, 1),
			},
			outside: heading,
		}
	case math.West:
		return side{
			rect: math.Rectangle2D{
				TopLeft:     pos,
				BottomRight: pos.AddXY(0, 1),
			},
			outside: heading,
		}
	}
	panic("invalid heading")
}

func (m Map) computeRegions() map[math.Point2D]map[math.Point2D]bool {
	regions := map[math.Point2D]map[math.Point2D]bool{}
	belongsToRegion := map[math.Point2D]math.Point2D{}
	math.Plan2D(m).WithEachPoint(func(pos math.Point2D) {
		if _, ok := belongsToRegion[pos]; ok {
			return
		}

		region := map[math.Point2D]bool{pos: true}
		regions[pos] = region
		m.growRegion(pos, region, pos, belongsToRegion)
	})
	return regions
}

func (m Map) growRegion(pos math.Point2D, region map[math.Point2D]bool, regionPos math.Point2D, belongsToRegion map[math.Point2D]math.Point2D) {
	plan := math.Plan2D(m)
	for _, n := range plan.Neighbours(pos, math.Plan2DNeighboursWithLimiter(func(p math.Point2D) bool { return plan.TileAt(p) == plan.TileAt(pos) && !region[p] })) {
		region[n] = true
		belongsToRegion[n] = regionPos
		m.growRegion(n, region, regionPos, belongsToRegion)
	}
}

// NeighbourOption is a function to provide an option to Map.Neighbours.
type NeighbourOption func(*neighboursConfig)

type neighboursConfig struct {
	includeExternal bool
	limiter         func(math.Point2D) bool
}

type side struct {
	rect    math.Rectangle2D
	outside math.Heading
}

func (s side) Contains(other side) bool {
	return s.rect.Contains(other.rect) && s.outside == other.outside
}

func (s side) Extends(other side) (side, bool) {
	a := s
	b := other
	if a.rect.TopLeft.IsGreaterThan(b.rect.TopLeft) {
		a, b = b, a
	}
	if b.rect.TopLeft != a.rect.BottomRight || a.outside != b.outside {
		return side{}, false
	}
	widthA := a.rect.BottomRight.X - a.rect.TopLeft.X
	heightA := a.rect.BottomRight.Y - a.rect.TopLeft.Y
	widthB := b.rect.BottomRight.X - b.rect.TopLeft.X
	heightB := b.rect.BottomRight.Y - b.rect.TopLeft.Y
	if widthA != widthB && heightA != heightB {
		return side{}, false
	}

	return side{rect: math.Rectangle2D{TopLeft: a.rect.TopLeft, BottomRight: b.rect.BottomRight}, outside: a.outside}, true
}
