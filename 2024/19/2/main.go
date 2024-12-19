package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := io.ReadLines()
	patterns := map[string]bool{}
	maxPatternLength := 0
	for _, pattern := range strings.Split(input[0], ", ") {
		patterns[pattern] = true
		if len(pattern) > maxPatternLength {
			maxPatternLength = len(pattern)
		}
	}
	requestedDesigns := input[2:]
	count := 0
	knownCounts := map[string]int{"": 1}
	for _, design := range requestedDesigns {
		count += countRealizations(design, patterns, maxPatternLength, knownCounts)
	}
	fmt.Println(count)
}

func countRealizations(design string, patterns map[string]bool, maxPatternLength int, knownCounts map[string]int) int {
	if count, ok := knownCounts[design]; ok {
		return count
	}

	count := 0
	for i := 1; i <= maxPatternLength && i <= len(design); i++ {
		prefix := design[:i]
		if patterns[prefix] {
			count += countRealizations(design[i:], patterns, maxPatternLength, knownCounts)
		}
	}
	knownCounts[design] = count
	return count
}
