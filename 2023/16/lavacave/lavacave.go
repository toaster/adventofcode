package lavacave

import (
	"slices"

	"github.com/toaster/advent_of_code/internal/math"
)

// Direction constants
const (
	North Direction = iota
	East
	South
	West
)

// Parse parses the input into a Cave.
func Parse(input []string) *Cave {
	cave := &Cave{tiles: map[math.Point2D]*tile{}}
	for y, line := range input {
		if cave.width == 0 {
			cave.width = len(line)
		}
		cave.height++
		for x, c := range line {
			cave.tiles[math.Point2D{X: x, Y: y}] = &tile{typ: c, beams: map[Direction]bool{}}
		}
	}
	return cave
}

// TopLeftBeamHeadingEast returns a beam entering at the top left tile heading east.
func TopLeftBeamHeadingEast() *Beam {
	return &Beam{
		pos:     math.Point2D{X: -1},
		heading: East,
	}
}

// Beam describes a beam of light entering the cave.
type Beam struct {
	pos     math.Point2D
	heading Direction
}

// Cave represents a cave to produce lava.
type Cave struct {
	height int
	tiles  map[math.Point2D]*tile
	width  int
}

// CountEnergizedFloorTiles counts the floor tiles being energized when the beam enters the cave.
func (c *Cave) CountEnergizedFloorTiles(beam *Beam) int {
	for _, t := range c.tiles {
		t.beams = map[Direction]bool{}
	}
	beams := []*Beam{beam}
	for len(beams) > 0 {
		for i := 0; i < len(beams); i++ {
			b := beams[i]
			switch b.heading {
			case North:
				b.pos.Y--
			case East:
				b.pos.X++
			case South:
				b.pos.Y++
			case West:
				b.pos.X--
			}
			t := c.tiles[b.pos]
			if t == nil || t.beams[b.heading] {
				beams = slices.Delete(beams, i, i+1)
				i--
				continue
			}

			t.beams[b.heading] = true
			switch t.typ {
			case '/':
				switch b.heading {
				case North:
					b.heading = East
				case East:
					b.heading = North
				case South:
					b.heading = West
				case West:
					b.heading = South
				}
			case '\\':
				switch b.heading {
				case North:
					b.heading = West
				case East:
					b.heading = South
				case South:
					b.heading = East
				case West:
					b.heading = North
				}
			case '-':
				if b.heading == North || b.heading == South {
					beams = slices.Delete(beams, i, i+1)
					b1 := &Beam{pos: b.pos, heading: East}
					b2 := &Beam{pos: b.pos, heading: West}
					beams = slices.Insert(beams, i, b1, b2)
					i++
				}
			case '|':
				if b.heading == East || b.heading == West {
					beams = slices.Delete(beams, i, i+1)
					b1 := &Beam{pos: b.pos, heading: North}
					b2 := &Beam{pos: b.pos, heading: South}
					beams = slices.Insert(beams, i, b1, b2)
					i++
				}
			}
		}
	}
	energized := 0
	for _, t := range c.tiles {
		if len(t.beams) > 0 {
			energized++
		}
	}
	return energized
}

// PossibleInputBeams returns all possible input beams.
func (c *Cave) PossibleInputBeams() []*Beam {
	var beams []*Beam
	for x := 0; x < c.width; x++ {
		beams = append(beams, &Beam{math.Point2D{X: x, Y: -1}, South})
		beams = append(beams, &Beam{math.Point2D{X: x, Y: c.height}, North})
	}
	for y := 0; y < c.height; y++ {
		beams = append(beams, &Beam{math.Point2D{X: -1, Y: y}, East})
		beams = append(beams, &Beam{math.Point2D{X: c.width, Y: y}, West})
	}
	return beams
}

// Direction describes the direction a beam of light is heading.
type Direction int

type tile struct {
	typ   rune
	beams map[Direction]bool
}
