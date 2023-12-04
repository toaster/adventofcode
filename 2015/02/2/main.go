package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	sum := 0
	for _, line := range io.ReadLines() {
		dims := io.ParseInts(line, "x")
		length := dims[0]
		width := dims[1]
		height := dims[2]
		sum += 2*min(length, width) + 2*min(max(length, width), height) + length*width*height
	}
	fmt.Println(sum)
}
