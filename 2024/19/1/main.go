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
	for _, design := range requestedDesigns {
		if canBeRealized(design, patterns, maxPatternLength) {
			count++
		}
	}
	fmt.Println(count)
}

func canBeRealized(design string, patterns map[string]bool, maxPatternLength int) bool {
	if realizable, ok := patterns[design]; ok {
		return realizable
	}

	realizable := false
	for i := 1; i <= maxPatternLength && i <= len(design); i++ {
		prefix := design[:i]
		if patterns[prefix] && canBeRealized(design[i:], patterns, maxPatternLength) {
			realizable = true
			break
		}
	}
	patterns[design] = realizable
	return realizable
}
