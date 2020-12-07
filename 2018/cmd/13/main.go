package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	north byte = '^'
	east  byte = '>'
	south byte = 'v'
	west  byte = '<'
)

type turn int

const (
	left turn = iota
	straight
	right
	direction
)

type pos struct {
	x, y int
}

type cart struct {
	h  byte
	nt turn
	i  bool
}

const (
	none         byte = ' '
	westEast     byte = '-'
	westSouth    byte = '\\'
	northSouth   byte = '|'
	northWest    byte = '/'
	intersection byte = '+'
)

func main() {
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")

	h := len(lines)
	carts := map[pos]*cart{}
	tracks := make([][]byte, h)
	for y, line := range lines {
		tracks[y] = parseLine(line, y, carts)
	}
	printMap(tracks, carts)
	for done := false; !done; {
		for y, l := range tracks {
			for x := range l {
				c := carts[pos{x, y}]
				if c == nil {
					continue
				}
				if c.i {
					c.i = false
					continue
				}
				delete(carts, pos{x, y})
				np := pos{x, y}
				switch c.h {
				case east:
					np.x++
					c.i = true
				case west:
					np.x--
				case north:
					np.y--
				case south:
					np.y++
					c.i = true
				}
				if carts[np] != nil {
					fmt.Println("crash", np.x, np.y)
					delete(carts, np)
					fmt.Println("remaining", len(carts))
					if len(carts) < 2 {
						for p, c := range carts {
							fmt.Println(p, *c)
						}
						done = true
					}
					continue
				} else {
					carts[np] = c
				}

				switch tracks[np.y][np.x] {
				case westSouth:
					switch c.h {
					case east:
						c.h = south
					case north:
						c.h = west
					case west:
						c.h = north
					case south:
						c.h = east
					}
				case northWest:
					switch c.h {
					case south:
						c.h = west
					case north:
						c.h = east
					case east:
						c.h = north
					case west:
						c.h = south
					}
				case intersection:
					switch c.h {
					case south:
						if c.nt%direction == left {
							c.h = east
						} else if c.nt%direction == right {
							c.h = west
						}
					case north:
						if c.nt%direction == left {
							c.h = west
						} else if c.nt%direction == right {
							c.h = east
						}
					case east:
						if c.nt%direction == left {
							c.h = north
						} else if c.nt%direction == right {
							c.h = south
						}
					case west:
						if c.nt%direction == left {
							c.h = south
						} else if c.nt%direction == right {
							c.h = north
						}
					}
					c.nt++
				}
			}
		}
		printMap(tracks, carts)
	}
	for p, c := range carts {
		fmt.Println(p, *c)
	}
}

func parseLine(l string, y int, carts map[pos]*cart) []byte {
	tracks := []byte(l)
	for x, b := range tracks {
		switch b {
		case '>':
			tracks[x] = '-'
			carts[pos{x, y}] = &cart{h: east}
		case '<':
			tracks[x] = '-'
			carts[pos{x, y}] = &cart{h: west}
		case '^':
			tracks[x] = '|'
			carts[pos{x, y}] = &cart{h: north}
		case 'v':
			tracks[x] = '|'
			carts[pos{x, y}] = &cart{h: south}
		}
	}
	return tracks
}

func printMap(m [][]byte, carts map[pos]*cart) {
	return
	for y, l := range m {
		for x, t := range l {
			if c := carts[pos{x, y}]; c != nil {
				fmt.Print(string(c.h))
			} else {
				fmt.Print(string(t))
			}
		}
		fmt.Println("")
	}
	time.Sleep(1 * time.Second)
}
