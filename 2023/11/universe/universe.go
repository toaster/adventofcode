package universe

import (
	"fmt"
	"slices"

	"github.com/toaster/advent_of_code/internal/math"
)

// Parse parses the observed universe.
func Parse(input []string) *Universe {
	galaxies := map[math.Point2D]bool{}
	width := 0
	height := 0
	for y, line := range input {
		width = len(line)
		for x, c := range line {
			if c == '#' {
				galaxies[math.Point2D{X: x, Y: y}] = true
			}
		}
		height++
	}
	return &Universe{
		galaxies: galaxies,
		height:   height,
		width:    width,
	}
}

// Universe represents the universe :).
type Universe struct {
	galaxies map[math.Point2D]bool
	height   int
	width    int
}

// ComputeShortestPathBetweenAllGalaxies sums up the shortest paths between all pairs of galaxies.
func (u *Universe) ComputeShortestPathBetweenAllGalaxies() int {
	l := 0
	for p1 := range u.galaxies {
		for p2 := range u.galaxies {
			l += p1.ManhattanDistance(p2)
		}
	}
	return l / 2
}

// Expand expands the universe by the given factor.
func (u *Universe) Expand(factor int) {
	additionalCount := factor - 1

	var rowsWithGalaxies []int
	for p := range u.galaxies {
		if !math.Contains(rowsWithGalaxies, p.Y) {
			rowsWithGalaxies = append(rowsWithGalaxies, p.Y)
		}
	}
	slices.Sort(rowsWithGalaxies)

	var columnsWithGalaxies []int
	for p := range u.galaxies {
		if !math.Contains(columnsWithGalaxies, p.X) {
			columnsWithGalaxies = append(columnsWithGalaxies, p.X)
		}
	}
	slices.Sort(columnsWithGalaxies)

	for y := 0; y < u.height; y++ {
		empty := true
		for x := 0; x < u.width; x++ {
			p := math.Point2D{X: x, Y: y}
			if u.galaxies[p] {
				empty = false
				break
			}
		}
		if empty {
			for i := len(rowsWithGalaxies) - 1; ; i-- {
				z := rowsWithGalaxies[i]
				if z < y {
					break
				}

				for _, x := range columnsWithGalaxies {
					oldP := math.Point2D{X: x, Y: z}
					if u.galaxies[oldP] {
						u.galaxies[oldP.AddXY(0, additionalCount)] = true
						delete(u.galaxies, oldP)
					}
				}
				rowsWithGalaxies[i] = z + additionalCount
			}
			y += additionalCount
			u.height += additionalCount
		}
	}

	for x := 0; x < u.width; x++ {
		empty := true
		for y := 0; y < u.height; y++ {
			p := math.Point2D{X: x, Y: y}
			if u.galaxies[p] {
				empty = false
				break
			}
		}
		if empty {
			for i := len(columnsWithGalaxies) - 1; ; i-- {
				z := columnsWithGalaxies[i]
				if z < x {
					break
				}

				for _, y := range rowsWithGalaxies {
					oldP := math.Point2D{X: z, Y: y}
					if u.galaxies[oldP] {
						u.galaxies[oldP.AddXY(additionalCount, 0)] = true
						delete(u.galaxies, oldP)
					}
				}
				columnsWithGalaxies[i] = z + additionalCount
			}
			x += additionalCount
			u.width += additionalCount
		}
	}
}

// Print prints the universe to stdout.
func (u *Universe) Print() {
	fmt.Println(u.width, "x", u.height)
	for y := 0; y < u.height; y++ {
		for x := 0; x < u.width; x++ {
			if u.galaxies[(math.Point2D{X: x, Y: y})] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
