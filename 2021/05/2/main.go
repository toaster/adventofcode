package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	io.ReportError("failed to read standard input", err)

	var max point
	plan := map[point]int{}
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		mp := parseLine(line, plan)
		if mp.x > max.x {
			max.x = mp.x
		}
		if mp.y > max.y {
			max.y = mp.y
		}
	}

	// printPlan(plan, max)
	count := 0
	for _, i := range plan {
		if i > 1 {
			count++
		}
	}
	fmt.Println(count)
}

func printPlan(plan map[point]int, max point) {
	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
			v := plan[point{x, y}]
			if v > 0 {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseLine(line string, plan map[point]int) point {
	points := strings.Split(line, " -> ")
	start := parsePoint(points[0])
	end := parsePoint(points[1])

	if start.x != end.x {
		if start.x > end.x {
			start, end = end, start
		}
		for x, y := start.x, start.y; x <= end.x; x++ {
			plan[point{x, y}]++
			if start.y < end.y {
				y++
			} else if start.y > end.y {
				y--
			}
		}
	} else {
		if start.y > end.y {
			start, end = end, start
		}
		for x, y := start.x, start.y; y <= end.y; y++ {
			plan[point{x, y}]++
		}
	}
	return point{max(start.x, end.x), max(start.y, end.y)}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func parsePoint(s string) point {
	nums := parseNumbers(s, ",")
	return point{nums[0], nums[1]}
}

func parseNumbers(line string, sep string) (numbers []int) {
	for _, s := range strings.Split(line, sep) {
		if s == "" {
			continue
		}

		v, err := strconv.Atoi(s)
		if err != nil {
			io.ReportError("failed to parse input", err)
		}

		numbers = append(numbers, v)
	}
	return
}

type point struct {
	x int
	y int
}
