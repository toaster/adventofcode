package main

import (
	"fmt"
	math2 "math"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

const (
	floor      tile = '.'
	slopeRight tile = '>'
	slopeLeft  tile = '<'
	slopeUp    tile = '^'
	slopeDown  tile = 'v'
	wall       tile = '#'
)

func main() {
	var start, dest *math.Point2D
	width := 0
	height := 0
	area := map[math.Point2D]tile{}
	for y, line := range io.ReadLines() {
		if width == 0 {
			width = len(line)
		}
		height++
		for x, c := range line {
			p := math.Point2D{X: x, Y: y}
			area[p] = tile(c)
			if area[p] != wall {
				if start == nil {
					start = &p
				}
				dest = &p
			}
		}
	}
	wholeMap := math.Rectangle2D{TopLeft: math.Point2D{}, BottomRight: math.Point2D{X: math2.MaxInt, Y: math2.MaxInt}}
	dist, path := math.FindLongestDistance(func(p math.Point2D) []math.Point2D {
		var adjacents []math.Point2D
		for _, n := range p.Neighbours(wholeMap) {
			if area[n] == wall {
				continue
			}
			adjacents = append(adjacents, n)
		}
		return adjacents
	}, *start, *dest)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := math.Point2D{X: x, Y: y}
			if path[p] {
				fmt.Print("O")
			} else {
				fmt.Print(string(area[p]))
			}
		}
		fmt.Println()
	}
	fmt.Println(dist)
}

type tile rune
