package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	input := lines[:len(lines)-2]
	tiles, start := parseMap(input)
	directions := parseDirections(lines[len(lines)-1])

	cur := tiles[start]
	heading := 0
	for _, d := range directions {
		for i := 0; i < d.move; i++ {
			var next *tile
			switch heading {
			case 0:
				next = cur.right
			case 1:
				next = cur.bottom
			case 2:
				next = cur.left
			case 3:
				next = cur.top
			}
			if !next.open {
				break
			}
			cur = next
		}
		heading = (heading + d.turn + 4) % 4
	}
	fmt.Println(1000*cur.pos.Y + 4*cur.pos.X + heading)
}

func connectTiles(tiles map[math.Point2D]*tile, maxX int, maxY int) {
	for p, t := range tiles {
		t.bottom = tiles[p.AddXY(0, 1)]
		for y := 1; t.bottom == nil; y++ {
			t.bottom = tiles[math.Point2D{X: p.X, Y: y}]
		}
		t.left = tiles[p.AddXY(-1, 0)]
		for x := maxX; t.left == nil; x-- {
			t.left = tiles[math.Point2D{X: x, Y: p.Y}]
		}
		t.right = tiles[p.AddXY(1, 0)]
		for x := 1; t.right == nil; x++ {
			t.right = tiles[math.Point2D{X: x, Y: p.Y}]
		}
		t.top = tiles[p.AddXY(0, -1)]
		for y := maxY; t.top == nil; y-- {
			t.top = tiles[math.Point2D{X: p.X, Y: y}]
		}
	}
}

func parseDirections(s string) (directions []direction) {
	for len(s) > 0 {
		if s[0] == 'R' {
			directions = append(directions, direction{0, 1})
			s = s[1:]
		} else if s[0] == 'L' {
			directions = append(directions, direction{0, -1})
			s = s[1:]
		}
		i := 1
		for ; i < len(s); i++ {
			if s[i] == 'R' || s[i] == 'L' {
				break
			}
		}
		directions = append(directions, direction{io.ParseInt(s[:i]), 0})
		s = s[i:]
	}
	return
}

func parseMap(input []string) (map[math.Point2D]*tile, math.Point2D) {
	var start math.Point2D
	tiles := map[math.Point2D]*tile{}
	maxX := 0
	maxY := 0
	for y, line := range input {
		if y+1 > maxY {
			maxY = y + 1
		}
		for x, c := range line {
			if x+1 > maxX {
				maxX = x + 1
			}
			if c == ' ' {
				continue
			}
			p := math.Point2D{X: x + 1, Y: y + 1}
			if len(tiles) == 0 {
				start = p
			}
			tiles[p] = &tile{pos: p, open: c == '.'}
		}
	}
	connectTiles(tiles, maxX, maxY)
	return tiles, start
}

type direction struct {
	move int
	turn int
}

type tile struct {
	bottom *tile
	left   *tile
	open   bool
	pos    math.Point2D
	right  *tile
	top    *tile
}
