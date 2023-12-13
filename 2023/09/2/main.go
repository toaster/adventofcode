package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	sum := 0
	for _, line := range io.ReadLines() {
		historyValues := io.ParseInts(line, " ")
		sum += predictPreviousOasisValue(historyValues)
	}
	fmt.Println(sum)
}

func predictPreviousOasisValue(input []int) int {
	var derivation []int
	isConstant := true
	for i := 1; i < len(input); i++ {
		v := input[i] - input[i-1]
		derivation = append(derivation, v)
		if v != 0 {
			isConstant = false
		}
	}
	if isConstant {
		return input[0]
	}

	return input[0] - predictPreviousOasisValue(derivation)
}
