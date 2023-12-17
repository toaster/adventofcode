package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := strings.TrimSpace(io.ReadAll())
	decompressedLength := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			for j := i + 1; j < len(input); j++ {
				if input[j] == ')' {
					nums := io.ParseInts(input[i+1:j], "x")
					decompressedLength += nums[0] * nums[1]
					i = j + nums[0]
					break
				}
			}
		} else {
			decompressedLength++
		}
	}
	fmt.Println(decompressedLength)
}
