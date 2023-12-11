package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	fmt.Println(computeDecompressedLength(strings.TrimSpace(io.ReadAll())))
}

func computeDecompressedLength(input string) int {
	decompressedLength := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			for j := i + 1; j < len(input); j++ {
				if input[j] == ')' {
					nums := io.ParseInts(input[i+1:j], "x")
					decompressedLength += computeDecompressedLength(input[j+1:j+1+nums[0]]) * nums[1]
					i = j + nums[0]
					break
				}
			}
		} else {
			decompressedLength++
		}
	}
	return decompressedLength
}
