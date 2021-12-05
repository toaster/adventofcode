package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var counts []int
	lineCount := 0
	for _, line := range io.ReadLines() {
		lineCount++
		for i, r := range line {
			var v int
			switch r {
			case '0':
				v = 0
			case '1':
				v = 1
			default:
				io.ReportError("", fmt.Errorf("unexpected bit: %c", r))
			}
			if i < len(counts) {
				counts[i] += v
			} else {
				counts = append(counts, v)
			}
		}
	}
	gamma := 0
	epsilon := 0
	for _, count := range counts {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if count > lineCount/2 {
			gamma++
		} else {
			epsilon++
		}
	}

	fmt.Println(gamma * epsilon)
}
