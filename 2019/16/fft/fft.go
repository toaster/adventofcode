package fft

import (
	"github.com/toaster/advent_of_code/internal/math"
)

var pattern = [4]int{0, 1, 0, -1}

// Perform performs phaseCount phases of the Flawed Frequency Transmission algorithm on input.
func Perform(input []int, phaseCount int) []int {
	data := input
	for p := 0; p < phaseCount; p++ {
		var nextData []int
		for i := 0; i < len(data); i++ {
			v := 0
			for l := 0; l < len(data); l++ {
				pi := ((l + 1) / (i + 1)) % 4
				v += data[l] * pattern[pi]
			}
			nextData = append(nextData, math.AbsInt(v)%10)
		}
		data = nextData
	}
	return data
}

// PerformHuge performs phaseCount phases of the Flawed Frequency Transmission algorithm on
// segment of the 10,000 times multiplied input and returns only the 8 significant digits.
func PerformHuge(fragment []int, phaseCount int) []int {
	inputLen := len(fragment) * 10000
	var offset int
	for i := 0; i < 7; i++ {
		offset *= 10
		offset += fragment[i]
	}
	if offset < inputLen/2 {
		panic("offset too low, cannot use highly optimized perform")
	}
	dataLen := inputLen - offset
	factor := dataLen/len(fragment) + 1
	var input []int
	for j := 0; j < factor; j++ {
		input = append(input, fragment...)
	}
	data := input[len(input)-dataLen:]
	for p := 0; p < phaseCount; p++ {
		for l := len(data) - 2; l >= 0; l-- {
			data[l] = (data[l] + data[l+1]) % 10
		}
	}
	return data[:8]
}
