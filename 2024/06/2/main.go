package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	plan := math.ParsePlan2D(io.ReadLines(), '^')
	route := plan.GetGuardRoute()
	candidates := route.Points()
	count := 0
	for _, candidate := range candidates {
		if candidate == plan.Start {
			continue
		}
		newPlan := plan.AddObstacle(candidate)
		if newPlan.GetGuardRoute().IsLoop() {
			count++
		}
	}
	fmt.Println(count)
}
