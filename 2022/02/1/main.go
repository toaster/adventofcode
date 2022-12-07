package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

var scores = map[string]int{
	"A X": 1 + 3,
	"A Y": 2 + 6,
	"A Z": 3 + 0,
	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,
	"C X": 1 + 6,
	"C Y": 2 + 0,
	"C Z": 3 + 3,
}

func main() {
	lines := io.ReadLines()
	total := 0
	for _, line := range lines {
		total += scores[line]
	}
	fmt.Println(total)
}
