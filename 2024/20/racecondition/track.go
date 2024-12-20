package racecondition

import (
	"github.com/toaster/advent_of_code/internal/math"
)

// ParseRacetrack returns a new race Racetrack from the given input.
func ParseRacetrack(lines []string) Racetrack {
	m := Racetrack{
		heading: math.East,
	}
	m.plan = math.ParsePlan2D(lines, 'S', math.Plan2DWithUnknownTileHandler(func(pos math.Point2D, r rune) {
		if r == 'E' {
			m.end = pos
		}
	}))
	return m
}

// Racetrack describes a race condition festival racetrack (https://adventofcode.com/2024/day/20).
type Racetrack struct {
	end     math.Point2D
	heading math.Heading
	plan    math.Plan2D
}

// CountCheatsSavingAtLeast counts the amount of possible cheats which save at least the given amount
// of steps.
func (r Racetrack) CountCheatsSavingAtLeast(maxCheatLength, steps int) int {
	count := 0
	fairPath := r.computePath()
	for pos := range fairPath {
		for _, p := range pos.PointsWithinManhattanDistance(maxCheatLength) {
			if s, ok := fairPath[p]; ok && s-fairPath[pos] >= steps+pos.ManhattanDistance(p) {
				count++
			}
		}
	}
	return count
}

func (r Racetrack) computePath() map[math.Point2D]int {
	cur := r.plan.Start
	path := map[math.Point2D]int{cur: 0}
	for cur != r.end {
		next := r.plan.Neighbours(cur, math.Plan2DNeighboursWithLimiter(func(pos math.Point2D) bool {
			_, found := path[pos]
			return !r.plan.Blocked(pos) && !found
		}))[0]
		path[next] = path[cur] + 1
		cur = next
	}
	return path
}
