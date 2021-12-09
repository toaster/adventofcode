package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var rows [][]int
	for _, line := range io.ReadLines() {
		var row []int
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		rows = append(rows, row)
	}
	height := len(rows)
	width := len(rows[0])

	sum := 0
	for y, row := range rows {
		for x, level := range row {
			fmt.Print(level)
			if y > 0 && rows[y-1][x] <= level {
				continue
			}
			if y < height-1 && rows[y+1][x] <= level {
				continue
			}
			if x > 0 && row[x-1] <= level {
				continue
			}
			if x < width-1 && row[x+1] <= level {
				continue
			}
			sum += level + 1
		}
		fmt.Println(" ", sum)
	}
	fmt.Println(sum)
}
