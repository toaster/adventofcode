package easterbunny

import (
	"github.com/toaster/advent_of_code/internal/math"
)

// ParseAntennaMap creates a new AntennaMap from the input.
func ParseAntennaMap(lines []string) AntennaMap {
	m := AntennaMap{positions: map[rune][]math.Point2D{}}
	m.height = len(lines)
	for y, line := range lines {
		if m.width == 0 {
			m.width = len(line)
		}
		for x, c := range line {
			if c == '.' {
				continue
			}
			m.positions[c] = append(m.positions[c], math.Point2D{X: x, Y: y})
		}
	}
	return m
}

// AntennaMap describes a map containing Easter Bunny antennas (https://adventofcode.com/2024/day/8).
type AntennaMap struct {
	height    int
	positions map[rune][]math.Point2D
	width     int
}

// CountAntinodes counts the antinodes of all antennas on the map.
func (m AntennaMap) CountAntinodes(considerResonantHarmonics bool) int {
	antinodes := map[math.Point2D]bool{}
	for _, positions := range m.positions {
		for i, pos := range positions {
			for j := i + 1; j < len(positions); j++ {
				otherPos := positions[j]
				difference := pos.Subtract(otherPos)
				if considerResonantHarmonics {
					antinodes[pos] = true
					antinodes[otherPos] = true
				}
				for candidate := pos.Add(difference); m.onMap(candidate); candidate = candidate.Add(difference) {
					antinodes[candidate] = true
					if !considerResonantHarmonics {
						break
					}
				}
				for candidate := otherPos.Subtract(difference); m.onMap(candidate); candidate = candidate.Subtract(difference) {
					antinodes[candidate] = true
					if !considerResonantHarmonics {
						break
					}
				}
			}
		}
	}
	return len(antinodes)
}

func (m AntennaMap) onMap(pos math.Point2D) bool {
	return pos.X >= 0 && pos.X < m.width && pos.Y >= 0 && pos.Y < m.height
}
