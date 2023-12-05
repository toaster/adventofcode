package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2023/05/farm"
	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	seeds := parseSeeds(lines[0])
	a := farm.ParseAlmanac(lines[3:])
	fmt.Println(a.NearestLocationForSeeds(seeds))
}

func parseSeeds(input string) (ranges []*math.Range) {
	nums := io.ParseInts(input[7:], " ")
	for i := 0; i < len(nums); i += 2 {
		ranges = append(ranges, &math.Range{Start: nums[i], End: nums[i] + nums[i+1] - 1})
	}
	return
}
