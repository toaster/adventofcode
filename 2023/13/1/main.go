package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
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

	if reflection := detectReflection(rows); reflection > 0 {
		result += 100 * reflection
	}

	if reflection := detectReflection(cols); reflection > 0 {
		result += reflection
	}

	return result
}

func detectReflection(values []int) int {
	last := 0
	for i, v := range values {
		if i > 0 && v == last && isReflection(values, i) {
			return i
		}
		last = v
	}
	return 0
}

func isReflection(values []int, index int) bool {
	for i := 0; i < len(values)-index && index-i > 0; i++ {
		if values[index+i] != values[index-1-i] {
			return false
		}
	}
	return true
}
