package repairdroid

import (
	"fmt"

	"github.com/toaster/advent_of_code/2019/icc"
	"github.com/toaster/advent_of_code/2019/space"
)

// Explore takes an ICC program, starts a robot with it and returns it after it has explored the map.
func Explore(program []int) *Robot {
	in := make(chan int)
	out := make(chan int)
	start := &node{typ: floor}
	r := &Robot{map[space.Point]*node{start.pos: start}, start, start, nil, in, out}
	c := icc.New(in, out)
	c.Load(program)
	go c.Run()

	for {
		if r.cur.typ == oxygenSystem {
			r.oxy = r.cur
		}
		var dir int
		if r.cur.children[north] == nil {
			dir = north
		} else if r.cur.children[south] == nil {
			dir = south
		} else if r.cur.children[west] == nil {
			dir = west
		} else if r.cur.children[east] == nil {
			dir = east
		} else {
			dir = r.cur.parent
		}
		if dir == 0 {
			break
		}
		r.tryMove(dir)
	}
	return r
}

const (
	north int = iota + 1
	south
	west
	east
)

type node struct {
	pos      space.Point
	typ      tile
	children [5]*node
	parent   int
	depth    int
	oxyDist  int
}

type tile string

const (
	unexplored   tile = " "
	floor        tile = "."
	wall         tile = "#"
	oxygenSystem tile = "o"
	droid        tile = "D"
	oxygen       tile = "O"
)

// Robot is a repair droid.
type Robot struct {
	Plan  map[space.Point]*node
	start *node
	cur   *node
	oxy   *node
	in    chan int
	out   chan int
}

// PrintMap prints the current map to stdout.
func (r *Robot) PrintMap() {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	for p := range r.Plan {
		if p.X < minX {
			minX = p.X
		} else if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		} else if p.Y > maxY {
			maxY = p.Y
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := space.Point{X: x, Y: y}
			if p == r.cur.pos {
				fmt.Print(droid)
			} else {
				if r.Plan[p] == nil {
					fmt.Print(unexplored)
				} else if r.Plan[p].oxyDist > -1 {
					fmt.Print(oxygen)
				} else {
					fmt.Print(r.Plan[p].typ)
				}
			}
		}
		fmt.Println()
	}
}

// OxygenDistance returns the distance of the oxygen generator to the robot's start position.
func (r *Robot) OxygenDistance() int {
	return r.oxy.depth
}

// FillWithOxygen starts the oxygen generator and returns the duration it took to fill the whole map.
func (r *Robot) FillWithOxygen() int {
	s := r.oxy
	s.oxyDist = 0
	return r.fill(s)
}

func (r *Robot) fill(n *node) int {
	max := n.oxyDist
	for _, child := range n.children {
		if child == nil || child.oxyDist > -1 || child.typ == wall {
			continue
		}
		child.oxyDist = n.oxyDist + 1
		t := r.fill(child)
		if t > max {
			max = t
		}
	}
	return max
}

func (r *Robot) tryMove(dir int) {
	r.in <- dir
	result := <-r.out

	n := r.cur.children[dir]
	if n == nil {
		nextPos := r.cur.pos
		switch dir {
		case north:
			nextPos.Y--
		case south:
			nextPos.Y++
		case west:
			nextPos.X--
		case east:
			nextPos.X++
		}
		var parentDir int
		if dir%2 == 1 {
			parentDir = dir + 1
		} else {
			parentDir = dir - 1
		}

		n = &node{pos: nextPos, parent: parentDir, depth: r.cur.depth + 1, oxyDist: -1}
		n.children[parentDir] = r.cur
		r.Plan[nextPos] = n
		r.cur.children[dir] = n
		switch result {
		case 0:
			n.typ = wall
			return
		case 1:
			n.typ = floor
		case 2:
			n.typ = oxygenSystem
		}
	}

	r.cur = n
}
