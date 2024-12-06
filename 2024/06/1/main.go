package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	plan := math.ParsePlan2D(io.ReadLines(), '^')
	fmt.Println(plan.CountVisitedLocationsOfGuard())
}
