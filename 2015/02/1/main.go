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
		one := length * width
		two := width * height
		three := length * height
		extra := min(min(one, two), three)
		sum += 2*one + 2*two + 2*three + extra
	}
	fmt.Println(sum)
}
