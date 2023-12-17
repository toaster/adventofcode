package main

import (
	"fmt"

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
		mxC := ' '
		mx := 0
		for c, i := range letter {
			if i > mx {
				mxC = c
				mx = i
			}
		}
		fmt.Printf("%c", mxC)
	}
	fmt.Println()
}
