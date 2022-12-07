package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

type rucksack struct {
	compartment1 []rune
	compartment2 []rune
}

func main() {
	var sacks []rucksack
	for _, line := range io.ReadLines() {
		itemCount := len(line) / 2
		sacks = append(sacks, rucksack{
			compartment1: []rune(line[0:itemCount]),
			compartment2: []rune(line[itemCount:]),
		})
	}

	sum := 0
	for _, sack := range sacks {
		common := math.CommonElement2(sack.compartment1, sack.compartment2)
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
