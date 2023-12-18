package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	cur := math.Point2D{}
	dug := map[math.Point2D]bool{cur: true}
	topLeft := math.Point2D{}
	bottomRight := math.Point2D{}
	for _, line := range io.ReadLines() {
		components := strings.Split(line, " ")
		direction := components[0]
		length := io.ParseInt(components[1])
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
		for i := 0; i < length; i++ {
			cur = cur.Add(delta)
			dug[cur] = true
		}
		if cur.X < topLeft.X {
			topLeft.X = cur.X
		} else if cur.X > bottomRight.X {
			bottomRight.X = cur.X
		}
		if cur.Y < topLeft.Y {
			topLeft.Y = cur.Y
		} else if cur.Y > bottomRight.Y {
			bottomRight.Y = cur.Y
		}
	}
	fmt.Println(len(dug))
	capacity := 0
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		inside := false
		trench := false
		trenchFromUp := false
		trenchFromDown := false
		for x := topLeft.X; x <= bottomRight.X+1; x++ {
			c := math.Point2D{X: x, Y: y}
			if dug[c] {
				trench = true
				if y > topLeft.Y && dug[c.SubtractXY(0, 1)] {
					trenchFromUp = !trenchFromUp
				}
				if y < bottomRight.Y && dug[c.AddXY(0, 1)] {
					trenchFromDown = !trenchFromDown
				}
				capacity++
				fmt.Print("#")
			} else {
				if trench {
					if trenchFromUp && trenchFromDown {
						inside = !inside
					}
					trench = false
					trenchFromUp = false
					trenchFromDown = false
				}
				if inside {
					fmt.Print("^")
				} else {
					fmt.Print(".")
				}
				if inside {
					capacity++
				}
			}
		}
		fmt.Println()
	}
	fmt.Println(capacity)
}
