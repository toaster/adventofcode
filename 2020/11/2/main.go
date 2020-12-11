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
			if count >= 5 {
				nextPlan[y][x] = 'L'
			}
		}
	}
	return nextPlan
}

func countOccupiedAdjacentSeats(plan [][]rune, x, y int) int {
	count := 0
	if isDirectionOccupied(plan, x, y, -1, -1) {
		count++
	}
	if isDirectionOccupied(plan, x, y, 0, -1) {
		count++
	}
	if isDirectionOccupied(plan, x, y, 1, -1) {
		count++
	}
	if isDirectionOccupied(plan, x, y, -1, 0) {
		count++
	}
	if isDirectionOccupied(plan, x, y, 1, 0) {
		count++
	}
	if isDirectionOccupied(plan, x, y, -1, 1) {
		count++
	}
	if isDirectionOccupied(plan, x, y, 0, 1) {
		count++
	}
	if isDirectionOccupied(plan, x, y, 1, 1) {
		count++
	}
	return count
}

func isDirectionOccupied(plan [][]rune, x, y, headingX, headingY int) bool {
	checkX := x + headingX
	if checkX < 0 || checkX >= len(plan[0]) {
		return false
	}
	checkY := y + headingY
	if checkY < 0 || checkY >= len(plan) {
		return false
	}
	// fmt.Println("check", checkX, checkY, string(plan[checkY][checkX]))
	if plan[checkY][checkX] == '#' {
		return true
	}
	if plan[checkY][checkX] == 'L' {
		return false
	}
	return isDirectionOccupied(plan, checkX, checkY, headingX, headingY)
}
