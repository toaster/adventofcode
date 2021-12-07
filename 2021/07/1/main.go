package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	nums := io.ParseInts(io.ReadLines()[0], ",")
	m := math.MedianInt(nums)
	var fuel int
	for _, num := range nums {
		fuel += math.AbsInt(num - m)
	}
	fmt.Println(fuel)
}
