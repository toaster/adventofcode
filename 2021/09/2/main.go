package main

import (
	"fmt"
	"sort"

	"github.com/toaster/advent_of_code/internal/io"
)

var pointCount = point{-1, -1}

func main() {
	rows := [][]int{{}}
	for _, line := range io.ReadLines() {
		row := []int{9}
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		row = append(row, 9)
		rows = append(rows, row)
	}
	height := len(rows) + 1
	width := len(rows[1])

	var borderRow []int
	for i := 0; i < width; i++ {
		borderRow = append(borderRow, 9)
	}
	rows[0] = borderRow
	rows = append(rows, borderRow)

	var basins []map[point]int
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if rows[y][x] < 9 && !isInAnyBasin(basins, x, y) {
				basin := map[point]int{}
				basins = append(basins, basin)
				fillBasin(basin, rows, x, y)
			}
		}
	}

	var sizes sort.IntSlice
	for _, basin := range basins {
		sizes = append(sizes, basin[pointCount])
	}
	sort.Sort(sort.Reverse(sizes))
	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}

func isInAnyBasin(basins []map[point]int, x, y int) bool {
	p := point{x, y}
	for _, basin := range basins {
		if basin[p] != 0 {
			return true
		}
	}
	return false
}

func fillBasin(basin map[point]int, rows [][]int, x, y int) {
	basin[point{x, y - 1}] = 10
	basin[point{x - 1, y}] = 10
	points := []point{{x, y}}
	for len(points) > 0 {
		count := len(points) - 1
		p := points[count]
		points = points[:count]
		if basin[p] != 0 {
			continue
		}

		level := rows[p.y][p.x]
		basin[p] = level + 1
		if level == 9 {
			continue
		}

		basin[pointCount]++
		points = considerPoint(basin, point{p.x, p.y - 1}, points) // up
		points = considerPoint(basin, point{p.x + 1, p.y}, points) // right
		points = considerPoint(basin, point{p.x, p.y + 1}, points) // down
		points = considerPoint(basin, point{p.x - 1, p.y}, points) // left
	}
}

func considerPoint(basin map[point]int, p point, points []point) []point {
	if basin[p] == 0 {
		points = append(points, p)
	}
	return points
}

type point struct {
	x int
	y int
}
