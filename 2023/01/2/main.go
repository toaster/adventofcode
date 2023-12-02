package main

import (
	"fmt"
	"strings"

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

	if strings.HasPrefix(s, "zero") {
		return true, 0
	} else if strings.HasPrefix(s, "one") {
		return true, 1
	} else if strings.HasPrefix(s, "two") {
		return true, 2
	} else if strings.HasPrefix(s, "three") {
		return true, 3
	} else if strings.HasPrefix(s, "four") {
		return true, 4
	} else if strings.HasPrefix(s, "five") {
		return true, 5
	} else if strings.HasPrefix(s, "six") {
		return true, 6
	} else if strings.HasPrefix(s, "seven") {
		return true, 7
	} else if strings.HasPrefix(s, "eight") {
		return true, 8
	} else if strings.HasPrefix(s, "nine") {
		return true, 9
	}
	return false, 0
}
