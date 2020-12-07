package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

const (
	none    byte = 0
	door    byte = '/'
	floor   byte = '.'
	start   byte = 'X'
	unknown byte = '?'
	wall    byte = '#'
)

type dungeon struct {
	m    map[pos]byte
	n    map[pos]*node
	s    pos
	w, h int
}

type node struct {
	p pos
	c []*node
}

func main() {
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	input := []byte(strings.TrimSuffix(string(b), "\n"))

	d := dungeon{
		m: map[pos]byte{pos{0, 0}: start},
		n: map[pos]*node{pos{0, 0}: &node{pos{0, 0}, []*node{}}},
		s: pos{-1, -1},
		w: 3,
		h: 3,
	}

	parsePath(&d, d.n[pos{0, 0}], input[1:len(input)-1])
	fmt.Println(biggestLength(d.n[pos{0, 0}]) - 1)
	fmt.Println(countRooms(d.n[pos{0, 0}], 0, 1000))
	//printMap(&d, pos{0, 0})
	//fmt.Println(biggestLength(&d, pos{0, 0}, input[1:len(input)-1]))
}

func parsePath(d *dungeon, s *node, input []byte) int {
	// fmt.Println("check:", string(input))
	// time.Sleep(500 * time.Millisecond)
	n := s
	// printMap(d, p)
	for i := 0; i < len(input); i++ {
		p := n.p
		//		fmt.Println("cur:", string(input[i:]))
		switch input[i] {
		case '(':
			i++
			i += parsePath(d, n, input[i:])
		case ')':
			// fmt.Println("subreturn:", lengths, max(lengths))
			return i
		case '|':
			// fmt.Println("push", length)
			n = s
		case 'W':
			d.m[pos{p.x - 1, p.y}] = door
			d.m[pos{p.x - 1, p.y - 1}] = wall
			d.m[pos{p.x - 1, p.y + 1}] = wall
			p = pos{p.x - 2, p.y}
			if d.m[p] == none {
				d.m[p] = floor
				if p.x < d.s.x {
					d.s.x = p.x - 1
					d.w += 2
				}
				d.n[p] = &node{p, []*node{}}
				n.c = append(n.c, d.n[p])
			}
			n = d.n[p]
		case 'E':
			d.m[pos{p.x + 1, p.y}] = door
			d.m[pos{p.x + 1, p.y - 1}] = wall
			d.m[pos{p.x + 1, p.y + 1}] = wall
			p = pos{p.x + 2, p.y}
			if d.m[p] == none {
				d.m[p] = floor
				if p.x >= d.s.x+d.w {
					d.w += 2
				}
				d.n[p] = &node{p, []*node{}}
				n.c = append(n.c, d.n[p])
			}
			n = d.n[p]
		case 'N':
			d.m[pos{p.x, p.y - 1}] = door
			d.m[pos{p.x - 1, p.y - 1}] = wall
			d.m[pos{p.x + 1, p.y - 1}] = wall
			p = pos{p.x, p.y - 2}
			if d.m[p] == none {
				d.m[p] = floor
				if p.y < d.s.y {
					d.s.y = p.y - 1
					d.h += 2
				}
				d.n[p] = &node{p, []*node{}}
				n.c = append(n.c, d.n[p])
			}
			n = d.n[p]
		case 'S':
			d.m[pos{p.x, p.y + 1}] = door
			d.m[pos{p.x - 1, p.y + 1}] = wall
			d.m[pos{p.x + 1, p.y + 1}] = wall
			p = pos{p.x, p.y + 2}
			if d.m[p] == none {
				d.m[p] = floor
				if p.y >= d.s.y+d.h {
					d.h += 2
				}
				d.n[p] = &node{p, []*node{}}
				n.c = append(n.c, d.n[p])
			}
			n = d.n[p]
		}
		//		printMap(d, p)
	}
	// fmt.Println("return:", length, lastIndex)
	//printMap(d, n.p)
	return len(input)
}

func countRooms(n *node, dist, min int) int {
	count := 0
	if dist >= min {
		count = 1
	}
	for _, c := range n.c {
		count += countRooms(c, dist+1, min)
	}
	return count
}

func biggestLength(n *node) int {
	if len(n.c) == 0 {
		//fmt.Println("end at", n.p)
		return 1
	}
	subLengths := []int{}
	for _, c := range n.c {
		subLengths = append(subLengths, biggestLength(c))
	}
	// fmt.Println("max at", n.p, max(subLengths))
	return max(subLengths) + 1
}

func max(values []int) int {
	m := 0
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}

func printMap(d *dungeon, c pos) {
	for j := 0; j < d.h; j++ {
		y := d.s.y + j
		for i := 0; i < d.w; i++ {
			x := d.s.x + i
			p := pos{x, y}
			t := d.m[p]
			if t == 0 {
				t = unknown
			}
			if p == c {
				t = 'C'
			}
			fmt.Print(string(t))
		}
		fmt.Println("")
	}
	fmt.Println("")
}
