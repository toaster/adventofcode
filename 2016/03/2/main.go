package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	count := 0
	var one []int
	var two []int
	var three []int
	for _, line := range io.ReadLines() {
		lengths := io.ParseInts(line, " ")
		one = append(one, lengths[0])
		two = append(two, lengths[1])
		three = append(three, lengths[2])
		if len(one) == 3 {
			if check(one) {
				count++
			}
			if check(two) {
				count++
			}
			if check(three) {
				count++
			}
			one = one[0:0]
			two = two[0:0]
			three = three[0:0]
		}
	}
	fmt.Println(count)
}

func check(lengths []int) bool {
	lengths = math.Sort(lengths)
	if lengths[0]+lengths[1] > lengths[2] {
		return true
	}
	return false
}
