package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	directions := strings.TrimSpace(io.ReadAll())
	m := map[math.Point2D]int{}
	loc := math.Point2D{}
	for _, d := range directions {
		loc = move(d, loc)
		m[loc]++
	}
	fmt.Println(len(m))
}

func move(direction int32, loc math.Point2D) math.Point2D {
	switch direction {
	case '>':
		loc = loc.AddXY(1, 0)
	case '<':
		loc = loc.SubtractXY(1, 0)
	case '^':
		loc = loc.SubtractXY(0, 1)
	case 'v':
		loc = loc.AddXY(0, 1)
	}
	return loc
}
