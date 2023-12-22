package util

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/math"
)

// ParsePlan2D parses a Plan2D from the given input.
func ParsePlan2D(lines []string, startMarker rune) *Plan2D {
	plan := &Plan2D{blocked: map[math.Point2D]bool{}}
	for y, line := range lines {
		plan.height++
		if plan.width == 0 {
			plan.width = len(line)
		}
		for x, c := range line {
			switch c {
			case startMarker:
				plan.Start = math.Point2D{X: x, Y: y}
			case '#':
				plan.blocked[math.Point2D{X: x, Y: y}] = true
			}
		}
	}
	plan.rangeX = math.Range{End: plan.width - 1}
	plan.rangeY = math.Range{End: plan.height - 1}
	return plan
}

// Plan2D represents a two-dimensional area.
type Plan2D struct {
	Start math.Point2D

	blocked map[math.Point2D]bool
	height  int
	rangeX  math.Range
	rangeY  math.Range
	width   int
}

// Neighbours returns the reachable neighbours of a given point on the plan.
func (p *Plan2D) Neighbours(point math.Point2D) (neighbours []math.Point2D) {
	if p.rangeY.Covers(point.Y) {
		if point.X > p.rangeX.Start {
			n := point.AddXY(-1, 0)
			if !p.blocked[n] {
				neighbours = append(neighbours, n)
			}
		}
		if point.X < p.rangeX.End {
			n := point.AddXY(1, 0)
			if !p.blocked[n] {
				neighbours = append(neighbours, n)
			}
		}
	}
	if p.rangeX.Covers(point.X) {
		if point.Y > p.rangeY.Start {
			n := point.AddXY(0, -1)
			if !p.blocked[n] {
				neighbours = append(neighbours, n)
			}
		}
		if point.Y < p.rangeY.End {
			n := point.AddXY(0, 1)
			if !p.blocked[n] {
				neighbours = append(neighbours, n)
			}
		}
	}
	return
}

// Print prints the plan on stdout.
func (p *Plan2D) Print(isMarked map[math.Point2D]bool) {
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			pos := math.Point2D{X: x, Y: y}
			if p.blocked[pos] {
				fmt.Print("#")
			} else if isMarked[pos] {
				fmt.Print("O")
			} else if pos == p.Start {
				fmt.Print("S")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
