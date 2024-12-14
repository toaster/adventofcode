package ebhq

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/math"
)

// ParseHallway creates a new Hallway from the given input.
func ParseHallway(lines []string) Hallway {
	h := Hallway{tiles: map[math.Point2D]int{}}
	for _, line := range lines {
		r := parseRobot(line)
		h.robots = append(h.robots, r)
		if r.pos.X >= h.width {
			h.width = r.pos.X + 1
		}
		if r.pos.Y >= h.height {
			h.height = r.pos.Y + 1
		}
		h.tiles[r.pos]++
	}
	return h
}

// Hallway represents a hallway patrolled by robots (https://adventofcode.com/2024/day/14).
type Hallway struct {
	height int
	robots []*robot
	tiles  map[math.Point2D]int
	width  int
}

// ComputeSafetyFactor computes the safety factor of the Hallway after the given amount of seconds.
func (h *Hallway) ComputeSafetyFactor(seconds int) int {
	for i := 0; i < seconds; i++ {
		h.evolveOneSecond()
	}
	f := 1
	middleX := h.width / 2
	middleY := h.height / 2
	f *= h.countRobots(math.Rectangle2D{BottomRight: math.Point2D{X: middleX - 1, Y: middleY - 1}})
	f *= h.countRobots(math.Rectangle2D{TopLeft: math.Point2D{X: middleX + 1}, BottomRight: math.Point2D{X: h.width - 1, Y: middleY - 1}})
	f *= h.countRobots(math.Rectangle2D{TopLeft: math.Point2D{Y: middleY + 1}, BottomRight: math.Point2D{X: middleX - 1, Y: h.height - 1}})
	f *= h.countRobots(math.Rectangle2D{TopLeft: math.Point2D{X: middleX + 1, Y: middleY + 1}, BottomRight: math.Point2D{X: h.width - 1, Y: h.height - 1}})
	return f
}

// EvolveToXMasTree returns the number of seconds until the robots will form an X-mas tree on the hallway floor.
func (h *Hallway) EvolveToXMasTree() int {
	middleX := h.width / 2
	middleY := h.height / 2
	for i := 1; ; i++ {
		h.evolveOneSecond()
		if h.tiles[math.Point2D{X: middleX, Y: middleY}] > 0 &&
			h.tiles[math.Point2D{X: middleX - 1, Y: middleY}] > 0 &&
			h.tiles[math.Point2D{X: middleX + 1, Y: middleY}] > 0 {
			h.print()
			return i
		}
	}
}

func (h *Hallway) countRobots(quadrant math.Rectangle2D) int {
	count := 0
	for _, r := range h.robots {
		if r.pos.IsInside(quadrant) {
			count++
		}
	}
	return count
}

func (h *Hallway) evolveOneSecond() {
	for _, r := range h.robots {
		h.tiles[r.pos]--
		r.pos = r.pos.Add(math.Point2D(r.velocity))
		for r.pos.X < 0 {
			r.pos.X += h.width
		}
		for r.pos.Y < 0 {
			r.pos.Y += h.height
		}
		r.pos.X %= h.width
		r.pos.Y %= h.height
		h.tiles[r.pos]++
	}
}

func (h *Hallway) print() {
	for y := 0; y < h.height; y++ {
		for x := 0; x < h.width; x++ {
			if h.tiles[math.Point2D{X: x, Y: y}] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func parseRobot(line string) *robot {
	rawPosAndVelocity := strings.Split(line, " ")
	p := math.ParsePoint2D(rawPosAndVelocity[0][2:], ",")
	v := math.ParseVector2D(rawPosAndVelocity[1][2:], ",")
	return &robot{pos: p, velocity: v}
}

type robot struct {
	pos      math.Point2D
	velocity math.Vector2D
}
