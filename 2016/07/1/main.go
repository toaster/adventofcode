package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	tlsCount := 0
	for _, address := range io.ReadLines() {
		if supportsTLS(address) {
			tlsCount++
		}
	}
	fmt.Println(tlsCount)
}

func supportsTLS(address string) bool {
	inHypernetSequence := false
	result := false
	for i := 0; i < len(address); i++ {
		c := address[i]
		if c == '[' {
			inHypernetSequence = true
			continue
		} else if c == ']' {
			inHypernetSequence = false
			continue
		}
		if i < 3 {
			continue
		}
		if address[i-3] == c && address[i-2] == address[i-1] && address[i-1] != c {
			if inHypernetSequence {
				return false
			}
			result = true
		}
	}
	return result
}
