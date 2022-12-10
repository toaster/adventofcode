package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

type move struct {
	direction string
	length    int
}

func main() {
	var rope []math.Point2D
	for i := 0; i < 10; i++ {
		rope = append(rope, math.Point2D{})
	}
	visited := map[math.Point2D]bool{rope[len(rope)-1]: true}
	var moves []move
	for _, line := range io.ReadLines() {
		components := strings.Split(line, " ")
		moves = append(moves, move{direction: components[0], length: io.ParseInt(components[1])})
	}

	for _, m := range moves {
		for i := 0; i < m.length; i++ {
			step(m.direction, rope)
			visited[rope[len(rope)-1]] = true
		}
	}
	fmt.Println(len(visited))
}

func step(direction string, rope []math.Point2D) {
	delta := math.Point2D{}
	switch direction {
	case "U":
		delta.Y = -1
	case "D":
		delta.Y = 1
	case "R":
		delta.X = 1
	case "L":
		delta.X = -1
	}
	rope[0] = rope[0].Add(delta)
	for i := 1; i < len(rope); i++ {
		if rope[i].IsAdjacent(rope[i-1]) {
			break
		}

		if rope[i].ManhattanDistance(rope[i-1]) == 2 {
			// same row or column
			if rope[i].X == rope[i-1].X {
				// same column
				if rope[i].Y > rope[i-1].Y {
					rope[i].Y--
				} else {
					rope[i].Y++
				}
			} else {
				// same row
				if rope[i].X > rope[i-1].X {
					rope[i].X--
				} else {
					rope[i].X++
				}
			}
		} else {
			if rope[i].Y > rope[i-1].Y {
				rope[i].Y--
			} else {
				rope[i].Y++
			}
			if rope[i].X > rope[i-1].X {
				rope[i].X--
			} else {
				rope[i].X++
			}
		}
	}
}
