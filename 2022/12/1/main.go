package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	var area [][]int
	start := math.Point2D{}
	end := math.Point2D{}
	for y, line := range io.ReadLines() {
		area = append(area, nil)
		for x, char := range line {
			switch char {
			case 'S':
				start.X = x
				start.Y = y
				area[y] = append(area[y], 0)
			case 'E':
				end.X = x
				end.Y = y
				area[y] = append(area[y], int('z'-'a'))
			default:
				area[y] = append(area[y], int(char-'a'))
			}
		}
	}

	distances := map[math.Point2D]int{start: 1}
	xRange := math.Range{End: len(area[0]) - 1}
	yRange := math.Range{End: len(area) - 1}
	candidates := []math.Point2D{start}
	for len(candidates) > 0 {
		candidates = trace(candidates, xRange, yRange, distances, area)
	}
	// printDistances(distances, xRange, yRange)
	fmt.Println(distances[end] - 1)
}

func printDistances(distances map[math.Point2D]int, xRange, yRange math.Range) {
	for y := yRange.Start; y <= yRange.End; y++ {
		for x := xRange.Start; x <= xRange.End; x++ {
			fmt.Printf("%3d ", distances[math.Point2D{X: x, Y: y}])
		}
		fmt.Println()
	}
}

func trace(candidates []math.Point2D, xRange math.Range, yRange math.Range, distances map[math.Point2D]int, area [][]int) (nextCandidates []math.Point2D) {
	for _, candidate := range candidates {
		for _, p := range candidate.Neighbours(xRange, yRange) {
			if distances[p] != 0 {
				continue
			}
			if area[p.Y][p.X] > area[candidate.Y][candidate.X]+1 {
				continue
			}

			distances[p] = distances[candidate] + 1
			nextCandidates = append(nextCandidates, p)
		}
	}
	return
}
