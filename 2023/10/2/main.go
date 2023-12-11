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
	width := 0
	height := 0
	for y, line := range io.ReadLines() {
		width = len(line)
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
		height++
	}
	cur := startPos
	loop := map[math.Point2D]bool{startPos: true}
	connectingTo[startPos] = connectedBy[startPos]
	for next := connectedBy[startPos][0]; next != startPos; {
		loop[next] = true
		var nextNext math.Point2D
		if connectingTo[next][0] == cur {
			nextNext = connectingTo[next][1]
		} else {
			nextNext = connectingTo[next][0]
		}
		cur = next
		next = nextNext
	}
	insideLoop := map[math.Point2D]bool{}
	for y := 1; y < height-1; y++ {
		isInside := false
		verticalConnected := y
		for x := 0; x < width; x++ {
			p := math.Point2D{X: x, Y: y}
			if loop[p] {
				// fmt.Println(p, "is in loop, vert:", verticalConnected)
				a := connectingTo[p][0]
				b := connectingTo[p][1]
				if a.Y != p.Y && b.Y != p.Y { // '|'
					// fmt.Println("    toggles inside", a, b)
					isInside = !isInside
				} else {
					var v math.Point2D
					if a.Y != p.Y {
						v = a
					} else if b.Y != p.Y {
						v = b
					}
					if v != (math.Point2D{}) { // !'-'
						if verticalConnected != y {
							// fmt.Println("    finishes horizontal to", v.Y)
							if v.Y != verticalConnected {
								isInside = !isInside
								// fmt.Println("    and toggles inside")
							}
							verticalConnected = y
						} else {
							verticalConnected = v.Y
							// fmt.Println("    goes west and connects vert to", verticalConnected)
						}
					}
				}
			} else if isInside {
				insideLoop[p] = true
			}
		}
	}
	fmt.Println(len(insideLoop))
}
