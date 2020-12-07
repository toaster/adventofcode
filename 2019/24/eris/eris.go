package eris

import (
	"fmt"
	"strings"
)

// Map is a map of Eris with bugs on it.
type Map struct {
	width  int
	height int
	m      map[point]*tile
}

// New creates a new Map of the input.
func New(input string) *Map {
	m := &Map{m: map[point]*tile{}}
	v := 1
	lines := strings.Split(strings.TrimSpace(input), "\n")
	m.height = len(lines)
	for y, line := range lines {
		m.width = len(line)
		for x, r := range line {
			p := point{x: x, y: y}
			m.m[p] = &tile{isBug: r == '#', value: v}
			v *= 2
		}
	}
	return m
}

// SimulateUntilRepeat simulates the bugs' life until the first pattern repeat.
func (m *Map) SimulateUntilRepeat() int {
	m.print()
	seen := map[int]bool{}
	b := m.computeBiodiversity()
	for !seen[b] {
		seen[b] = true
		bm := map[point]int{}
		for p := range m.m {
			bm[p] = m.countAdjacentBugs(p)
		}
		for p, t := range m.m {
			if t.isBug {
				if bm[p] != 1 {
					t.isBug = false
				}
			} else if bm[p] == 1 || bm[p] == 2 {
				t.isBug = true
			}
		}
		m.print()
		b = m.computeBiodiversity()
		// time.Sleep(2 * time.Second)
	}
	return b
}

func (m *Map) computeBiodiversity() int {
	b := 0
	for _, t := range m.m {
		if t.isBug {
			b += t.value
		}
	}
	return b
}

func (m *Map) countAdjacentBugs(p point) (c int) {
	if p.x > 0 && m.m[point{x: p.x - 1, y: p.y}].isBug {
		c++
	}
	if p.x < m.width-1 && m.m[point{x: p.x + 1, y: p.y}].isBug {
		c++
	}
	if p.y > 0 && m.m[point{x: p.x, y: p.y - 1}].isBug {
		c++
	}
	if p.y < m.height-1 && m.m[point{x: p.x, y: p.y + 1}].isBug {
		c++
	}
	return
}

func (m *Map) print() {
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.m[point{x, y}].isBug {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type point struct {
	x int
	y int
}

type tile struct {
	isBug bool
	value int
}
