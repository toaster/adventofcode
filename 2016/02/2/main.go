package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	pad := map[math.Point2D]byte{
		{0, -2}:  '1',
		{-1, -1}: '2',
		{0, -1}:  '3',
		{1, -1}:  '4',
		{-2, 0}:  '5',
		{-1, 0}:  '6',
		{0, 0}:   '7',
		{1, 0}:   '8',
		{2, 0}:   '9',
		{-1, 1}:  'A',
		{0, 1}:   'B',
		{1, 1}:   'C',
		{0, 2}:   'D',
	}
	var code []byte
	loc := math.Point2D{X: -2}
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
	fmt.Println(string(code))
}
