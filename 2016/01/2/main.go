package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	directions := strings.Split(strings.TrimSpace(io.ReadAll()), ", ")
	start := math.Point2D{}
	loc := start
	heading := 0
	m := map[math.Point2D]bool{loc: true}
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	for _, direction := range directions {
		fmt.Println(direction)
		distance := io.ParseInt(direction[1:])
		switch direction[0] {
		case 'R':
			heading++
		case 'L':
			heading += 3
		}
		heading %= 4
		for i := 0; i < distance; i++ {
			switch heading {
			case 0:
				loc = loc.SubtractXY(0, 1)
				if loc.Y < minY {
					minY--
				}
			case 1:
				loc = loc.AddXY(1, 0)
				if loc.X > maxX {
					maxX++
				}
			case 2:
				loc = loc.AddXY(0, 1)
				if loc.Y > maxY {
					maxY++
				}
			case 3:
				loc = loc.SubtractXY(1, 0)
				if loc.X < minX {
					minX--
				}
			}
			if m[loc] {
				fmt.Println(start.ManhattanDistance(loc))
				return
			}
			m[loc] = true
		}
		printMap(minX, maxX, minY, maxY, m, loc, heading)
	}
}

func printMap(minX, maxX, minY, maxY int, m map[math.Point2D]bool, loc math.Point2D, heading int) {
	// fmt.Println(loc)
	start := math.Point2D{X: 0, Y: 0}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := math.Point2D{X: x, Y: y}
			if p == start {
				fmt.Print("S")
			} else if p == loc {
				switch heading {
				case 0:
					fmt.Print("^")
				case 1:
					fmt.Print(">")
				case 2:
					fmt.Print("v")
				case 3:
					fmt.Print("<")
				}
			} else {
				if m[p] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}
