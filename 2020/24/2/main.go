package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/toaster/advent_of_code/2019/math"
)

type point struct {
	x int
	y int
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	tiles := map[point]bool{}
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		var p point
		for len(line) > 0 {
			// fmt.Println("line", line, "p", p)
			if strings.HasPrefix(line, "se") {
				p = se(p)
				line = line[2:]
			} else if strings.HasPrefix(line, "sw") {
				p = sw(p)
				line = line[2:]
			} else if strings.HasPrefix(line, "w") {
				p = w(p)
				line = line[1:]
			} else if strings.HasPrefix(line, "nw") {
				p = nw(p)
				line = line[2:]
			} else if strings.HasPrefix(line, "ne") {
				p = ne(p)
				line = line[2:]
			} else if strings.HasPrefix(line, "e") {
				p = e(p)
				line = line[1:]
			}
		}
		// fmt.Println("flip", p)
		tiles[p] = !tiles[p]
	}
	// fmt.Println(tiles)
	// {
	// 	count := 0
	// 	for _, b := range tiles {
	// 		if b {
	// 			count++
	// 		}
	// 	}
	// 	fmt.Println("black tiles:", count)
	// }
	for d := 0; d < 100; d++ {
		flipPoints := map[point]bool{}
		for p, b := range tiles {
			if b {
				n := countBlackNeighbours(p, tiles)
				if n == 0 || n > 2 {
					flipPoints[p] = true
				}
				for _, np := range neighbours(p) {
					if !tiles[np] && countBlackNeighbours(np, tiles) == 2 {
						flipPoints[np] = true
					}
				}
			}
			// fmt.Println(p, b, "->", n, "=>", flipPoints)
		}
		for flipPoint := range flipPoints {
			tiles[flipPoint] = !tiles[flipPoint]
		}
		count := 0
		for _, b := range tiles {
			if b {
				count++
			}
		}
		fmt.Println("day:", d+1, "black tiles:", count)
	}
	count := 0
	for _, b := range tiles {
		if b {
			count++
		}
	}
	fmt.Println("black tiles:", count)
}

func neighbours(p point) []point {
	return []point{
		w(p),
		nw(p),
		ne(p),
		e(p),
		se(p),
		sw(p),
	}
}

func countBlackNeighbours(p point, tiles map[point]bool) int {
	count := 0
	if tiles[w(p)] {
		count++
	}
	if tiles[nw(p)] {
		count++
	}
	if tiles[ne(p)] {
		count++
	}
	if tiles[e(p)] {
		count++
	}
	if tiles[se(p)] {
		count++
	}
	if tiles[sw(p)] {
		count++
	}
	return count
}

func sw(p point) point {
	xOff := math.AbsInt(p.y) % 2
	return point{x: p.x - 1 + xOff, y: p.y + 1}
}

func se(p point) point {
	xOff := math.AbsInt(p.y) % 2
	return point{x: p.x + xOff, y: p.y + 1}
}

func e(p point) point {
	return point{x: p.x + 1, y: p.y}
}

func ne(p point) point {
	xOff := math.AbsInt(p.y) % 2
	return point{x: p.x + xOff, y: p.y - 1}
}

func nw(p point) point {
	xOff := math.AbsInt(p.y) % 2
	return point{x: p.x - 1 + xOff, y: p.y - 1}
}

func w(p point) point {
	return point{x: p.x - 1, y: p.y}
}
