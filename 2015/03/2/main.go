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
	roboLoc := math.Point2D{}
	for i := 0; i < len(directions); i += 2 {
		loc = move(directions[i], loc)
		roboLoc = move(directions[i+1], roboLoc)
		m[loc]++
		m[roboLoc]++
	}
	fmt.Println(len(m))
}

func move(direction byte, loc math.Point2D) math.Point2D {
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
