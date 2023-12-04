package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	pad := map[math.Point2D]int{
		{-1, -1}: 1,
		{0, -1}:  2,
		{1, -1}:  3,
		{-1, 0}:  4,
		{0, 0}:   5,
		{1, 0}:   6,
		{-1, 1}:  7,
		{0, 1}:   8,
		{1, 1}:   9,
	}
	var code = []int{}
	loc := math.Point2D{}
	for _, line := range io.ReadLines() {
		for _, c := range line {
			var next math.Point2D
			switch c {
			case 'U':
				next = loc.SubtractXY(0, 1)
			case 'D':
				next = loc.AddXY(0, 1)
			case 'L':
				next = loc.SubtractXY(1, 0)
			case 'R':
				next = loc.AddXY(1, 0)
			}
			if pad[next] != 0 {
				loc = next
			}
		}
		code = append(code, pad[loc])
	}
	fmt.Println(code)
}
