package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := io.ReadLines()
	sum := 0
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if isDigit, value := detectDigit(line[i:]); isDigit {
				sum += value * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if isDigit, value := detectDigit(line[i:]); isDigit {
				sum += value
				break
			}
		}
	}
	fmt.Println(sum)
}

func detectDigit(s string) (bool, int) {
	c := s[0]
	if c >= '0' && c <= '9' {
		return true, int(c - '0')
	}
	return false, 0
}
