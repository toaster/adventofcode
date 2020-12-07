package grid

import (
	"math"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

func Mark(grid map[Pos]int, rawDirections string) {
	directions := strings.Split(rawDirections, ",")
	p := Pos{0, 0}
	l := 0
	for _, direction := range directions {
		distance, _ := strconv.Atoi(direction[1:])
		for d := 0; d < distance; d++ {
			l++
			switch direction[0] {
			case 'U':
				p.y++
			case 'D':
				p.y--
			case 'R':
				p.x++
			case 'L':
				p.x--
			}
			grid[p] = l
		}
	}
}

func Search(grid map[Pos]int, rawDirections string) int {
	directions := strings.Split(rawDirections, ",")
	p := Pos{0, 0}
	l := 0
	dist := math.MaxInt64
	for _, direction := range directions {
		distance, _ := strconv.Atoi(direction[1:])
		for d := 0; d < distance; d++ {
			l++
			switch direction[0] {
			case 'U':
				p.y++
			case 'D':
				p.y--
			case 'R':
				p.x++
			case 'L':
				p.x--
			}
			dist = newDistance(p, grid, l, dist)
		}
	}
	return dist
}

func newDistance(p Pos, grid map[Pos]int, length, distance int) int {
	if grid[p] > 0 {
		d := length + grid[p]
		if d < distance {
			return d
		}
	}
	return distance
}
