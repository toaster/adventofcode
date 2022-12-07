package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

type group struct {
	a []rune
	b []rune
	c []rune
}

func main() {
	var groups []group
	lines := io.ReadLines()
	for i := 0; i < len(lines); i += 3 {
		groups = append(groups, group{
			a: []rune(lines[i]),
			b: []rune(lines[i+1]),
			c: []rune(lines[i+2]),
		})
	}

	sum := 0
	for _, g := range groups {
		common := math.CommonElement3(g.a, g.b, g.c)
		if common == nil {
			panic("unexpectedly no common element")
		}
		sum += priority(*common)
	}
	fmt.Println(sum)
}

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	return int(item - 'A' + 27)
}
