package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := io.ReadAll()
	markerLength := 14
	for i := markerLength - 1; i < len(input); i++ {
		if isMarker(input[i-(markerLength-1):], markerLength) {
			fmt.Println(i + 1)
			return
		}
	}
}

func isMarker(input string, markerLength int) bool {
	i := markerLength - 1
	detector := map[uint8]bool{input[i]: true}
	for j := 0; j < i; j++ {
		if detector[input[j]] {
			return false
		}
		detector[input[j]] = true
	}
	return true
}
