package math

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/util"
)

// ParsePlan2D parses a Plan2D from the given input.
func ParsePlan2D(lines []string, startMarker rune, options ...Plan2DOption) Plan2D {
	cfg := plan2DConfig{handleUnknownTile: func(Point2D, rune) {}}
	for _, option := range options {
		option(&cfg)
	}
	plan := Plan2D{
		blocked: map[Point2D]bool{},
		heading: Heading(startMarker),
		height:  len(lines),
		tiles:   map[Point2D]rune{},
	}
	for y, line := range lines {
		for x, c := range line {
			if plan.width == 0 {
				plan.width = len(line)
			}
			pos := Point2D{X: x, Y: y}
			plan.tiles[pos] = c
			switch c {
			case startMarker:
				plan.Start = pos
			case '#':
				plan.blocked[pos] = true
			case '.':
			default:
				cfg.handleUnknownTile(pos, c)
			}
		}
	}
	plan.rangeX = Range{End: plan.width - 1}
	plan.rangeY = Range{End: plan.height - 1}
	return plan
}

// Plan2DNeighboursWithExternal specifies to include neighbours off the map in Plan2D.Neighbours.
func Plan2DNeighboursWithExternal() Plan2DNeighbourOption {
	return func(config *plan2DNeighboursConfig) {
		config.includeExternal = true
	}
}

// Plan2DNeighboursWithLimiter specifies a limiter to use with Map.Neighbours.
func Plan2DNeighboursWithLimiter(l func(Point2D) bool) Plan2DNeighbourOption {
	return func(config *plan2DNeighboursConfig) {
		config.limiter = l
	}
}

// Plan2DPrintAt specifies a printer for printing tiles which are considered empty.
// The printer might return `false` if it does not print anything.
// In this case, the default for empty spaces (`.`) is printed.
func Plan2DPrintAt(print func(Point2D) bool) Plan2DPrintOption {
	return func(config *plan2DPrintConfig) {
		config.print = print
	}
}

// Plan2DWithUnknownTileHandler specifies a handler for parsing unknown tiles.
func Plan2DWithUnknownTileHandler(handleUnknownTile func(Point2D, rune)) Plan2DOption {
	return func(config *plan2DConfig) {
		config.handleUnknownTile = handleUnknownTile
	}
}

// Plan2D represents a two-dimensional area.
type Plan2D struct {
	Start Point2D

	blocked map[Point2D]bool
	heading Heading
	height  int
	rangeX  Range
	rangeY  Range
	tiles   map[Point2D]rune
	width   int
}

// AddObstacle returns a new plan with an obstacle added at the given position
func (p Plan2D) AddObstacle(pos Point2D) Plan2D {
	p.blocked = util.CopyMap(p.blocked)
	p.blocked[pos] = true
	return p
}

// Blocked returns whether the given position is blocked.
func (p Plan2D) Blocked(position Point2D) bool {
	return p.blocked[position]
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
// It allows to specify a limited (default: !blocked) and to include neighbours off the map.
func (p Plan2D) Neighbours(pos Point2D, options ...Plan2DNeighbourOption) (neighbours []Point2D) {
	cfg := &plan2DNeighboursConfig{limiter: func(n Point2D) bool { return !p.blocked[n] }}
	for _, option := range options {
		option(cfg)
	}
	if p.rangeY.Covers(pos.Y) {
		if cfg.includeExternal || pos.X > p.rangeX.Start {
			n := pos.AddXY(-1, 0)
			if cfg.limiter(n) {
				neighbours = append(neighbours, n)
			}
		}
		if cfg.includeExternal || pos.X < p.rangeX.End {
			n := pos.AddXY(1, 0)
			if cfg.limiter(n) {
				neighbours = append(neighbours, n)
			}
		}
	}
	if p.rangeX.Covers(pos.X) {
		if cfg.includeExternal || pos.Y > p.rangeY.Start {
			n := pos.AddXY(0, -1)
			if cfg.limiter(n) {
				neighbours = append(neighbours, n)
			}
		}
		if cfg.includeExternal || pos.Y < p.rangeY.End {
			n := pos.AddXY(0, 1)
			if cfg.limiter(n) {
				neighbours = append(neighbours, n)
			}
		}
	}
	return
}

// Print prints the plan on stdout.
func (p Plan2D) Print(options ...Plan2DPrintOption) {
	cfg := plan2DPrintConfig{
		print: func(pos Point2D) bool {
			if pos == p.Start {
				fmt.Print("S")
				return true
			}
			return false
		},
	}
	for _, option := range options {
		option(&cfg)
	}
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			pos := Point2D{X: x, Y: y}
			if p.blocked[pos] {
				fmt.Print("#")
			} else if !cfg.print(pos) {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// TileAt returns the value of the tile at the given position of the map.
func (p Plan2D) TileAt(pos Point2D) rune {
	return p.tiles[pos]
}

// WithEachPoint iterates over all the points on the map and calls the given function.
func (p Plan2D) WithEachPoint(callback func(Point2D)) {
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			callback(Point2D{X: x, Y: y})
		}
	}
}

func (p Plan2D) onMap(pos Point2D) bool {
	return pos.X >= 0 && pos.X < p.width && pos.Y >= 0 && pos.Y < p.height
}

// Plan2DNeighbourOption is a function to provide an option to Plan2D.Neighbours.
type Plan2DNeighbourOption func(*plan2DNeighboursConfig)

type plan2DNeighboursConfig struct {
	includeExternal bool
	limiter         func(Point2D) bool
}

// Plan2DOption is a function to provide an option to ParsePlan2D.
type Plan2DOption func(*plan2DConfig)

type plan2DConfig struct {
	handleUnknownTile func(Point2D, rune)
}

// Plan2DPrintOption is a function to provide an option to Plan2D.Print.
type Plan2DPrintOption func(*plan2DPrintConfig)

type plan2DPrintConfig struct {
	print func(Point2D) bool
}
