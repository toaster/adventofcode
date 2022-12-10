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
	head := math.Point2D{}
	tail := math.Point2D{}
	visited := map[math.Point2D]bool{tail: true}
	var moves []move
	for _, line := range io.ReadLines() {
		components := strings.Split(line, " ")
		moves = append(moves, move{direction: components[0], length: io.ParseInt(components[1])})
	}

	for _, m := range moves {
		for i := 0; i < m.length; i++ {
			head, tail = step(m.direction, head, tail)
			visited[tail] = true
		}
	}
	fmt.Println(len(visited))
}

func step(direction string, head, tail math.Point2D) (math.Point2D, math.Point2D) {
	delta := math.Point2D{}
	additionalTailDelta := math.Point2D{}
	switch direction {
	case "U":
		delta.Y = -1
		additionalTailDelta.X = head.X - tail.X
	case "D":
		delta.Y = 1
		additionalTailDelta.X = head.X - tail.X
	case "R":
		delta.X = 1
		additionalTailDelta.Y = head.Y - tail.Y
	case "L":
		delta.X = -1
		additionalTailDelta.Y = head.Y - tail.Y
	}
	head = head.Add(delta)
	if !tail.IsAdjacent(head) {
		tail = tail.Add(delta)
		if tail.ManhattanDistance(head) > 1 {
			tail = tail.Add(additionalTailDelta)
		}
	}
	return head, tail
}
