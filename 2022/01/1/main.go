package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := io.ReadLines()
	max := 0
	cur := 0
	for _, line := range lines {
		if line == "" {
			cur = 0
		} else {
			cur += io.ParseInt(line)
		}
		if cur > max {
			max = cur
		}
	}
	fmt.Println(max)
}
