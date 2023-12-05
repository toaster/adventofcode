package main

import (
	"fmt"
	"slices"

	"github.com/toaster/advent_of_code/internal/io"
)

var naughtyPairs = [][2]rune{
	{'a', 'b'},
	{'c', 'd'},
	{'p', 'q'},
	{'x', 'y'},
}

func main() {
	nice := 0
	for _, s := range io.ReadLines() {
		if isNice(s) {
			nice++
		}
	}
	fmt.Println(nice)
}

func isNice(s string) bool {
	vowels := 0
	pairs := 0
	var last rune
	for _, c := range s {
		if c == last {
			pairs++
		}
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		if slices.Contains(naughtyPairs, [2]rune{last, c}) {
			return false
		}
		last = c
	}
	return vowels >= 3 && pairs >= 1
}
