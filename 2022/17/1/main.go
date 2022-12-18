package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2022/17/tetris"
	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	jets := io.ReadLines()[0]
	pieces := []*tetris.Piece{
		{ // -
			Height: 1,
			Shape: map[math.Point2D]bool{
				{0, 0}: true,
				{1, 0}: true,
				{2, 0}: true,
				{3, 0}: true,
			},
			Width: 4,
		},
		{ // +
			Height: 3,
			Shape: map[math.Point2D]bool{
				{1, 0}: true,
				{0, 1}: true,
				{1, 1}: true,
				{2, 1}: true,
				{1, 2}: true,
			},
			Width: 3,
		},
		{ // ⌟
			Height: 3,
			Shape: map[math.Point2D]bool{
				{0, 0}: true,
				{1, 0}: true,
				{2, 0}: true,
				{2, 1}: true,
				{2, 2}: true,
			},
			Width: 3,
		},
		{ // |
			Height: 4,
			Shape: map[math.Point2D]bool{
				{0, 0}: true,
				{0, 1}: true,
				{0, 2}: true,
				{0, 3}: true,
			},
			Width: 1,
		},
		{ // .
			Height: 2,
			Shape: map[math.Point2D]bool{
				{0, 0}: true,
				{0, 1}: true,
				{1, 0}: true,
				{1, 1}: true,
			},
			Width: 2,
		},
	}
	fmt.Println(tetris.Play(7, 2022, pieces, []rune(jets)))
}
