package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	start := 0
	var rows []int
	var cols []int
	sum := 0
	lines := io.ReadLines()
	lines = append(lines, "")
	for i, line := range lines {
		if line == "" {
			sum += analyze(rows, cols)
			start = i + 1
			rows = nil
			cols = nil
			continue
		}

		y := i - start
		if len(cols) == 0 {
			cols = make([]int, len(line))
		}
		rows = append(rows, 0)
		for x, c := range line {
			if c == '#' {
				rows[y] += 1 << x
				cols[x] += 1 << y
			}
		}
	}
	fmt.Println(sum)
}

func analyze(rows, cols []int) int {
	result := 0

	if reflection := detectSmudgedReflection(rows); reflection > 0 {
		result += 100 * reflection
	}

	if reflection := detectSmudgedReflection(cols); reflection > 0 {
		result += reflection
	}

	return result
}

func detectSmudgedReflection(values []int) int {
	for i := range values {
		if i > 0 && isSmudgedReflection(values, i) {
			return i
		}
	}
	return 0
}

func isSmudgedReflection(values []int, index int) bool {
	smudgeCount := 0
	for i := 0; i < len(values)-index && index-i > 0; i++ {
		a := values[index+i]
		b := values[index-1-i]
		if a != b {
			if math.CountSetBits(a^b) == 1 && smudgeCount == 0 {
				smudgeCount++
			} else {
				return false
			}
		}
	}
	return smudgeCount == 1
}
