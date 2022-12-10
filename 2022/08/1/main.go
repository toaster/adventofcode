package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

type tree struct {
	height  int
	visible bool
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

	for x := 0; x < width; x++ {
		trees[0][x].visible = true
		trees[height-1][x].visible = true
		limit := trees[0][x].height
		for y := 1; y < height; y++ {
			if trees[y][x].height > limit {
				trees[y][x].visible = true
				limit = trees[y][x].height
			}
		}
		limit = trees[height-1][x].height
		for y := height - 2; y > 0; y-- {
			if trees[y][x].height > limit {
				trees[y][x].visible = true
				limit = trees[y][x].height
			}
		}
	}
	for y := 0; y < height; y++ {
		trees[y][0].visible = true
		trees[y][width-1].visible = true
		limit := trees[y][0].height
		for x := 1; x < width; x++ {
			if trees[y][x].height > limit {
				trees[y][x].visible = true
				limit = trees[y][x].height
			}
		}
		limit = trees[y][width-1].height
		for x := width - 2; x > 0; x-- {
			if trees[y][x].height > limit {
				trees[y][x].visible = true
				limit = trees[y][x].height
			}
		}
	}

	printTrees(trees)

	count := 0
	for _, row := range trees {
		for _, t := range row {
			if t.visible {
				count++
			}
		}
	}
	fmt.Println(count)
}

func printTrees(trees [][]tree) {
	for _, row := range trees {
		for _, t := range row {
			if t.visible {
				fmt.Printf(" %d ", t.height)
			} else {
				fmt.Printf("(%d)", t.height)
			}
		}
		fmt.Println()
	}
}
