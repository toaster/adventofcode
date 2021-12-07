package space

import (
	"strconv"
	"strings"

	math2019 "github.com/toaster/advent_of_code/2019/math"
	"github.com/toaster/advent_of_code/internal/math"
)

// System represents a system of moons.
type System struct {
	Moons []*Moon
}

// Moon represents a moon in a system.
type Moon struct {
	Pos Vect
	Vel Vect
}

// Vect is a three dimensional vector used for position or velocity.
type Vect struct {
	X int
	Y int
	Z int
}

// ParseSystem parses the input and returns the System it describes.
func ParseSystem(input string) *System {
	s := &System{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		coords := strings.Split(line[1:len(line)-1], ",")
		s.Moons = append(s.Moons, &Moon{Pos: Vect{
			X: parseCoord(coords[0]),
			Y: parseCoord(coords[1]),
			Z: parseCoord(coords[2]),
		}})
	}
	return s
}

func parseCoord(coord string) int {
	c, _ := strconv.Atoi(strings.Split(coord, "=")[1])
	return c
}

// ComputePeriod returns the amount of steps after which the system returns to its initial state.
func (s *System) ComputePeriod() int {
	ix, iy, iz := s.getState()
	var px, py, pz int
	for i := 1; ; i++ {
		s.step()
		cx, cy, cz := s.getState()
		if px == 0 && s.statesEqual(cx, ix) {
			px = i
		}
		if py == 0 && s.statesEqual(cy, iy) {
			py = i
		}
		if pz == 0 && s.statesEqual(cz, iz) {
			pz = i
		}
		if px != 0 && py != 0 && pz != 0 {
			break
		}
	}
	return math2019.LCM([]int{px, py, pz})
}

func (s *System) getState() (sx, sy, sz [][]int) {
	for _, moon := range s.Moons {
		sx = append(sx, []int{moon.Pos.X, moon.Vel.X})
		sy = append(sy, []int{moon.Pos.Y, moon.Vel.Y})
		sz = append(sz, []int{moon.Pos.Z, moon.Vel.Z})
	}
	return
}

func (s *System) statesEqual(a, b [][]int) bool {
	for i, s := range a {
		o := b[i]
		if s[0] != o[0] || s[1] != o[1] {
			return false
		}
	}
	return true
}

// Energy computes the current energy of the system.
func (s *System) Energy() int {
	e := 0
	for _, moon := range s.Moons {
		pe := 0
		pe += math.AbsInt(moon.Pos.X)
		pe += math.AbsInt(moon.Pos.Y)
		pe += math.AbsInt(moon.Pos.Z)
		ke := 0
		ke += math.AbsInt(moon.Vel.X)
		ke += math.AbsInt(moon.Vel.Y)
		ke += math.AbsInt(moon.Vel.Z)
		e += pe * ke
	}
	return e
}

// Simulate simulates the effects of gravity in the system.
func (s *System) Simulate(steps int) {
	for i := 0; i < steps; i++ {
		s.step()
	}
}

func (s *System) step() {
	for j, moon := range s.Moons {
		for k, other := range s.Moons {
			if j == k {
				continue
			}
			moon.Vel.X += computeVelocity(moon.Pos.X, other.Pos.X)
			moon.Vel.Y += computeVelocity(moon.Pos.Y, other.Pos.Y)
			moon.Vel.Z += computeVelocity(moon.Pos.Z, other.Pos.Z)
		}
	}
	for _, moon := range s.Moons {
		moon.Pos.X += moon.Vel.X
		moon.Pos.Y += moon.Vel.Y
		moon.Pos.Z += moon.Vel.Z
	}
}

func computeVelocity(a, b int) int {
	if a < b {
		return 1
	}
	if a > b {
		return -1
	}
	return 0
}
