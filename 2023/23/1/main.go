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
	area := map[math.Point2D]tile{}
	for y, line := range io.ReadLines() {
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
	dist, _ := math.FindLongestDistance(func(p math.Point2D) []math.Point2D {
		switch area[p] {
		case slopeRight:
			return []math.Point2D{p.AddXY(1, 0)}
		case slopeLeft:
			return []math.Point2D{p.SubtractXY(1, 0)}
		case slopeUp:
			return []math.Point2D{p.SubtractXY(0, 1)}
		case slopeDown:
			return []math.Point2D{p.AddXY(0, 1)}
		}
		var adjacents []math.Point2D
		for _, n := range p.Neighbours(wholeMap) {
			if area[n] == wall {
				continue
			}
			if area[n] == slopeRight && n.X == p.X-1 {
				continue
			}
			if area[n] == slopeLeft && n.X == p.X+1 {
				continue
			}
			if area[n] == slopeDown && n.Y == p.Y-1 {
				continue
			}
			if area[n] == slopeUp && n.Y == p.Y+1 {
				continue
			}
			adjacents = append(adjacents, n)
		}
		return adjacents
	}, *start, *dest)
	fmt.Println(dist)
}

type tile rune
