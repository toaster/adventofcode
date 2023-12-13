package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	connectingTo := map[math.Point2D][]math.Point2D{}
	connectedBy := map[math.Point2D][]math.Point2D{}
	startPos := math.Point2D{}
	for y, line := range io.ReadLines() {
		for x, c := range line {
			p := math.Point2D{X: x, Y: y}
			var a, b math.Point2D
			switch c {
			case '|':
				a = p.AddXY(0, 1)
				b = p.SubtractXY(0, 1)
			case '-':
				a = p.AddXY(1, 0)
				b = p.SubtractXY(1, 0)
			case 'L':
				a = p.AddXY(1, 0)
				b = p.SubtractXY(0, 1)
			case 'J':
				a = p.SubtractXY(1, 0)
				b = p.SubtractXY(0, 1)
			case '7':
				a = p.SubtractXY(1, 0)
				b = p.AddXY(0, 1)
			case 'F':
				a = p.AddXY(1, 0)
				b = p.AddXY(0, 1)
			case 'S':
				startPos = p
			}
			connectingTo[p] = append(connectingTo[p], a, b)
			connectedBy[a] = append(connectedBy[a], p)
			connectedBy[b] = append(connectedBy[b], p)
		}
	}
	l := 1
	cur := startPos
	for next := connectedBy[startPos][0]; next != startPos; {
		var nextNext math.Point2D
		if connectingTo[next][0] == cur {
			nextNext = connectingTo[next][1]
		} else {
			nextNext = connectingTo[next][0]
		}
		cur = next
		next = nextNext
		l++
	}
	fmt.Println(l, l/2)
}
