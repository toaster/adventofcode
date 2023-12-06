package main

import (
	"fmt"
	math2 "math"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	messages := io.ReadLines()
	letters := [8]map[rune]int{{}, {}, {}, {}, {}, {}, {}, {}}
	for _, message := range messages {
		for i, c := range message {
			letters[i][c]++
		}
	}
	for _, letter := range letters {
		mnC := ' '
		mn := math2.MaxInt
		for c, i := range letter {
			if i < mn {
				mnC = c
				mn = i
			}
		}
		fmt.Printf("%c", mnC)
	}
	fmt.Println()
}
