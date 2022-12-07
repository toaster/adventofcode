package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := io.ReadAll()
	for i := 3; i < len(input); i++ {
		if input[i] != input[i-1] && input[i] != input[i-2] && input[i] != input[i-3] &&
			input[i-1] != input[i-2] && input[i-1] != input[i-3] &&
			input[i-2] != input[i-3] {
			fmt.Println(i + 1)
			break
		}
	}
}
