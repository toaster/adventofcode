package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

const (
	up direction = iota
	right
	down
	left
	directionCount
	turnRight = 3
	turnLeft  = 1
)

func main() {
	cur := math.Point2D{}
	var edges []math.LineSegment2D
	var contour *math.LinkedLineSegment2D
	var contourStart *math.LinkedLineSegment2D
	side := 0
	var lastDirection byte
	for _, line := range io.ReadLines() {
		components := strings.Split(line, " ")
		length := io.ParseInt("0x" + components[2][2:7])
		delta := math.Point2D{}
		d := components[2][7]
		switch d {
		case '3': // up
			delta.Y = -length
			if lastDirection == '0' {
				side--
			} else {
				side++
			}
		case '1': // down
			delta.Y = length
			if lastDirection == '0' {
				side++
			} else {
				side--
			}
		case '0': // right
			delta.X = length
			if lastDirection == '3' {
				side++
			} else {
				side--
			}
		case '2': // left
			delta.X = -length
			if lastDirection == '3' {
				side--
			} else {
				side++
			}
		}
		if lastDirection == 0 {
			side = 0
		}
		lastDirection = d
		edge := math.LineSegment2D{A: cur}
		cur = cur.Add(delta)
		edge.B = cur
		edges = append(edges, edge)
		contour = &math.LinkedLineSegment2D{
			LineSegment2D: edge,
			Previous:      contour,
		}
		if contourStart == nil {
			contourStart = contour
		} else {
			contour.Previous.Next = contour
		}
	}
	// This could be part of the loop, including sanity checks.
	// However, we know the data creates a valid contour.
	contour.Next = contourStart
	contourStart.Previous = contour

	for c := contourStart; c.Next != contourStart; c = c.Next {
		fmt.Printf("=> %#v\n", c.LineSegment2D)
	}

	volume := 0
	for contourStart != nil {
		var area int
		area, contourStart = reduce(contourStart, side)
		volume += area
	}
	fmt.Println(volume)
}

type direction int

