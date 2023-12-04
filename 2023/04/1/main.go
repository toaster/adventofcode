package main

import (
	"fmt"
	math2 "math"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	var wins [][]int
	var nums [][]int
	for _, line := range lines {
		raw := strings.Split(strings.Split(line, ": ")[1], " | ")
		wins = append(wins, io.ParseInts(raw[0], " "))
		nums = append(nums, io.ParseInts(raw[1], " "))
	}
	sum := 0
	for i := 0; i < len(wins); i++ {
		common := math.Intersection(nums[i], wins[i])
		count := len(common)
		worth := 0
		if count > 0 {
			worth = int(math2.Pow(2, float64(count-1)))
		}
		fmt.Printf("- Card %d has %d winning number(s) (%#v), so it is worth %d points.\n", i+1, count, common, worth)
		sum += worth
	}
	fmt.Println(sum)
}
