package math

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/util"
)

// ParsePlan2D parses a Plan2D from the given input.
func ParsePlan2D(lines []string, startMarker rune) Plan2D {
	plan := Plan2D{blocked: map[Point2D]bool{}, heading: Heading(startMarker)}
	for y, line := range lines {
		plan.height++
		if plan.width == 0 {
			plan.width = len(line)
		}
		for x, c := range line {
			switch c {
			case startMarker:
				plan.Start = Point2D{X: x, Y: y}
			case '#':
				plan.blocked[Point2D{X: x, Y: y}] = true
			}
		}
	}
	plan.rangeX = Range{End: plan.width - 1}
	plan.rangeY = Range{End: plan.height - 1}
	return plan
}

// Plan2D represents a two-dimensional area.
type Plan2D struct {
	Start Point2D

	blocked map[Point2D]bool
	heading Heading
	height  int
	rangeX  Range
	rangeY  Range
	width   int
}

// AddObstacle returns a new plan with an obstacle added at the given position
func (p Plan2D) AddObstacle(pos Point2D) Plan2D {
	p.blocked = util.CopyMap(p.blocked)
	p.blocked[pos] = true
	return p
}

// CountVisitedLocationsOfGuard follows the trace of the lab guard (https://adventofcode.com/2024/day/6)
// and returns the amount of distinct locations it reaches before moving out of the area
func (p Plan2D) CountVisitedLocationsOfGuard() any {
	return p.GetGuardRoute().PointCount()
}

// GetGuardRoute computes the route of the guard until it leaves the map or loops.
func (p Plan2D) GetGuardRoute() *Route {
	pos := p.Start
	heading := p.heading
	route := &Route{}
	for p.onMap(pos) && !route.IsLoop() {
		route.Add(pos, heading)
		next := heading.Facing(pos)
		for p.blocked[next] {
			heading = heading.TurnRight()
			next = heading.Facing(pos)
		}
		pos = next
	}
	return route
}

// Neighbours returns the reachable neighbours of a given point on the plan.
func (p Plan2D) Neighbours(point Point2D) (neighbours []Point2D) {
	if p.rangeY.Covers(point.Y) {
		if point.X > p.rangeX.Start {
			n := point.AddXY(-1, 0)
			if !p.blocked[n] {
				neighbours = append(neighbours, n)
			}
		}
		if point.X < p.rangeX.End {
			n := point.AddXY(1, 0)
			if !p.blocked[n] {
				neighbours = append(neighbours, n)
			}
		}
	}
	if p.rangeX.Covers(point.X) {
		if point.Y > p.rangeY.Start {
			n := point.AddXY(0, -1)
			if !p.blocked[n] {
				neighbours = append(neighbours, n)
			}
		}
		if point.Y < p.rangeY.End {
			n := point.AddXY(0, 1)
			if !p.blocked[n] {
				neighbours = append(neighbours, n)
			}
		}
	}
	return
}

// Print prints the plan on stdout.
func (p Plan2D) Print(isMarked map[Point2D]bool) {
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			pos := Point2D{X: x, Y: y}
			if p.blocked[pos] {
				fmt.Print("#")
			} else if isMarked[pos] {
				fmt.Print("O")
			} else if pos == p.Start {
				fmt.Print("S")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (p Plan2D) onMap(pos Point2D) bool {
	return pos.X >= 0 && pos.X < p.width && pos.Y >= 0 && pos.Y < p.height
}
