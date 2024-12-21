package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/21/space"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	codes := io.ReadLines()
	result := 0
	for _, code := range codes {
		result += io.ParseInt(code[:len(code)-1]) * space.ShortestSequence(code, 25)
	}
	fmt.Println(result)
}
