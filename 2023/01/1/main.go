package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	sum := 0
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if value, isDigit := math.DetectDigit(line[i]); isDigit {
				sum += value * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if value, isDigit := math.DetectDigit(line[i]); isDigit {
				sum += value
				break
			}
		}
	}
	fmt.Println(sum)
}
