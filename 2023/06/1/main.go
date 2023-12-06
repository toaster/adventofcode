package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := io.ReadLines()
	times := io.ParseInts(lines[0][6:], " ")
	distances := io.ParseInts(lines[1][9:], " ")
	product := 1
	for i, time := range times {
		waysToWin := 0
		distanceToBeat := distances[i]
		for j := 1; j < time; j++ {
			if j*(time-j) > distanceToBeat {
				waysToWin++
			}
		}
		product *= waysToWin
	}
	fmt.Println(product)
}
