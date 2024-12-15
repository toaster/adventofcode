package deepsea

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/math"
)

// LanternfishWarehouse describes a lanternfish warehouse (https://adventofcode.com/2024/day/15).
type LanternfishWarehouse struct {
	boxes    map[math.Point2D]bool
	plan     math.Plan2D
	robot    math.Point2D
	scaledUp bool
}

// ParseLanternfishWarehouse creates a new LanternfishWarehouse from the given input
func ParseLanternfishWarehouse(lines []string) *LanternfishWarehouse {
	wh := &LanternfishWarehouse{
		boxes: map[math.Point2D]bool{},
	}
	wh.plan = math.ParsePlan2D(lines, '@', math.Plan2DWithUnknownTileHandler(func(pos math.Point2D, t rune) {
		wh.boxes[pos] = true
	}))
	wh.robot = wh.plan.Start
	return wh
}

// ParseScaledUpLanternfishWarehouse creates a scaled-up LanternfishWarehouse from the given input
func ParseScaledUpLanternfishWarehouse(lines []string) *LanternfishWarehouse {
	var newLines []string
	for _, line := range lines {
		newLine := ""
		for _, c := range line {
			newLine += string(c)
			switch c {
			case '@', 'O':
				newLine += "."
			default:
				newLine += string(c)
			}
		}
		newLines = append(newLines, newLine)
	}
	w := ParseLanternfishWarehouse(newLines)
	w.scaledUp = true
	return w
}

// MoveRobot moves the robot according to the given moves.
func (w *LanternfishWarehouse) MoveRobot(moves string) {
	for _, move := range moves {
		heading := math.Heading(move)
		target := heading.Facing(w.robot)
		var boxTarget *math.Point2D
		if w.boxes[target] {
			boxTarget = &target
		} else if w.scaledUp {
			left := target.AddXY(-1, 0)
			if w.boxes[left] {
				boxTarget = &left
			}
		}
		if !w.plan.Blocked(target) && (boxTarget == nil || w.moveBox(*boxTarget, heading)) {
			w.robot = target
		}
	}
}

// Print prints the warehouse map.
func (w *LanternfishWarehouse) Print() {
	w.plan.Print(math.Plan2DPrintAt(func(pos math.Point2D) bool {
		if w.scaledUp {
			if w.boxes[pos] {
				fmt.Print("[")
				return true
			}

			if w.boxes[pos.SubtractXY(1, 0)] {
				fmt.Print("]")
				return true
			}
		} else {
			if w.boxes[pos] {
				fmt.Print("O")
				return true
			}
		}

		if pos == w.robot {
			fmt.Print("@")
			return true
		}

		return false
	}))
}

// SumUpGPSCoordinates sums up the GPS coordinates of all the boxes.
func (w *LanternfishWarehouse) SumUpGPSCoordinates() int {
	sum := 0
	for p := range w.boxes {
		sum += p.Y*100 + p.X
	}
	return sum
}

func (w *LanternfishWarehouse) moveBox(pos math.Point2D, heading math.Heading) bool {
	if !w.scaledUp {
		target := heading.Facing(pos)
		if !w.plan.Blocked(target) && (!w.boxes[target] || w.moveBox(target, heading)) {
			w.boxes[target] = true
			delete(w.boxes, pos)
			return true
		}

		return false
	}

	targetLeft := heading.Facing(pos)
	targetRight := targetLeft.AddXY(1, 0)
	if w.plan.Blocked(targetLeft) || w.plan.Blocked(targetRight) {
		return false
	}

	boxTargetLeft := targetLeft.SubtractXY(1, 0)
	switch heading {
	case math.West:
		if !w.boxes[boxTargetLeft] || w.moveBox(boxTargetLeft, heading) {
			w.boxes[targetLeft] = true
			delete(w.boxes, pos)
			return true
		}
	case math.East:
		if !w.boxes[targetRight] || w.moveBox(targetRight, heading) {
			w.boxes[targetLeft] = true
			delete(w.boxes, pos)
			return true
		}
	case math.North, math.South:
		if w.boxes[targetLeft] {
			if w.moveBox(targetLeft, heading) {
				w.boxes[targetLeft] = true
				delete(w.boxes, pos)
				return true
			}
			return false
		}
		if !w.boxes[boxTargetLeft] && !w.boxes[targetRight] {
			w.boxes[targetLeft] = true
			delete(w.boxes, pos)
			return true
		}
		if !w.boxes[boxTargetLeft] && w.boxes[targetRight] {
			if w.moveBox(targetRight, heading) {
				w.boxes[targetLeft] = true
				delete(w.boxes, pos)
				return true
			}
			return false
		}
		if w.boxes[boxTargetLeft] && !w.boxes[targetRight] {
			if w.moveBox(boxTargetLeft, heading) {
				w.boxes[targetLeft] = true
				delete(w.boxes, pos)
				return true
			}
			return false
		}
		if w.scaledUpBoxMovableVertically(boxTargetLeft, heading) && w.moveBox(targetRight, heading) {
			w.moveBox(boxTargetLeft, heading)
			w.boxes[targetLeft] = true
			delete(w.boxes, pos)
			return true
		}
	}
	return false
}

func (w *LanternfishWarehouse) scaledUpBoxMovableVertically(pos math.Point2D, heading math.Heading) bool {
	targetLeft := heading.Facing(pos)
	targetRight := targetLeft.AddXY(1, 0)
	if w.plan.Blocked(targetLeft) || w.plan.Blocked(targetRight) {
		return false
	}

	if w.boxes[targetLeft] {
		return w.scaledUpBoxMovableVertically(targetLeft, heading)
	}

	boxTargetLeft := targetLeft.SubtractXY(1, 0)
	return (!w.boxes[boxTargetLeft] || w.scaledUpBoxMovableVertically(boxTargetLeft, heading)) &&
		(!w.boxes[targetRight] || w.scaledUpBoxMovableVertically(targetRight, heading))
}
