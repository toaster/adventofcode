package main

import (
	"fmt"
	"strconv"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	count := 0
	last := []int{0, 0, 0}
	for _, line := range io.ReadLines() {
		value, err := strconv.Atoi(line)
		io.ReportError("failed to parse input", err)

		if last[0] != 0 {
			if value+last[2]+last[1] > last[0]+last[1]+last[2] {
				count++
			}
		}
		last[0] = last[1]
		last[1] = last[2]
		last[2] = value
	}

	fmt.Println(count)
}
