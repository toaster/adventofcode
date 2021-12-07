package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	nums := io.ParseInts(io.ReadLines()[0], ",")
	avgL := math.AverageInt(nums)
	avgU := avgL + 1
	fuelL := computeFuelConsumption(nums, avgL)
	fuelU := computeFuelConsumption(nums, avgU)
	fmt.Println(math.MinInt(fuelL, fuelU))
}

func computeFuelConsumption(nums []int, avg int) int {
	var fuel int
	for _, num := range nums {
		n := math.AbsInt(num - avg)
		fuel += n * (n + 1) / 2
	}
	return fuel
}
