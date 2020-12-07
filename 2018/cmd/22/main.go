package main

import "fmt"

type pos struct {
	x, y int
}

const (
	narrow byte = '|'
	rocky  byte = '.'
	wet    byte = '='
)

type cavern struct {
	gi, el, rl int
	t          byte
}

func main() {
	depth := 11394
	target := pos{7, 701}
	width := target.x + 1
	height := target.y + 1
	m := map[pos]cavern{}
	trl := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := cavern{}
			if x == 0 {
				if y == 0 {
					c.gi = 0
				} else {
					c.gi = y * 48271
				}
			} else if y == 0 {
				c.gi = x * 16807
			} else if x == target.x && y == target.y {
				c.gi = 0
			} else {
				c.gi = m[pos{x - 1, y}].el * m[pos{x, y - 1}].el
			}
			c.el = (c.gi + depth) % 20183
			c.rl = c.el % 3
			switch c.rl {
			case 0:
				c.t = rocky
			case 1:
				c.t = wet
			case 2:
				c.t = narrow
			}
			m[pos{x, y}] = c
			trl += c.rl
		}
	}
	fmt.Println("risk level:", trl)
}
