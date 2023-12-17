package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

const maxStraight = 3

const (
	north direction = iota
	east
	south
	west
	directionCount
	oppositeDiff direction = 2
)

func main() {
	city := map[math.Point2D]int{}
	width := 0
	height := 0
	for y, line := range io.ReadLines() {
		if width == 0 {
			width = len(line)
		}
		height++
		for x, v := range line {
			city[math.Point2D{X: x, Y: y}] = int(v - '0')
		}
	}
	start := point{heading: east}
	end := point{pos: math.Point2D{X: width - 1, Y: height - 1}}
	distance := math.FindShortestWeightedDistance(
		func(p point) []point {
			var adjacents []point
			for d := north; d < directionCount; d++ {
				opposite := (d + oppositeDiff) % directionCount
				if p.heading != opposite && (p.heading != d || p.straight < maxStraight) {
					n := point{pos: p.pos, heading: d, straight: 1}
					switch d {
					case north:
						n.pos.Y--
						if n.pos.Y < 0 {
							continue
						}
					case east:
						n.pos.X++
						if n.pos.X == width {
							continue
						}
					case south:
						n.pos.Y++
						if n.pos.Y == height {
							continue
						}
					case west:
						n.pos.X--
						if n.pos.X < 0 {
							continue
						}
					}
					if p.heading == d {
						n.straight = p.straight + 1
					}
					adjacents = append(adjacents, n)
				}
			}
			return adjacents
		},
		func(_, b point) int {
			return city[b.pos]
		},
		func(p point) bool {
			return p.pos == end.pos
		},
		start,
	)
	fmt.Println(distance)
}

type direction int

type point struct {
	pos      math.Point2D
	heading  direction
	straight int
}
