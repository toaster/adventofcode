package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	floor := map[math.Point2D]*cucumber{}
	var movableEast []*cucumber
	var movableSouth []*cucumber
	var height int
	var width int
	for y, line := range io.ReadLines() {
		width = len(line)
		height++
		for x, tile := range line {
			p := math.Point2D{X: x, Y: y}
			switch tile {
			case '>':
				floor[p] = &cucumber{pos: p, south: false}
			case 'v':
				floor[p] = &cucumber{pos: p, south: true}
			}
		}
	}
	for _, cuke := range floor {
		p := cuke.destination(width, height)
		if floor[p] != nil {
			continue
		}

		if !cuke.south {
			movableEast = append(movableEast, cuke)
			continue
		}

		if c := floor[west(p, width)]; c != nil && !c.south {
			continue
		}

		movableSouth = append(movableSouth, cuke)
	}

	steps := 0
	for len(movableEast) > 0 || len(movableSouth) > 0 {
		var newMovableEast []*cucumber
		for _, cuke := range movableEast {
			prev := cuke.pos
			cuke.pos = cuke.destination(width, height)
			delete(floor, prev)
			floor[cuke.pos] = cuke
			p := cuke.destination(width, height)
			if floor[p] == nil {
				if c := floor[north(p, height)]; c == nil || !c.south {
					newMovableEast = append(newMovableEast, cuke)
				}
			}
			if c := floor[north(prev, height)]; c != nil && c.south {
				movableSouth = append(movableSouth, c)
			} else if c := floor[west(prev, width)]; c != nil && !c.south {
				newMovableEast = append(newMovableEast, c)
			}
		}
		movableEast = newMovableEast

		var newMovableSouth []*cucumber
		for _, cuke := range movableSouth {
			prev := cuke.pos
			cuke.pos = cuke.destination(width, height)
			delete(floor, prev)
			floor[cuke.pos] = cuke
			p := cuke.destination(width, height)
			if floor[p] == nil {
				if c := floor[west(p, width)]; c == nil || c.south {
					newMovableSouth = append(newMovableSouth, cuke)
				}
			}
			if c := floor[west(prev, width)]; c != nil && !c.south {
				movableEast = append(movableEast, c)
			} else if c := floor[north(prev, height)]; c != nil && c.south {
				newMovableSouth = append(newMovableSouth, c)
			}
			movableSouth = newMovableSouth
		}
		steps++
	}
	fmt.Println(steps)
}

func printFloor(width int, height int, floor map[math.Point2D]*cucumber) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := floor[math.Point2D{X: x, Y: y}]
			if c == nil {
				fmt.Print(".")
			} else if c.south {
				fmt.Print("v")
			} else {
				fmt.Print(">")
			}
		}
		fmt.Println()
	}
}

func north(pos math.Point2D, height int) math.Point2D {
	p := pos.SubtractXY(0, 1)
	if p.Y < 0 {
		p.Y = height - 1
	}
	return p
}

func west(pos math.Point2D, width int) math.Point2D {
	p := pos.SubtractXY(1, 0)
	if p.X < 0 {
		p.X = width - 1
	}
	return p
}

type cucumber struct {
	pos   math.Point2D
	south bool
}

func (c *cucumber) String() string {
	if c.south {
		return fmt.Sprintf("%d,%d: v", c.pos.X, c.pos.Y)
	}

	return fmt.Sprintf("%d,%d: >", c.pos.X, c.pos.Y)
}

func (c *cucumber) destination(width, height int) math.Point2D {
	if c.south {
		p := c.pos.AddXY(0, 1)
		if p.Y >= height {
			p.Y = 0
		}
		return p
	}

	p := c.pos.AddXY(1, 0)
	if p.X >= width {
		p.X = 0
	}
	return p
}
