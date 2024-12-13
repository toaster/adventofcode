package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/2024/13/tropicalisland"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var machines []tropicalisland.ClawMachine
	for _, input := range strings.Split(io.ReadAll(), "\n\n") {
		machines = append(machines, tropicalisland.ParseClawMachine(input, 0))
	}
	total := 0
	for _, machine := range machines {
		total += machine.ComputeMinimalCostsIfPossible()
	}
	fmt.Println(total)
}
