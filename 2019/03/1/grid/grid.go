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

func Mark(grid map[Pos]bool, rawDirections string) {
	directions := strings.Split(rawDirections, ",")
	p := Pos{0, 0}
	for _, direction := range directions {
		distance, _ := strconv.Atoi(direction[1:])
		switch direction[0] {
		case 'U':
			for dy := 0; dy < distance; dy++ {
				p.y++
				grid[p] = true
			}
		case 'D':
			for dy := 0; dy < distance; dy++ {
				p.y--
				grid[p] = true
			}
		case 'R':
			for dy := 0; dy < distance; dy++ {
				p.x++
				grid[p] = true
			}
		case 'L':
			for dy := 0; dy < distance; dy++ {
				p.x--
				grid[p] = true
			}
		}
	}
}

func Search(grid map[Pos]bool, rawDirections string) int {
	directions := strings.Split(rawDirections, ",")
	p := Pos{0, 0}
	d := math.MaxInt64
	for _, direction := range directions {
		distance, _ := strconv.Atoi(direction[1:])
		switch direction[0] {
		case 'U':
			for dy := 0; dy < distance; dy++ {
				p.y++
				d = newDistance(p, grid, d)
			}
		case 'D':
			for dy := 0; dy < distance; dy++ {
				p.y--
				d = newDistance(p, grid, d)
			}
		case 'R':
			for dy := 0; dy < distance; dy++ {
				p.x++
				d = newDistance(p, grid, d)
			}
		case 'L':
			for dy := 0; dy < distance; dy++ {
				p.x--
				d = newDistance(p, grid, d)
			}
		}
	}
	return d
}

func newDistance(p Pos, grid map[Pos]bool, distance int) int {
	if grid[p] {
		d := abs(p.x) + abs(p.y)
		if d < distance {
			return d
		}
	}
	return distance
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
