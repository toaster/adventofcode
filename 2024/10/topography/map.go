package topography

import (
	"github.com/toaster/advent_of_code/internal/math"
)

// ParseMap creates a new Map from the input lines.
func ParseMap(lines []string) Map {
	m := Map{
		height: len(lines),
		tiles:  map[math.Point2D]uint8{},
	}
	for y, line := range lines {
		for x, c := range line {
			if m.width == 0 {
				m.width = len(line)
			}
			v := uint8(c - '0')
			pos := math.Point2D{X: x, Y: y}
			m.tiles[pos] = v
			if v == 0 {
				m.trailHeads = append(m.trailHeads, pos)
			}
			if v == 9 {
				m.summits = append(m.summits, pos)
			}
		}
	}
	m.rangeX = math.Range{
		Start: 0,
		End:   m.width - 1,
	}
	m.rangeY = math.Range{
		Start: 0,
		End:   m.height - 1,
	}
	return m
}

// Map is a topographic map.
type Map struct {
	height     int
	rangeX     math.Range
	rangeY     math.Range
	summits    []math.Point2D
	tiles      map[math.Point2D]uint8
	trailHeads []math.Point2D
	width      int
}

// SumUpRatings sums up all ratings of all trailheads on the Map.
func (m Map) SumUpRatings() int {
	ratings := map[math.Point2D]int{}
	rating := 0
	for _, head := range m.trailHeads {
		rating += m.computeRating(head, ratings)
	}
	return rating
}

// SumUpTrailheadScores sums up all scores of all trailheads on the Map.
func (m Map) SumUpTrailheadScores() int {
	reachableSummits := map[math.Point2D]map[math.Point2D]bool{}
	score := 0
	for _, head := range m.trailHeads {
		m.trackTrail(head, reachableSummits)
		score += len(reachableSummits[head])
	}
	return score
}

func (m Map) computeRating(pos math.Point2D, ratings map[math.Point2D]int) int {
	if ratings[pos] == -1 {
		return 0
	}
	if ratings[pos] > 0 {
		return ratings[pos]
	}

	if m.tiles[pos] == 9 {
		ratings[pos] = 1
		return 1
	}

	rating := 0
	for _, n := range m.nextNeighbours(pos) {
		rating += m.computeRating(n, ratings)
	}
	ratings[pos] = rating
	if ratings[pos] == 0 {
		ratings[pos] = -1
	}
	return rating
}

func (m Map) trackTrail(pos math.Point2D, reachableSummits map[math.Point2D]map[math.Point2D]bool) {
	if reachableSummits[pos] != nil {
		return
	}

	reachableSummits[pos] = map[math.Point2D]bool{}
	if m.tiles[pos] == 9 {
		reachableSummits[pos][pos] = true
		return
	}

	for _, n := range m.nextNeighbours(pos) {
		m.trackTrail(n, reachableSummits)
		for s := range reachableSummits[n] {
			reachableSummits[pos][s] = true
		}
	}
	return
}

func (m Map) nextNeighbours(pos math.Point2D) (neighbours []math.Point2D) {
	l := m.tiles[pos]
	if m.rangeY.Covers(pos.Y) {
		if pos.X > m.rangeX.Start {
			n := pos.AddXY(-1, 0)
			if m.tiles[n] == l+1 {
				neighbours = append(neighbours, n)
			}
		}
		if pos.X < m.rangeX.End {
			n := pos.AddXY(1, 0)
			if m.tiles[n] == l+1 {
				neighbours = append(neighbours, n)
			}
		}
	}
	if m.rangeX.Covers(pos.X) {
		if pos.Y > m.rangeY.Start {
			n := pos.AddXY(0, -1)
			if m.tiles[n] == l+1 {
				neighbours = append(neighbours, n)
			}
		}
		if pos.Y < m.rangeY.End {
			n := pos.AddXY(0, 1)
			if m.tiles[n] == l+1 {
				neighbours = append(neighbours, n)
			}
		}
	}
	return
}
