package main

import (
	"fmt"
	"math"

	"github.com/toaster/advent_of_code/internal/io"
	aocmath "github.com/toaster/advent_of_code/internal/math"
)

func main() {
	area := map[aocmath.Point2D]*elf{}
	for y, line := range io.ReadLines() {
		for x, c := range line {
			if c == '#' {
				pos := aocmath.Point2D{X: x, Y: y}
				area[pos] = &elf{pos: pos, proposedPos: pos}
			}
		}
	}

	last := &direction{
		checkA: aocmath.Point2D{X: 1, Y: -1},
		checkB: aocmath.Point2D{X: 1, Y: 1},
		target: aocmath.Point2D{X: 1, Y: 0},
	}
	cur := &direction{
		checkA: aocmath.Point2D{X: -1, Y: -1},
		checkB: aocmath.Point2D{X: 1, Y: -1},
		target: aocmath.Point2D{X: 0, Y: -1},
		next: &direction{
			checkA: aocmath.Point2D{X: -1, Y: 1},
			checkB: aocmath.Point2D{X: 1, Y: 1},
			target: aocmath.Point2D{X: 0, Y: 1},
			next: &direction{
				checkA: aocmath.Point2D{X: -1, Y: -1},
				checkB: aocmath.Point2D{X: -1, Y: 1},
				target: aocmath.Point2D{X: -1, Y: 0},
				next:   last,
			},
		},
	}
	last.next = cur

	rounds := 0
	for {
		rounds++
		newArea := map[aocmath.Point2D]*elf{}
		anyMove := false
		// printArea(area, -5, -5, 15, 15)
		for _, e := range area {
			move := false
			for _, n := range e.pos.Adjacents() {
				if area[n] != nil {
					move = true
					anyMove = true
					break
				}
			}
			if move {
				p := cur
				for {
					proposal := e.pos.Add(p.target)
					if area[proposal] == nil &&
						area[e.pos.Add(p.checkA)] == nil &&
						area[e.pos.Add(p.checkB)] == nil {
						e.proposedPos = proposal
						break
					}
					p = p.next
					if p == cur {
						break
					}
				}
			}
		}

		for _, e := range area {
			other := newArea[e.proposedPos]
			if other != nil {
				other.proposedPos = other.pos
				e.proposedPos = e.pos
			} else {
				newArea[e.proposedPos] = e
			}
		}

		for pos, e := range area {
			if e.pos == e.proposedPos {
				continue
			}

			e.pos = e.proposedPos
			area[e.pos] = e
			delete(area, pos)
		}
		cur = cur.next
		if !anyMove {
			break
		}
	}

	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for pos := range area {
		if pos.X < minX {
			minX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		}
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}
	fmt.Println((maxX+1-minX)*(maxY+1-minY)-len(area), rounds)
}

func printArea(area map[aocmath.Point2D]*elf, minX int, minY int, maxX int, maxY int) {
	for y := minY; y < maxY+1; y++ {
		for x := minX; x < maxX+1; x++ {
			if area[aocmath.Point2D{X: x, Y: y}] != nil {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type elf struct {
	pos         aocmath.Point2D
	proposedPos aocmath.Point2D
}

type direction struct {
	checkA aocmath.Point2D
	checkB aocmath.Point2D
	next   *direction
	target aocmath.Point2D
}
