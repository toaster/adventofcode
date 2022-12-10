package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

type tree struct {
	height      int
	scenicScore int
}

func main() {
	var trees [][]tree
	width := 0
	height := 0
	for y, line := range io.ReadLines() {
		if width == 0 {
			width = len(line)
		}
		height++
		trees = append(trees, make([]tree, width))
		for x, v := range line {
			trees[y][x].height = int(v - '0')
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			top := 0
			for u := y - 1; u > -1; u-- {
				top++
				if trees[u][x].height >= trees[y][x].height {
					break
				}
			}
			down := 0
			for d := y + 1; d < height; d++ {
				down++
				if trees[d][x].height >= trees[y][x].height {
					break
				}
			}
			left := 0
			for l := x - 1; l > -1; l-- {
				left++
				if trees[y][l].height >= trees[y][x].height {
					break
				}
			}
			right := 0
			for r := x + 1; r < width; r++ {
				right++
				if trees[y][r].height >= trees[y][x].height {
					break
				}
			}
			trees[y][x].scenicScore = top * down * left * right
		}
	}

	max := 0
	for _, row := range trees {
		for _, t := range row {
			if t.scenicScore > max {
				max = t.scenicScore
			}
		}
	}
	fmt.Println(max)
}
