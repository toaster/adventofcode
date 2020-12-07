package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
	}
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	result := 1
	for _, slope := range []struct {
		x, y int
	}{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	} {
		result *= countTreesOnSlope(slope.x, slope.y, lines)
	}
	fmt.Println("result:", result)
}

func countTreesOnSlope(headingX int, headingY int, lines []string) int {
	patternWidth := len(lines[0])
	x := 0
	trees := 0
	for y := 0; y < len(lines); y += headingY {
		if lines[y][x] == '#' {
			trees++
		}
		x += headingX
		x = x % patternWidth
	}
	return trees
}
