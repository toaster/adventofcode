package reflectordish

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/math"
)

const (
	ball rockType = 'O'
	cube rockType = '#'
	none rockType = '.'
)

// ParseControlPlatform parses a ControlPlatform from the given input.
func ParseControlPlatform(input []string) *ControlPlatform {
	p := &ControlPlatform{plan: map[math.Point2D]rockType{}}
	for y, line := range input {
		if p.width == 0 {
			p.width = len(line)
		}
		p.height++
		for x, r := range []rockType(line) {
			p.plan[math.Point2D{X: x, Y: y}] = r
		}
	}
	return p
}

// ControlPlatform is a control platform wor a reflector dish
type ControlPlatform struct {
	width  int
	height int
	plan   map[math.Point2D]rockType
}

// NorthLoad computes the load on the north side of the platform.
func (p *ControlPlatform) NorthLoad() int {
	load := 0
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			if p.plan[math.Point2D{X: x, Y: y}] == ball {
				load += p.height - y
			}
		}
	}
	return load
}

// Print prints the platform.
func (p *ControlPlatform) Print() {
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			fmt.Print(string(p.plan[math.Point2D{X: x, Y: y}]))
		}
		fmt.Println()
	}
}

// SpinCycle performs a “spin cycle” by tilting north, west, south and east.
func (p *ControlPlatform) SpinCycle() {
	p.TiltNorth()
	p.TiltWest()
	p.TiltSouth()
	p.TiltEast()
}

// TiltEast tilts the platform to east.
func (p *ControlPlatform) TiltEast() {
	unchanged := false
	for !unchanged {
		unchanged = true
		for y := 0; y < p.height; y++ {
			for x := p.height - 2; x >= 0; x-- {
				c := math.Point2D{X: x, Y: y}
				n := math.Point2D{X: x + 1, Y: y}
				if p.plan[c] == ball && p.plan[n] == none {
					p.plan[n] = ball
					p.plan[c] = none
					unchanged = false
				}
			}
		}
	}
}

// TiltNorth tilts the platform to north.
func (p *ControlPlatform) TiltNorth() {
	unchanged := false
	for !unchanged {
		unchanged = true
		for x := 0; x < p.width; x++ {
			for y := 1; y < p.height; y++ {
				c := math.Point2D{X: x, Y: y}
				n := math.Point2D{X: x, Y: y - 1}
				if p.plan[c] == ball && p.plan[n] == none {
					p.plan[n] = ball
					p.plan[c] = none
					unchanged = false
				}
			}
		}
	}
}

// TiltSouth tilts the platform to south.
func (p *ControlPlatform) TiltSouth() {
	unchanged := false
	for !unchanged {
		unchanged = true
		for x := 0; x < p.width; x++ {
			for y := p.height - 2; y >= 0; y-- {
				c := math.Point2D{X: x, Y: y}
				n := math.Point2D{X: x, Y: y + 1}
				if p.plan[c] == ball && p.plan[n] == none {
					p.plan[n] = ball
					p.plan[c] = none
					unchanged = false
				}
			}
		}
	}
}

// TiltWest tilts the platform to north.
func (p *ControlPlatform) TiltWest() {
	unchanged := false
	for !unchanged {
		unchanged = true
		for y := 0; y < p.height; y++ {
			for x := 0; x < p.height; x++ {
				c := math.Point2D{X: x, Y: y}
				n := math.Point2D{X: x - 1, Y: y}
				if p.plan[c] == ball && p.plan[n] == none {
					p.plan[n] = ball
					p.plan[c] = none
					unchanged = false
				}
			}
		}
	}
}

type rockType rune
