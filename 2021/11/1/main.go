package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	width := 10
	height := 10
	c := cavern{}
	for y, line := range io.ReadLines() {
		for x, e := range io.ParseInts(line, "") {
			var neighbours []point
			if x > 0 {
				neighbours = append(neighbours, point{x - 1, y})
				if y > 0 {
					neighbours = append(neighbours, point{x - 1, y - 1})
				}
				if y < height-1 {
					neighbours = append(neighbours, point{x - 1, y + 1})
				}
			}
			if x < width-1 {
				neighbours = append(neighbours, point{x + 1, y})
				if y > 0 {
					neighbours = append(neighbours, point{x + 1, y - 1})
				}
				if y < height-1 {
					neighbours = append(neighbours, point{x + 1, y + 1})
				}
			}
			if y > 0 {
				neighbours = append(neighbours, point{x, y - 1})
			}
			if y < height-1 {
				neighbours = append(neighbours, point{x, y + 1})
			}
			c[point{x, y}] = &octopus{energy: e, neighbours: neighbours}
		}
	}

	flashes := 0
	for i := 0; i < 100; i++ {
		for _, o := range c {
			o.energy++
		}
		for _, o := range c {
			flashes += flashIfNecessary(c, o)
		}
		for _, o := range c {
			if o.flashed {
				o.flashed = false
				o.energy = 0
			}
		}
	}
	fmt.Println(flashes)
}

func flashIfNecessary(c cavern, o *octopus) (flashes int) {
	if o.energy > 9 && !o.flashed {
		o.flashed = true
		flashes++
		for _, p := range o.neighbours {
			neighbour := c[p]
			neighbour.energy++
			flashes += flashIfNecessary(c, neighbour)
		}
	}
	return
}

type cavern map[point]*octopus

type point struct {
	x int
	y int
}

type octopus struct {
	energy     int
	flashed    bool
	neighbours []point
}
