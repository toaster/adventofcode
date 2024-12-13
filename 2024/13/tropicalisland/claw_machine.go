package tropicalisland

import (
	"math"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// ParseClawMachine creates a new ClawMachine from the given input adding the given offset to the
// X and Y values.
func ParseClawMachine(input string, offset int) ClawMachine {
	lines := strings.Split(input, "\n")
	ax, ay := parseLine(lines[0], 10)
	bx, by := parseLine(lines[1], 10)
	x, y := parseLine(lines[2], 7)
	return ClawMachine{
		ax: ax,
		ay: ay,
		bx: bx,
		by: by,
		x:  x + offset,
		y:  y + offset,
	}
}

// ClawMachine is an arcade on the tropical island (https://adventofcode.com/2024/day/13).
type ClawMachine struct {
	ax int
	ay int
	bx int
	by int
	x  int
	y  int
}

// ComputeMinimalCostsIfPossible returns the minimal costs to win with the claw machine or `0` if it
// is not possible to win.
func (m ClawMachine) ComputeMinimalCostsIfPossible() int {
	fA := (float64(m.x) - float64(m.bx*m.y)/float64(m.by)) / (float64(m.ax) - float64(m.bx*m.ay)/float64(m.by))
	fB := (float64(m.y) - fA*float64(m.ay)) / float64(m.by)
	nA := int(math.Round(fA))
	nB := int(math.Round(fB))
	cx := nA*m.ax + nB*m.bx
	cy := nA*m.ay + nB*m.by
	if cx != m.x || cy != m.y {
		return 0
	}

	return 3*nA + nB
}

func parseLine(line string, offset int) (int, int) {
	raw := strings.Split(line[offset:], ", ")
	return io.ParseInt(raw[0][2:]), io.ParseInt(raw[1][2:])
}
