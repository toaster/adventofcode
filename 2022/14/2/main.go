package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	a := parseScan(io.ReadLines())
	drawArea(a)
	var s *math.Point2D
	count := 0
	for {
		if s == nil {
			s = &math.Point2D{}
			*s = a.sandSource
		}
		t := s.AddXY(0, 1)
		if !isFloor(t, a) {
			if a.tiles[t] == air {
				s = &t
				continue
			}

			t = s.AddXY(-1, 1)
			if a.tiles[t] == air {
				s = &t
				continue
			}

			t = s.AddXY(1, 1)
			if a.tiles[t] == air {
				s = &t
				continue
			}
		}

		count++
		a.tiles[*s] = sand
		if *s == a.sandSource {
			break
		}

		s = nil
	}
	drawArea(a)
	fmt.Println(count)
}

func isFloor(p math.Point2D, a *area) bool {
	return p.Y == a.bottomRight.Y+2
}

func drawArea(a *area) {
	floorY := a.bottomRight.Y + 2
	for y := a.topLeft.Y; y <= floorY; y++ {
		for x := a.topLeft.X - 10; x <= a.bottomRight.X+10; x++ {
			p := math.Point2D{X: x, Y: y}
			switch a.tiles[p] {
			case rock:
				fmt.Print("#")
			case sand:
				fmt.Print("o")
			default:
				if p == a.sandSource {
					fmt.Print("+")
				} else if p.Y == floorY {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func parsePath(line string, a *area) {
	var last *math.Point2D
	for _, c := range strings.Split(line, " -> ") {
		coords := io.ParseInts(c, ",")
		p := math.Point2D{X: coords[0], Y: coords[1]}
		if p.X < a.topLeft.X {
			a.topLeft.X = p.X
		}
		if p.Y < a.topLeft.Y {
			a.topLeft.Y = p.Y
		}
		if p.X > a.bottomRight.X {
			a.bottomRight.X = p.X
		}
		if p.Y > a.bottomRight.Y {
			a.bottomRight.Y = p.Y
		}
		if last != nil {
			add := math.Point2D{}
			if last.X == p.X {
				if last.Y > p.Y {
					add.Y = -1
				} else {
					add.Y = 1
				}
			} else {
				if last.X > p.X {
					add.X = -1
				} else {
					add.X = 1
				}
			}
			for *last != p {
				*last = last.Add(add)
				a.tiles[*last] = rock
			}
		} else {
			a.tiles[p] = rock
		}
		last = &p
	}
}

func parseScan(lines []string) *area {
	src := math.Point2D{X: 500}
	a := &area{
		bottomRight: src,
		sandSource:  src,
		tiles:       map[math.Point2D]material{},
		topLeft:     src,
	}
	for _, line := range lines {
		parsePath(line, a)
	}
	return a
}

type area struct {
	topLeft     math.Point2D
	bottomRight math.Point2D
	sandSource  math.Point2D
	tiles       map[math.Point2D]material
}

type material int8

const (
	air material = iota
	rock
	sand
)
