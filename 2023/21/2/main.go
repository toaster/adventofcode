package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	plan := math.ParsePlan2D(io.ReadLines(), 'S')
	reachable := math.PossibleDestinationsForExactDistance(plan.Neighbours, plan.Start, 26501365)
	fmt.Println(len(reachable))
}
