package main

import (
	"fmt"

	calibration "github.com/toaster/advent_of_code/2024/07"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var equations []calibration.Equation
	for _, line := range io.ReadLines() {
		equations = append(equations, calibration.ParseEquation(line))
	}

	sum := 0
	for _, e := range equations {
		if e.CanBeTrue(3) {
			sum += e.Result()
		}
	}
	fmt.Println(sum)
}
