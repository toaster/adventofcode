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
				p.x += math.AbsInt(p.y) % 2
				p.y++
				line = line[2:]
			} else if strings.HasPrefix(line, "sw") {
				p.y++
				p.x -= math.AbsInt(p.y) % 2
				line = line[2:]
			} else if strings.HasPrefix(line, "w") {
				p.x--
				line = line[1:]
			} else if strings.HasPrefix(line, "nw") {
				p.y--
				p.x -= math.AbsInt(p.y) % 2
				line = line[2:]
			} else if strings.HasPrefix(line, "ne") {
				p.x += math.AbsInt(p.y) % 2
				p.y--
				line = line[2:]
			} else if strings.HasPrefix(line, "e") {
				p.x++
				line = line[1:]
			}
		}
		// fmt.Println("flip", p)
		tiles[p] = !tiles[p]
	}
	count := 0
	for _, b := range tiles {
		if b {
			count++
		}
	}
	fmt.Println("black tiles:", count)
}
