package pluto

import (
	"fmt"
	"math"
	"strings"
)

// Maze represents a donut shaped Pluto maze.
type Maze struct {
	plan     map[pos]bool
	outer    map[string]pos
	outerPos map[pos]string
	inner    map[string]pos
	innerPos map[pos]string
	tree     *node
	nodes    map[string]*node
	// start *pathNode
	// outer map[string]*pathNode
	// inner map[string]*pathNode
}

// Parse returns a Maze based on the input.
func Parse(input string) *Maze {
	m := &Maze{outer: map[string]pos{}, outerPos: map[pos]string{}, inner: map[string]pos{}, innerPos: map[pos]string{}, plan: map[pos]bool{}, nodes: map[string]*node{}}
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	height := len(lines) - 4
	width := len(lines[2]) - 2
	for i, r := range lines[0] {
		if r != ' ' {
			m.setOuter(string([]rune{r, rune(lines[1][i])}), pos{i - 2, 0})
		}
	}
	for i, r := range lines[height+2] {
		if r != ' ' {
			m.setOuter(string([]rune{r, rune(lines[height+3][i])}), pos{i - 2, height - 1})
		}
	}
	for y := 0; y < height; y++ {
		l := lines[y+2]
		if l[0] != ' ' {
			m.setOuter(l[0:2], pos{0, y})
		}
		if len(l) > width+2 {
			m.setOuter(l[width+2:width+4], pos{width - 1, y})
		}
		for x, r := range l[2 : 2+width] {
			if r != '#' && r != '.' && r != ' ' {
				if lines[y+1][x+2] == '.' {
					m.setInner(string([]rune{r, rune(lines[y+3][x+2])}), pos{x, y - 1})
				}
				if lines[y+4][x+2] == '.' && lines[y+3][x+2] != '.' && lines[y+3][x+2] != '#' {
					m.setInner(string([]rune{r, rune(lines[y+3][x+2])}), pos{x, y + 2})
				}
				if l[x+1] == '.' {
					m.setInner(l[x+2:x+4], pos{x - 1, y})
				}
				if l[x+4] == '.' && l[x+3] != '.' && l[x+3] != '#' {
					m.setInner(l[x+2:x+4], pos{x + 2, y})
				}
			}
			if r == '.' {
				m.plan[pos{x, y}] = true
			}
		}
	}
	m.tree = &node{name: "AA", neighbours: map[*node]int{}}
	m.nodes["AA"] = m.tree
	for name := range m.outer {
		if m.nodes[name] != nil {
			continue
		}
		m.nodes[name] = &node{
			name:       name,
			neighbours: map[*node]int{},
			dist:       math.MaxInt64,
		}
	}
	for name := range m.inner {
		if m.nodes[name] != nil {
			continue
		}
		m.nodes[name] = &node{
			name:       name,
			neighbours: map[*node]int{},
			dist:       math.MaxInt64,
		}
	}
	for name, p := range m.outer {
		m.explore(m.nodes[name], p)
	}
	for name, p := range m.inner {
		m.explore(m.nodes[name], p)
	}
	return m
}

// ShortestPath returns the length of the shortest path from AA to ZZ
func (m *Maze) ShortestPath() int {
	q := map[*node]bool{}
	for _, n := range m.nodes {
		q[n] = true
	}
	for len(q) > 0 {
		min := math.MaxInt64
		var minN *node
		for n := range q {
			if n.dist < min {
				minN = n
				min = n.dist
			}
		}
		var t []string
		for n := range minN.neighbours {
			t = append(t, n.name)
		}
		fmt.Println("select", *minN, "->", t)
		delete(q, minN)
		for n, i := range minN.neighbours {
			if q[n] {
				d := minN.dist + 1 + i
				if d < n.dist {
					fmt.Println("reduce", n.dist, "to", d, "for", n.name, "from", minN.name)
					n.dist = d
					n.predecessor = minN
				} else {
					fmt.Println("don't reduce", n.dist, "to", d, "for", n.name, "from", minN.name)
				}
			}
		}
	}
	n := m.nodes["ZZ"]
	for n != nil {
		fmt.Println(*n)
		n = n.predecessor
	}
	return m.nodes["ZZ"].dist - 1
}

type node struct {
	name        string
	predecessor *node
	dist        int
	neighbours  map[*node]int
}

type pos struct {
	x int
	y int
}

type pathNode struct {
	p        pos
	children [4]*pathNode
	parent   int
	depth    int
}

const (
	north int = iota
	south
	west
	east
)

func (m *Maze) setOuter(k string, p pos) {
	m.outer[k] = p
	m.outerPos[p] = k
}

func (m *Maze) setInner(k string, p pos) {
	m.inner[k] = p
	m.innerPos[p] = k
}

func (m *Maze) explore(start *node, startPos pos) {
	fmt.Println("explore", start.name, startPos)
	cur := &pathNode{p: startPos, parent: -1}
	move := func(cur *pathNode, dir int) *pathNode {
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

			n = &pathNode{p: nextPos, parent: parentDir, depth: cur.depth + 1}
			n.children[parentDir] = cur
			cur.children[dir] = n

			if !m.plan[n.p] {
				return cur
			}
		}
		name := m.innerPos[n.p]
		// jumpPos := m.outer[name]
		if name == "" {
			name = m.outerPos[n.p]
			// jumpPos = m.inner[name]
		}
		if name != "" && name != start.name {
			fmt.Println("connect", start.name, "with", name, "dist", n.depth) //, "jumper to", jumpPos)
			neighbour := m.nodes[name]
			// if neighbour == nil {
			// 	neighbour = &node{
			// 		name:       name,
			// 		neighbours: map[*node]int{},
			// 		dist:       math.MaxInt64,
			// 	}
			// 	m.nodes[name] = neighbour
			// 	// if (jumpPos != pos{0, 0}) {
			// 	// 	m.explore(neighbour, jumpPos)
			// 	// }
			// }
			start.neighbours[neighbour] = n.depth
			neighbour.neighbours[start] = n.depth
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
