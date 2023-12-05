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

func parseSeeds(input string) (result []*math.Range) {
	for _, start := range io.ParseInts(input[7:], " ") {
		result = append(result, &math.Range{Start: start, End: start})
	}
	return
}
