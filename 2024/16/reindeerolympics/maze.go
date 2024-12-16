package reindeerolympics

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/math"
	"github.com/toaster/advent_of_code/internal/util"
)

// ParseMaze returns a new Maze from the given input.
func ParseMaze(lines []string, costsPerStep, costsPerTurn int) Maze {
	m := Maze{
		costsPerStep: costsPerStep,
		costsPerTurn: costsPerTurn,
		heading:      math.East,
	}
	m.plan = math.ParsePlan2D(lines, 'S', math.Plan2DWithUnknownTileHandler(func(pos math.Point2D, r rune) {
		if r == 'E' {
			m.end = pos
		}
	}))
	return m
}

// Maze describes a Reindeer Olympics maze (https://adventofcode.com/2024/day/16).
type Maze struct {
	costsPerStep int
	costsPerTurn int
	end          math.Point2D
	heading      math.Heading
	plan         math.Plan2D
}

// CountBestPathTiles counts the tiles of all best paths, i.e., paths with the cheapest costs.
func (m Maze) CountBestPathTiles() int {
	p := m.computeCheapestPath()
	seats := map[math.Point2D]bool{}
	for pos := range p.stepMap {
		seats[pos] = true
	}
	for _, prefix := range p.prefixes {
		for pos := range prefix.stepMap {
			seats[pos] = true
		}
	}
	m.printSeats(seats)
	return len(seats)
}

// FindCheapestPath finds the cheapest path through the maze.
func (m Maze) FindCheapestPath() int {
	p := m.computeCheapestPath()
	m.printPath(p)
	return p.costs
}

func (m Maze) computeCheapestPath() *path {
	startEntry := entry{heading: m.heading, pos: m.plan.Start}
	paths := []*path{{heading: m.heading, pos: m.plan.Start, stepMap: map[math.Point2D]math.Heading{m.plan.Start: m.heading}}}
	entered := map[entry]entryInfo{startEntry: {p: paths[0]}}
	for {
		if paths[0].pos == m.end {
			break
		}

		var nextSteps []step
		for _, n := range []struct {
			heading math.Heading
			costs   int
		}{
			{heading: paths[0].heading.TurnRight(), costs: 1001},
			{heading: paths[0].heading.TurnLeft(), costs: 1001},
			{heading: paths[0].heading, costs: 1},
		} {
			if next := n.heading.Facing(paths[0].pos); !m.plan.Blocked(next) {
				nextSteps = append(nextSteps, step{entry: entry{heading: n.heading, pos: next}, costs: n.costs})
			}
		}
		if len(nextSteps) == 0 {
			// dead end
			paths = paths[1:]
			continue
		}

		for i := 1; i < len(nextSteps); i++ {
			p := *paths[0]
			p.prefixes = util.CopySlice(p.prefixes)
			p.stepMap = util.CopyMap(p.stepMap)
			paths = append([]*path{&p}, paths...)
		}

		for _, target := range nextSteps {
			paths[0].costs += target.costs
			if info, ok := entered[target.entry]; ok {
				if paths[0].costs >= info.costs {
					if paths[0].costs == info.costs {
						info.p.prefixes = append(info.p.prefixes, paths[0])
						info.p.prefixes = append(info.p.prefixes, paths[0].prefixes...)
					}
					// circuit breaker (already entered this path on a cheaper route)
					paths = paths[1:]
					continue
				}
			}

			paths[0].pos = target.pos
			paths[0].heading = target.heading
			paths[0].stepMap[target.pos] = target.heading
			entered[target.entry] = entryInfo{p: paths[0], costs: paths[0].costs}
			for i := 1; i < len(paths); i++ {
				if paths[i-1].costs < paths[i].costs {
					break
				}
				paths[i-1], paths[i] = paths[i], paths[i-1]
			}
		}
	}
	return paths[0]
}

func (m Maze) printPath(p *path) {
	m.plan.Print(math.Plan2DPrintAt(func(pos math.Point2D) bool {
		if h, ok := p.stepMap[pos]; ok {
			fmt.Print("\x1b[38;5;15m")
			fmt.Print(h.String())
			fmt.Print("\x1b[0m")
			return true
		}
		return false
	}))
}

func (m Maze) printSeats(seats map[math.Point2D]bool) {
	m.plan.Print(math.Plan2DPrintAt(func(pos math.Point2D) bool {
		if seats[pos] {
			fmt.Print("\x1b[38;5;15m")
			fmt.Print("O")
			fmt.Print("\x1b[0m")
			return true
		}
		return false
	}))
}

type entry struct {
	heading math.Heading
	pos     math.Point2D
}

type entryInfo struct {
	p     *path
	costs int
}

type path struct {
	costs    int
	heading  math.Heading
	pos      math.Point2D
	prefixes []*path
	stepMap  map[math.Point2D]math.Heading
}

type step struct {
	entry
	costs int
}
