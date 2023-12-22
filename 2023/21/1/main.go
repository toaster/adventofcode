package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
	"github.com/toaster/advent_of_code/internal/util"
)

func main() {
	plan := util.ParsePlan2D(io.ReadLines(), 'S')
	reachable := math.PossibleDestinationsForExactDistance(plan.Neighbours, plan.Start, 64)
	fmt.Println(len(reachable))
}
