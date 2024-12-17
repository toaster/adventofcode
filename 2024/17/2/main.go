package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/2024/17/cc"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := io.ReadLines()
	c := cc.Parse(lines)
	expectedInts := io.ParseInts(lines[4][9:], ",")
	// The last iteration has to stop.
	// Given my input, this is the case for an A of 0-7 (divided by 2^3=8 before JNZ).
	possibleValues := []int{0, 1, 2, 3, 5, 6, 7}
	// For every tail of digits (starting from the end) compute the possible A’s leading to it
	// restricted by the expected A of the earlier run with a shorter tail.
	for i := len(expectedInts) - 1; i >= 0; i-- {
		var expectedStrings []string
		for _, value := range expectedInts[i:] {
			expectedStrings = append(expectedStrings, strconv.Itoa(value))
		}
		expected := strings.Join(expectedStrings, ",")
		var nextPossibleValues []int
		for _, regAValue := range possibleValues {
			out := make(chan int)
			c.PatchRegisterA(regAValue)
			go c.Run(out)
			var output []string
			for value := range out {
				output = append(output, strconv.Itoa(value))
			}
			if expected == strings.Join(output, ",") {
				if i == 0 {
					fmt.Println(regAValue)
					return
				}

				start := 8 * regAValue
				for j := 0; j < 8; j++ {
					nextPossibleValues = append(nextPossibleValues, start+j)
				}
			}
		}
		possibleValues = nextPossibleValues
	}
}
