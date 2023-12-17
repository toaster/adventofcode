package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	count := 0
	for _, line := range io.ReadLines() {
		lengths := io.ParseInts(line, " ")
		lengths = math.Sort(lengths)
		if lengths[0]+lengths[1] > lengths[2] {
			count++
		}
	}
	fmt.Println(count)
}
