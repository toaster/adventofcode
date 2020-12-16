package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	validRangesByClasses := map[string][]int{}
	blocks := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	for _, line := range strings.Split(blocks[0], "\n") {
		classAndRanges := strings.Split(line, ": ")
		validRangesByClasses[classAndRanges[0]] = make([]int, 4)
		for i, n := range strings.Split(strings.ReplaceAll(classAndRanges[1], " or ", "-"), "-") {
			num, _ := strconv.Atoi(n)
			validRangesByClasses[classAndRanges[0]][i] = num
		}
	}
	var nearbyTickets [][]int
	for _, line := range strings.Split(blocks[2], "\n")[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(line))
	}

	errorRate := 0
	for _, ticket := range nearbyTickets {
		for _, num := range ticket {
			isValid := false
			for _, r := range validRangesByClasses {
				if (num >= r[0] && num <= r[1]) || (num >= r[2] && num <= r[3]) {
					isValid = true
					break
				}
			}
			if !isValid {
				errorRate += num
			}
		}
	}

	fmt.Println("error rate:", errorRate)
}

func parseTicket(line string) []int {
	var ticket []int
	for _, n := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(n)
		ticket = append(ticket, num)
	}
	return ticket
}
