package main

import (
	"fmt"
	"sort"

	"github.com/toaster/advent_of_code/internal/io"
)

var scores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var closers = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	var allScores []int
	for _, line := range io.ReadLines() {
		score := parseLine(line)
		if score != 0 {
			allScores = append(allScores, score)
		}
	}
	sort.Ints(allScores)
	fmt.Println(allScores[len(allScores)/2])
}

func parseLine(line string) int {
	var open []rune
	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			open = append(open, r)
		default:
			if len(open) == 0 {
				break
			}
			i := len(open) - 1
			if closers[open[i]] == r {
				open = open[:i]
			} else {
				return 0
			}
		}
	}
	score := 0
	for i := len(open) - 1; i >= 0; i-- {
		score *= 5
		score += scores[closers[open[i]]]
	}
	return score
}
