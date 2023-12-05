package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

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
	inBetweens := 0
	multiPairs := 0
	pairs := map[[2]rune]int{}
	var last rune
	var beforeLast rune
	didOverlap := false
	for _, c := range s {
		overlap := false
		if c == beforeLast {
			inBetweens++
		}
		if last != rune(0) {
			pair := [2]rune{last, c}
			if !didOverlap && beforeLast != rune(0) {
				pairBefore := [2]rune{beforeLast, last}
				if pairBefore == pair {
					overlap = true
				}
			}
			if !overlap {
				pairs[pair]++
				if pairs[pair] == 2 {
					multiPairs++
				}
			}
		}
		didOverlap = overlap
		beforeLast = last
		last = c
	}
	return inBetweens > 0 && multiPairs > 0
}
