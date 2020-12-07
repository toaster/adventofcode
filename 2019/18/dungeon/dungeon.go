package dungeon

import (
	"strings"
)

// Map is a map of a dungeon.
type Map struct {
	width  int
	height int
	plan   map[pos]tile
	keys   map[tile]pos
	doors  map[tile]pos
	start  *node
}

type pos struct {
	x int
	y int
}

type tile string

const (
	floor    tile = "."
	wall     tile = "#"
	entrance tile = "@"
)

type node struct {
	p        pos
	children [4]*node
	parent   int
	depth    int
}

const (
	north int = iota
	south
	west
	east
)

// Parse creates a Map from the input.
func Parse(input string) *Map {
	m := &Map{plan: map[pos]tile{}, keys: map[tile]pos{}, doors: map[tile]pos{}}
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		m.height = y + 1
		m.width = len(line)
		for x, t := range strings.Split(line, "") {
			p := pos{x, y}
			m.plan[p] = tile(t)
			if tile(t) == entrance {
				m.start = &node{p: p, parent: -1}
			}
		}
	}
	m.explore()
	return m
}

// MinimalStepsToCollectAllKeys computes how many steps are needed to gather all keys starting at the entrance.
func (m *Map) MinimalStepsToCollectAllKeys() int {
	return 0
}

func (m *Map) explore() {
	cur := m.start
	move := func(cur *node, dir int) *node {
		n := cur.children[dir]
		if n == nil {
			nextPos := cur.p
			switch dir {
			case north:
				nextPos.y--
			case south:
				nextPos.y++
			case west:
				nextPos.x--
			case east:
				nextPos.x++
			}
			var parentDir int
			if dir%2 == 0 {
				parentDir = dir + 1
			} else {
				parentDir = dir - 1
			}

			n = &node{p: nextPos, parent: parentDir, depth: cur.depth + 1}
			n.children[parentDir] = cur
			cur.children[dir] = n

			t := m.plan[n.p]
			switch t {
			case wall:
				return cur
			case floor, entrance:
				// nothing special
			default:
				if t[0] >= 'A' {
					m.doors[t] = n.p
				} else {
					m.keys[t] = n.p
				}
			}
		}
		return n
	}
	for {
		var moved bool
		for dir := 0; dir < 4; dir++ {
			if cur.children[dir] == nil {
				cur = move(cur, dir)
				moved = true
				break
			}
		}
		if !moved {
			if cur.parent == -1 {
				break
			}
			cur = move(cur, cur.parent)
		}
	}
}
