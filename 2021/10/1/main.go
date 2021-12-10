package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

var illegal = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var closers = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	score := 0
	for _, line := range io.ReadLines() {
		score += parseLine(line)
	}
	fmt.Println(score)
}

func parseLine(line string) int {
	var open []rune
	for x, r := range line {
		switch r {
		case '(', '[', '{', '<':
			open = append(open, r)
		default:
			i := len(open) - 1
			if closers[open[i]] == r {
				open = open[:i]
			} else {
				fmt.Printf("%d: %c (%d)\n", x, r, illegal[r])
				return illegal[r]
			}
		}
	}
	return 0
}
