package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	var plan, nextPlan [][]rune
	for _, line := range lines {
		plan = append(plan, []rune(line))
	}
	i := 0
	// for _, row := range plan {
	// 	fmt.Println(string(row))
	// }
	for ; !arePlansEqual(plan, nextPlan); i++ {
		nextPlan = performOneRound(plan)
		plan, nextPlan = nextPlan, plan
		// fmt.Println(i)
		// for _, row := range plan {
		// 	fmt.Println(string(row))
		// }
		// time.Sleep(1 * time.Millisecond)
	}
	occupiedSeats := 0
	for _, row := range plan {
		for _, seat := range row {
			if seat == '#' {
				occupiedSeats++
			}
		}
	}
	fmt.Println("occupied seats:", occupiedSeats)
}

func arePlansEqual(a, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			return false
		}
	}
	return true
}

func performOneRound(plan [][]rune) [][]rune {
	nextPlan := make([][]rune, len(plan))
	for y, row := range plan {
		nextPlan[y] = make([]rune, len(row))
		for x, seat := range row {
			nextPlan[y][x] = seat
			if seat == '.' {
				continue
			}
			count := countOccupiedAdjacentSeats(plan, x, y)
			if seat == 'L' {
				if count == 0 {
					nextPlan[y][x] = '#'
				}
				continue
			}
			if count >= 4 {
				nextPlan[y][x] = 'L'
			}
		}
	}
	return nextPlan
}

func countOccupiedAdjacentSeats(plan [][]rune, x, y int) int {
	count := 0
	if y > 0 {
		row := plan[y-1]
		if x > 0 && row[x-1] == '#' {
			count++
		}
		if row[x] == '#' {
			count++
		}
		if x < len(row)-1 && row[x+1] == '#' {
			count++
		}
	}
	{
		row := plan[y]
		if x > 0 && row[x-1] == '#' {
			count++
		}
		if x < len(row)-1 && row[x+1] == '#' {
			count++
		}
	}
	if y < len(plan)-1 {
		row := plan[y+1]
		if x > 0 && row[x-1] == '#' {
			count++
		}
		if row[x] == '#' {
			count++
		}
		if x < len(row)-1 && row[x+1] == '#' {
			count++
		}
	}
	return count
}