func reduce(start *math.LinkedLineSegment2D, side int) (int, *math.LinkedLineSegment2D) {
	var expectedTurn direction
	if side > 0 {
		expectedTurn = turnRight
	} else {
		expectedTurn = turnLeft
	}
	fmt.Printf("start: %#v; expect turn: %d\n", start.LineSegment2D, expectedTurn)
	cur := start
	if cur.Next.Next == cur.Previous.Previous {
		// rectangle
		return (cur.A.ManhattanDistance(cur.B) + 1) * (cur.Previous.A.ManhattanDistance(cur.Previous.B) + 1), nil
	}

Next:
	for {
		prevDir := computeDirection(cur.Previous.LineSegment2D)
		dir := computeDirection(cur.LineSegment2D)
		turn := (prevDir + directionCount - dir) % directionCount
		if turn != expectedTurn {
			cur = cur.Next
			fmt.Printf("A turn (%d) wrong => next: %#v\n", turn, cur.LineSegment2D)
			if cur == start {
				panic("ooops, no rectangle found")
			}
			continue Next
		}

		nextDir := computeDirection(cur.Next.LineSegment2D)
		nextTurn := (dir + directionCount - nextDir) % directionCount
		if nextTurn != expectedTurn {
			cur = cur.Next
			fmt.Printf("B turn (%d) wrong => next: %#v\n", nextTurn, cur.LineSegment2D)
			if cur == start {
				panic("ooops, no rectangle found")
			}
			continue Next
		}

		// found edge with turns in the right direction
		lenPrev := cur.Previous.B.ManhattanDistance(cur.Previous.A)
		lenNext := cur.Next.B.ManhattanDistance(cur.Next.A)
		cutLen := min(lenPrev, lenNext)

		{
			// check that new edge does not cross existing one
			newEdge := math.LineSegment2D{A: cur.A, B: cur.B}

			switch prevDir {
			case up:
				newEdge.A.Y += cutLen
				newEdge.B.Y += cutLen
			case right:
				newEdge.A.X -= cutLen
				newEdge.B.X -= cutLen
			case down:
				newEdge.A.Y -= cutLen
				newEdge.B.Y -= cutLen
			case left:
				newEdge.A.X += cutLen
				newEdge.B.X += cutLen
			}

			for c := cur; c.Next != cur; c = c.Next {
				if crosses(newEdge, c.LineSegment2D) {
					cur = cur.Next
					fmt.Printf("New edge crosses %#v => next: %#v\n", c.LineSegment2D, cur.LineSegment2D)
					if cur == start {
						panic("ooops, no rectangle found")
					}

					continue Next
				}
			}
		}

		area := (cur.B.ManhattanDistance(cur.A) + 1) * cutLen

		if lenPrev == cutLen {
			if computeDirection(cur.Previous.Previous.LineSegment2D) == dir {
				// cut prev and join prev.prev into cur
				cur.A = cur.Previous.Previous.A
				cur.Previous.Previous.Previous.Next = cur
				cur.Previous = cur.Previous.Previous.Previous
			} else {
				// cut prev
				cur.A = cur.Previous.A
				cur.Previous.Previous.Next = cur
				cur.Previous = cur.Previous.Previous
			}
		} else {
			switch prevDir {
			case up:
				cur.Previous.B.Y += cutLen
			case right:
				cur.Previous.B.X -= cutLen
			case down:
				cur.Previous.B.Y -= cutLen
			case left:
				cur.Previous.B.X += cutLen
			}
			cur.A = cur.Previous.B
		}
		if lenNext == cutLen {
			if computeDirection(cur.Next.Next.LineSegment2D) == dir {
				// cut next and join next.next into cur
				cur.B = cur.Next.Next.B
				cur.Next.Next.Next.Previous = cur
				cur.Next = cur.Next.Next.Next
			} else {
				// cut next
				cur.B = cur.Next.B
				cur.Next.Next.Previous = cur
				cur.Next = cur.Next.Next
			}
		} else {
			switch nextDir {
			case up:
				cur.Next.A.Y -= cutLen
			case right:
				cur.Next.A.X += cutLen
			case down:
				cur.Next.A.Y += cutLen
			case left:
				cur.Next.A.X -= cutLen
			}
			cur.B = cur.Next.A
		}

		return area, cur
	}
}

func computeDirection(d math.LineSegment2D) direction {
	if d.A.X > d.B.X {
		return left
	}
	if d.A.X < d.B.X {
		return right
	}
	if d.A.Y > d.B.Y {
		return up
	}
	return down
}

// Crosses returns whether two line segments cross
func crosses(a, b math.LineSegment2D) bool {
	if max(a.A.X, a.B.X) < min(b.A.X, b.B.X) {
		return false
	}

	if min(a.A.X, a.B.X) > max(b.A.X, b.B.X) {
		return false
	}

	if max(a.A.Y, a.B.Y) < min(b.A.Y, b.B.Y) {
		return false
	}

	if min(a.A.Y, a.B.Y) > max(b.A.Y, b.B.Y) {
		return false
	}

	if a.A.X == a.B.X {
		if b.A.X == b.B.X {
			return b.A.X == a.A.X && max(min(a.A.Y, a.B.Y), min(b.A.Y, b.B.Y)) < min(max(a.A.Y, a.B.Y), max(b.A.Y, b.B.Y))
		}

		return b.A.Y > min(a.A.Y, a.B.Y) && b.A.Y < max(a.A.Y, a.B.Y) && a.A.X > min(b.A.X, b.B.X) && a.A.X < max(b.A.X, b.B.X)
	}

	if b.A.Y == b.B.Y {
		return b.A.Y == a.A.Y && max(min(a.A.X, a.B.X), min(b.A.X, b.B.X)) < min(max(a.A.X, a.B.X), max(b.A.X, b.B.X))
	}

	return b.A.X > min(a.A.X, a.B.X) && b.A.X < max(a.A.X, a.B.X) && a.A.Y > min(b.A.Y, b.B.Y) && a.A.Y < max(b.A.Y, b.B.Y)
}
