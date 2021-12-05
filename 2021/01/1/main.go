package main

import (
	"fmt"
	"strconv"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	count := 0
	last := 0
	for _, line := range io.ReadLines() {
		value, err := strconv.Atoi(line)
		io.ReportError("failed to parse input", err)

		if value > last && last != 0 {
			count++
		}
		last = value
	}

	fmt.Println(count)
}
