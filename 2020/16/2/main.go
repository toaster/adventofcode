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
	myTicket := parseTicket(strings.Split(blocks[1], "\n")[1])
	var nearbyTickets [][]int
	for _, line := range strings.Split(blocks[2], "\n")[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(line))
	}

	var validTickets [][]int
	for _, ticket := range nearbyTickets {
		ticketIsValid := true
		for _, num := range ticket {
			numIsValid := false
			for _, r := range validRangesByClasses {
				if (num >= r[0] && num <= r[1]) || (num >= r[2] && num <= r[3]) {
					numIsValid = true
					break
				}
			}
			if !numIsValid {
				ticketIsValid = false
				break
			}
		}
		if ticketIsValid {
			validTickets = append(validTickets, ticket)
		}
	}
	var candidatesForTickets [][][]string
	for _, ticket := range validTickets {
		var candidates [][]string
		for i, num := range ticket {
			for class, r := range validRangesByClasses {
				if (num >= r[0] && num <= r[1]) || (num >= r[2] && num <= r[3]) {
					if len(candidates) < i+1 {
						candidates = append(candidates, []string{class})
					} else if !contains(candidates[i], class) {
						candidates[i] = append(candidates[i], class)
					}
				} else {
				}
			}
		}
		candidatesForTickets = append(candidatesForTickets, candidates)
	}
	var candidates [][]string
	for _, candidatesForTicket := range candidatesForTickets {
		if len(candidates) == 0 {
			candidates = make([][]string, len(candidatesForTicket))
			copy(candidates, candidatesForTicket)
			continue
		}

		for i, c := range candidatesForTicket {
			candidates[i] = intersection(candidates[i], c)
		}
	}
	var found []string
	for {
		done := true
		for i, candidate := range candidates {
			if len(candidate) == 1 {
				if !contains(found, candidate[0]) {
					found = append(found, candidate[0])
				}
				continue
			}

			candidates[i] = minus(candidates[i], found)
			done = false
		}
		if done {
			break
		}
	}
	result := 1
	for i, classes := range candidates {
		if len(classes) == 0 {
			continue
		}
		class := classes[0]
		if strings.HasPrefix(class, "departure") {
			result *= myTicket[i]
		}
	}
	fmt.Println("result:", result)
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func intersection(a, b []string) []string {
	var result []string
	for _, s := range a {
		if contains(b, s) {
			result = append(result, s)
		}
	}
	return result
}

func minus(a, b []string) []string {
	var result []string
	for _, s := range a {
		if !contains(b, s) {
			result = append(result, s)
		}
	}
	return result
}

func parseTicket(line string) []int {
	var ticket []int
	for _, n := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(n)
		ticket = append(ticket, num)
	}
	return ticket
}
