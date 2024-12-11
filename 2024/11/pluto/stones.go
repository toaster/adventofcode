package pluto

import (
	"strconv"

	"github.com/toaster/advent_of_code/internal/io"
)

// ParseStones creates new Stones from the input.
func ParseStones(input string) Stones {
	return io.ParseInts(input, " ")
}

// Stones represent a line of stones (https://adventofcode.com/2024/day/11).
type Stones []int

// CountAfter returns the stone count after the given amount of generations.
func (s Stones) CountAfter(generations int) int {
	trees := map[int]*genTree{}
	stones := map[int]bool{}
	for _, value := range s {
		stones[value] = true
	}
	for gen := 0; gen < generations; gen++ {
		next := map[int]bool{}
		for stone := range stones {
			if trees[stone] != nil {
				continue
			}

			if stone == 0 {
				trees[stone] = &genTree{stones: []int{1}, genCounts: map[int]int{}}
				next[1] = true
			} else if valueAsString := strconv.Itoa(stone); len(valueAsString)%2 == 0 {
				half := len(valueAsString) / 2
				first := io.ParseInt(valueAsString[0:half])
				second := io.ParseInt(valueAsString[half:])
				trees[stone] = &genTree{stones: []int{first, second}, genCounts: map[int]int{}}
				next[first] = true
				next[second] = true
			} else {
				newValue := stone * 2024
				trees[stone] = &genTree{stones: []int{newValue}, genCounts: map[int]int{}}
				next[newValue] = true
			}
		}
		stones = next
	}
	count := 0
	for _, value := range s {
		count += computeCount(value, trees, generations-1)
	}
	return count
}

func computeCount(value int, trees map[int]*genTree, generations int) int {
	tree := trees[value]
	if tree.genCounts[generations] == 0 {
		if generations == 0 {
			tree.genCounts[generations] = len(tree.stones)
		} else {
			count := 0
			for _, stone := range tree.stones {
				count += computeCount(stone, trees, generations-1)
			}
			tree.genCounts[generations] = count
		}
	}
	return tree.genCounts[generations]
}

type genTree struct {
	stones    []int
	genCounts map[int]int
}
